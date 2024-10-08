# Query

## Query 1

- Query 1: Tính lương của các nhân viên kinh doanh dựa trên số tài khoản ngân hàng mà họ đã tạo được cho khách hàng trong mỗi tháng. Mỗi tài khoản ghi nợ (debit) được tạo ra sẽ được cộng 500 nghìn, mỗi tài khoản tín dụng (credit) sẽ được cộng 2% số tiền mà khách gửi lần đầu.

```javascript
db.employees.aggregate([
  {
    $match: {
      position: "sale",
    },
  },
  {
    $unwind: "$accounts_created",
  },
  {
    $addFields: {
      "accounts_created.initial_deposit": {
        $convert: {
          input: "$accounts_created.initial_deposit",
          to: "double",
          onError: 0,
        },
      },
    },
  },
  {
    $group: {
      _id: {
        employee_id: "$employee_id",
        employee_name: "$fullname",
        month: {
          $dateToString: {
            format: "%Y-%m",
            date: {
              $toDate: "$accounts_created.created_at",
            },
          },
        },
      },
      total_salary: {
        $sum: {
          $cond: [
            {
              $eq: ["$accounts_created.type", "credit"],
            },
            500,
            {
              $multiply: ["$accounts_created.initial_deposit", 0.02],
            },
          ],
        },
      },
    },
  },
  {
    $project: {
      employee_id: "$_id.employee_id",
      employee_name: "$_id.employee_name",
      month: "$_id.month",
      total_salary: {
        $round: ["$total_salary", 3],
      },
    },
  },
]);
```

## Query 2

- Query 2: Liệt kê tên khách hàng cùng số tiền giao dịch trên mỗi tài khoản tín dụng (credit) mà họ thực hiện trong khoảng thời gian từ ngày bắt đầu đến ngày kết thúc.

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
    $unwind: "$accounts.transactions",
  },
  {
    $match: {
      "accounts.transactions.created_at": {
        $gte: new Date("2022-07-01").toISOString(),
        $lte: new Date("2024-11-31").toISOString(),
      },
    },
  },
  {
    $project: {
      _id: 0,
      customer_name: "$fullname",
      transaction_amount: "$accounts.transactions.amount",
      transaction_date: "$accounts.transactions.created_at",
    },
  },
]);
```

## Query 3

- Query 3: Liệt kê thông tin các tài khoản tín dụng (credit) cùng tổng số nợ tồn đọng tại thời điểm truy vấn, danh sách được sắp xếp theo thứ tự giảm dần số dự nợ.

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
    $addFields: {
      "accounts.total_outstanding_debt": {
        $convert: {
          input: "$accounts.outstanding_balance",
          to: "double",
          onError: 0,
          onNull: 0,
        },
      },
    },
  },
  {
    $sort: {
      "accounts.total_outstanding_debt": -1,
    },
  },
  {
    $project: {
      _id: 0,
      customer_id: 1,
      fullname: 1,
      account_number: "$accounts.account_number",
      account_name: "$accounts.account_name",
      total_outstanding_debt: "$accounts.total_outstanding_debt",
    },
  },
]);
```

## Query 4

- Query 4: Liệt kê thông tin của 10 khách hàng có tổng số tiền gửi trên các tài khoản của họ là nhiều nhất.

```javascript
db.customer.aggregate([
  {
    $unwind: "$accounts",
  },
  {
    $match: {
      "accounts.account_type": "debit",
    },
  },
  {
    $group: {
      _id: "$_id",
      fullname: { $first: "$fullname" },
      customer_id: { $first: "$customer_id" },
      total_balance: { $sum: { $toDouble: "$accounts.balance" } },
    },
  },
  {
    $project: {
      _id: 1,
      fullname: 1,
      customer_id: 1,
      total_balance: { $round: ["$total_balance", 3] },
    },
  },
  {
    $sort: { total_balance: -1 },
  },
  {
    $limit: 10,
  },
]);
```

## Query 5

Thực hiện các giao dịch thanh toán nợ cho tài khoản tín dụng từ tài khoản ghi nợ.

- Việc thực hiện giao dịch thanh toán tín dụng cho tài khoản tín dụng từ tài khoản ghi nợ cần thực hiên qua nhiều bước, không thể thực hiện trong 1 truy vấn duy nhất.
- Các truy vấn cần thực hiện để hoàn thành yêu cầu trên (Các bước này được thực hiện trong 1 function javascript để minh hoạ):

```javascript
function payCreditAccount(
  credit_account_number,
  debit_account_number,
  payment_amount,
  employee_id,
  customer_id,
) {
  // Tìm tài khoản ghi nợ
  const debitAccount = db.customer.findOne(
    {
      "accounts.account_number": debit_account_number,
      "accounts.account_type": "debit",
    },
    { "accounts.$": 1 },
  );

  // Kiểm tra số dư tài khoản ghi nợ có đủ để thực hiện giao dịch không
  if (
    debitAccount.accounts[0].balance - payment_amount <
    debitAccount.accounts[0].min_balance
  ) {
    throw new Error("Số dư tài khoản không đủ để thực hiện giao dịch.");
  }

  // Tìm tài khoản tín dụng
  const creditAccount = db.customer.findOne(
    {
      "accounts.account_number": credit_account_number,
      "accounts.account_type": "credit",
    },
    { "accounts.$": 1 },
  );

  // Cập nhật số dư tài khoản ghi nọ, thêm lịch sử giao dịch
  db.customer.updateOne(
    {
      "accounts.account_number": debit_account_number,
      "accounts.account_type": "debit",
    },
    { $inc: { "accounts.$.balance": -payment_amount } },
    {
      $push: {
        "accounts.$.transactions": {
          transaction_id: "3344cfde-10fe-4bbf-9a99-054cf4387619", // Generate using uuid() function
          amount: -payment_amount,
          date: new Date(), // Format ISODate
        },
      },
    },
  );

  // Cập nhật số nợ tồn đọng trong tài khoản tín dụng
  db.customer.updateOne(
    {
      "accounts.account_number": credit_account_number,
      "accounts.account_type": "credit",
    },
    { $inc: { "accounts.$.outstanding_balance": -payment_amount } },
    {
      $push: {
        "accounts.$.transactions": {
          transaction_id: "2b3f977d-de79-45e6-b64e-df8e3556ee65", // Generate using uuid() function
          amount: -payment_amount,
          date: new Date(), // Format ISODate
        },
      },
    },
  );

  // Tạo lịch sử giao dịch lưu vào collection transaction
  db.transaction.insertMany(
    {
      transaction_id: "3344cfde-10fe-4bbf-9a99-054cf4387619", // Same ID with transaction in account debit
      account_number: debit_account_number,
      customer_id: customer_id,
      employee_id: employee_id,
      amount: -payment_amount,
      date: "2024-11-19T04:52:09.526Z",
      type: "transfer",
    },
    {
      transaction_id: "2b3f977d-de79-45e6-b64e-df8e3556ee65", // Same ID with transaction in account credit
      account_number: credit_account_number,
      customer_id: customer_id,
      employee_id: employee_id,
      amount: -payment_amount,
      date: "2024-11-19T04:52:09.526Z",
      type: "transfer",
    },
  );

  return "Thanh toán thành công!";
}
```
