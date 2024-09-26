# DATABASE DESIGN

## Input design

**Real-world scenario: Building a branch bank management system.**

**Database requirements include:**

- Bank employees have the following basic information: employee ID, identity card number, name, date of birth, address, job rank, seniority, and job position.

- Basic customer information includes: identity card number, customer ID, name, date of birth, and address.

- Information about bank accounts includes: account ID, account type, and balance.

- Each customer can open up to 2 credit accounts and 3 savings accounts. Credit accounts have additional information such as credit limits. Savings accounts have additional information such as monthly interest rates and minimum balance.

- A customer can perform multiple transactions on each credit account, but the total debt on each account (including ongoing transactions) cannot exceed the credit limit. Whenever a customer makes a payment on a credit account, the debt on that account is reduced by the payment amount.

- Customers can pay off credit accounts using funds from their savings accounts. Customers can also pay off other people's credit accounts. The payment amount for each credit account will be deducted from the balance of the savings account. However, the transaction can only proceed if the remaining balance in the savings account is greater than or equal to the minimum balance.

- Bank transactions can be supported by any bank employee. The process of creating customer accounts is carried out only by sales staff.

## Design

### Employee

```go
type Employee struct {
 EmployeeID       string `bson:"employee_id"`
 IdentityCard     string `bson:"identity_card"`
 Name             string `bson:"name"`
 DateOfBirth      string `bson:"date_of_birth"`
 Address          string `bson:"address"`
 JobRank          string `bson:"job_rank"`
 Seniority        int    `bson:"seniority"`
 JobPosition      string `bson:"job_position"`
 IsSalesStaff     bool   `bson:"is_sales_staff"`
}
```

### Customer

```go
type Customer struct {
 IdentityCard string   `bson:"identity_card"`
 CustomerID   string   `bson:"customer_id"`
 Name         string   `bson:"name"`
 DateOfBirth  string   `bson:"date_of_birth"`
 Address      string   `bson:"address"`
 Accounts     []string `bson:"accounts"`
}
```

### Account

```go
type Account struct {
 AccountID   string  `bson:"account_id"`
 Type        string  `bson:"type"`
 Balance     float64 `bson:"balance"`
 CustomerID  string  `bson:"customer_id"`
 CreditLimit *float64 `bson:"credit_limit,omitempty"`
 InterestRate *float64 `bson:"interest_rate,omitempty"`
 MinBalance  *float64 `bson:"min_balance,omitempty"`
}
```

### Transaction

```go
type Transaction struct {
 TransactionID string  `bson:"transaction_id"`
 AccountID     string  `bson:"account_id"`
 Amount        float64 `bson:"amount"`
 Type         string  `bson:"type"` // "deposit", "withdraw", "payment"
 EmployeeID    string  `bson:"employee_id"`
 Date          string  `bson:"date"`
 Description   string  `bson:"description"`
 SourceAccount *string `bson:"source_account,omitempty"`
}
```

### Explanation

- **Employee:** Stores basic employee information, including a flag `IsSalesStaff` to identify employees who can create customer accounts.
- **Customer:** Stores basic customer information and an array `Accounts` to store the IDs of the accounts associated with the customer.
- **Account:** Stores information about each bank account, including account type, balance, and customer ID. For credit accounts, it includes `CreditLimit`. For savings accounts, it includes `InterestRate` and `MinBalance`.
- **Transaction:** Stores details of each transaction, including the account ID, amount, type, employee ID, date, description, and the source account for payment transactions.

### Relationships

- **One-to-Many:**
  - An employee can handle multiple transactions.
  - A customer can have multiple accounts.
  - An account can have multiple transactions.
- **Many-to-Many:**
  - A customer can have multiple accounts.
  - A customer can make payments from their savings account to multiple credit accounts.

### Constraints

- A customer can have up to 2 credit accounts and 3 savings accounts.
- The total debt on a credit account, including ongoing transactions, cannot exceed the credit limit.
- A transaction can only proceed if the remaining balance in the savings account is greater than or equal to the minimum balance.

### Considerations

- The `Transaction` model can be extended to include additional information, such as transaction fees, transaction status, and transaction notes.
- The database can be further optimized by using indexes to improve query performance.
- You can add validation logic to ensure that transactions comply with the defined constraints.

This database model provides a foundation for building a branch bank management system. It can be further customized and expanded to meet the specific needs of your application.
