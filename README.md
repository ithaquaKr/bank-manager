# BANK-MANAGER

**Real-world scenario: Building a branch bank management system.**

**Database requirements include:**

- Bank employees have the following basic information: employee ID, identity card number, name, date of birth, address, job rank, seniority, and job position.

- Basic customer information includes: identity card number, customer ID, name, date of birth, and address.

- Information about bank accounts includes: account ID, account type, and balance.

- Each customer can open up to 2 credit accounts and 3 savings accounts. Credit accounts have additional information such as credit limits. Savings accounts have additional information such as monthly interest rates and minimum balance.

- A customer can perform multiple transactions on each credit account, but the total debt on each account (including ongoing transactions) cannot exceed the credit limit. Whenever a customer makes a payment on a credit account, the debt on that account is reduced by the payment amount.

- Customers can pay off credit accounts using funds from their savings accounts. Customers can also pay off other people's credit accounts. The payment amount for each credit account will be deducted from the balance of the savings account. However, the transaction can only proceed if the remaining balance in the savings account is greater than or equal to the minimum balance.

- Bank transactions can be supported by any bank employee. The process of creating customer accounts is carried out only by sales staff.

Query requirements:

    1. Calculate the salaries of sales employees based on the number of bank accounts they have created for customers each month. Each created credit account will earn an additional 500,000 VND, and each savings account will earn 2% of the initial deposit made by the customer.
    2. List the names of customers along with the transaction amount for each credit account they made within the period from the start date to the end date.
    3. List information of credit accounts along with the total outstanding debt at the time of the query, sorted in descending order by debt balance.
    4. List information of the 10 customers with the highest total deposits in their accounts.
    5. Execute credit account payment transactions from savings accounts. Ensure that the database constraints are reflected in the application.
