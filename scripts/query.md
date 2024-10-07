# Query

Các yêu cầu truy vấn:

- Query 1: Tính lương của các nhân viên kinh doanh dựa trên số tài khoản ngân hàng mà họ đã tạo được cho khách hàng trong mỗi tháng. Mỗi tài khoản ghi nợ (debit) được tạo ra sẽ được cộng 500 nghìn, mỗi tài khoản tín dụng (credit) sẽ được cộng 2% số tiền mà khách gửi lần đầu.
- Query 2: Liệt kê tên khách hàng cùng số tiền giao dịch trên mỗi tài khoản tín dụng (credit) mà họ thực hiện trong khoảng thời gian từ ngày bắt đầu đến ngày kết thúc.
- Query 3: Liệt kê thông tin các tài khoản ghi nợ (debit) cùng tổng số nợ tồn đọng tại thời điểm truy vấn, danh sách được sắp xếp theo thứ tự giảm dần số dự nợ.
- Query 4: Liệt kê thông tin của 10 khách hàng có tổng số tiền gửi trên các tài khoản của họ là nhiều nhất.
- Query 5: Thực hiện các giao dịch thanh toán nợ cho tài khoản ghi nợ từ tài khoản tín dụng.

## Query 1

```javascript
db.employees.aggregate([
  {
    $unwind: "$accounts_created",
  },
  {
    $addFields: {
      month: { $month: "$accounts_created.created_at" },
    },
  },
  {
    $group: {
      _id: { employee_id: "$employee_id", month: "$month" },
      total_credit_bonus: {
        $sum: {
          $cond: [{ $eq: ["$accounts_created.type", "credit"] }, 500000, 0],
        },
      },
      total_savings_bonus: {
        $sum: {
          $cond: [
            { $eq: ["$accounts_created.type", "savings"] },
            { $multiply: [0.02, "$accounts_created.initial_deposit"] },
            0,
          ],
        },
      },
    },
  },
  {
    $addFields: {
      total_salary: { $add: ["$total_credit_bonus", "$total_savings_bonus"] },
    },
  },
  {
    $project: {
      _id: 0,
      employee_id: "$_id.employee_id",
      month: "$_id.month",
      total_salary: 1,
    },
  },
]);
```

## Query 2

```javascript
db.customer.aggregate([
  {
    $unwind: "$accounts",
  },
  {
    $match: { "accounts.account_type": "credit" },
  },
  {
    $unwind: "$accounts.transactions",
  },
  {
    $match: {
      "accounts.transactions.date": {
        $gte: ISODate("2024-01-01"),
        $lte: ISODate("2024-12-31"),
      },
    },
  },
  {
    $project: {
      _id: 0,
      customer_name: "$name",
      transaction_amount: "$accounts.transactions.amount",
      transaction_date: "$accounts.transactions.date",
    },
  },
]);
```

## Query 3

```javascript
db.customer.aggregate([
  {
    $unwind: "$accounts",
  },
  {
    $match: { "accounts.account_type": "credit" },
  },
  {
    $group: {
      _id: "$accounts.account_id",
      customer_name: { $first: "$name" },
      credit_limit: { $first: "$accounts.credit_limit" },
      outstanding_debt: { $first: "$accounts.outstanding_debt" },
    },
  },
  {
    $sort: { outstanding_debt: -1 },
  },
  {
    $project: {
      _id: 0,
      account_id: "$_id",
      customer_name: 1,
      credit_limit: 1,
      outstanding_debt: 1,
    },
  },
]);
```

## Query 4

- Liệt kê thông tin của 10 khách hàng có tổng số tiền gửi trên các tài khoản của họ là nhiều nhất.

```javascript
db.customer.aggregate([
  {
    $unwind: "$accounts",
  },
  {
    $match: {
      "accounts.account_type": "credit",
    },
  },
  {
    $group: {
      _id: "$_id",
      customer_name: {
        $first: "$fullname",
      },
      total_deposit: {
        $sum: "$accounts.balance",
      },
    },
  },
  {
    $sort: {
      total_deposit: -1,
    },
  },
  {
    $limit: 10,
  },
  {
    $project: {
      _id: 0,
      customer_name: 1,
      total_deposit: 1,
    },
  },
]);
```

## Query 5

Thực hiện các giao dịch thanh toán nợ cho tài khoản ghi nợ từ tài khoản tín dụng.

- Việc thực hiện giao dịch thanh toán tín dụng cho tài khoản ghi nợ từ tài khoản tín dụng cần thực hiên qua nhiều bước, không thể thực hiện trong 1 truy vấn duy nhất.
- Các truy vấn cần thực hiện để hoàn thành yêu cầu trên (Các bước này được thực hiện trong 1 function javascript để minh hoạ):

```javascript
function payCreditAccount(
  credit_account_number,
  debit_account_number,
  payment_amount,
  employee_id,
  customer_id,
) {
  // Tìm tài khoản tín dụng
  const creditAccount = db.customer.findOne(
    {
      "accounts.account_number": credit_account_number,
      "accounts.account_type": "credit",
    },
    { "accounts.$": 1 },
  );

  // Kiểm tra số dư tài khoản tín dụng có đủ để thực hiện giao dịch không
  if (
    creditAccount.accounts[0].balance - payment_amount <
    creditAccount.accounts[0].min_balance
  ) {
    throw new Error(
      "Số dư tài khoản tín dụng không đủ để thực hiện giao dịch.",
    );
  }

  // Tìm tài khoản ghi nợ
  const debitAccount = db.customer.findOne(
    {
      "accounts.account_number": debit_account_number,
      "accounts.account_type": "debit",
    },
    { "accounts.$": 1 },
  );

  // Cập nhật số dư tài khoản tiết kiệm
  db.customer.updateOne(
    {
      "accounts.account_number": credit_account_id,
      "accounts.account_type": "credit",
    },
    { $inc: { "accounts.$.balance": -payment_amount } },
  );

  // Cập nhật số nợ tồn đọng trong tài khoản ghi nợ
  db.customer.updateOne(
    {
      "accounts.account_id": debit_account_id,
      "accounts.account_type": "debit",
    },
    { $inc: { "accounts.$.outstanding_debt": -payment_amount } },
  );

  // Thêm giao dịch vào lịch sử giao dịch của tài khoản tín dụng (Tài khoản người gửi)
  db.customer.updateOne(
    {
      "accounts.account_id": credit_account_id,
      "accounts.account_type": "credit",
    },
    {
      $push: {
        "accounts.$.transactions": {
          amount: payment_amount,
          date: new Date(),
          type: "credit_payment",
        },
      },
    },
  );

  // Thêm giao dịch vào lịch sử giao dịch của tài khoản ghi nợ
  db.customer.updateOne(
    {
      "accounts.account_id": debit_account_id,
      "accounts.account_type": "debit",
    },
    {
      $push: {
        "accounts.$.transactions": {
          amount: payment_amount,
          date: new Date(),
          type: "credit_payment",
        },
      },
    },
  );

  // Tạo lịch sử giao dịch lưu vào collection transaction
  db.transaction.insert({
    _id: ObjectId("59fef89dde1deb9eba370ef2"),
    transaction_id: "99f6d391-39fc-436e-a5d2-e07c17f64b35",
    account_number: "134070496",
    customer_id: "baa69552-cba6-4d1c-978f-1089c6b6d948",
    employee_id: "2008839e-cca8-49a9-bd8f-6fa38d49574f",
    amount: 414.58,
    date: "2024-11-19T04:52:09.526Z",
    type: "transfer",
  });

  return "Thanh toán thành công!";
}
```
