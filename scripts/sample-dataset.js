// Import dependencies
const { MongoClient } = require("mongodb");
const { faker } = require("@faker-js/faker");

// MongoDB URL
const uri = "mongodb://admin:secretpassword@127.0.0.1:27017";
const dbName = "bank-manager";

function getRandomNumberInRange(min, max) {
  return Math.floor(Math.random() * (max - min + 1)) + min;
}

async function createDateset() {
  const client = new MongoClient(uri);

  try {
    await client.connect();
    console.log("Connected to MongoDB server");
    // Start a session
    const session = client.startSession();
    // Start a transaction
    session.startTransaction();

    const db = client.db(dbName);
    const employeeCollection = db.collection("employee");
    const customerCollection = db.collection("customer");
    const transactionCollection = db.collection("transaction");

    // Create Employee data
    const employees = [];
    const total_accounts_created = []; // Store all account has been created
    for (let i = 0; i < 10; i++) {
      // Create random accounts_created for employee
      const accounts_created = [];
      const number_account = getRandomNumberInRange(2, 6);
      for (let i = 0; i < number_account; i++) {
        account = {
          account_number: faker.finance.accountNumber({ length: 9 }),
          type: faker.helpers.arrayElement(["credit", "debit"]),
          initial_deposit: faker.finance.amount(100, 1000, 2),
          created_at: faker.date.recent().toISOString(),
        };
        accounts_created.push(account);
        total_accounts_created.push(account);
      }

      const employee = {
        employee_id: faker.string.uuid(),
        identity_card: faker.commerce.isbn(),
        fullname: faker.person.fullName(),
        date_of_birth: faker.date
          .birthdate({ mode: "year", min: 1900, max: 2005 })
          .toISOString(),
        address: faker.location.streetAddress({ useFullAddress: true }),
        job_rank: faker.helpers.arrayElement(["entry", "middle", "senior"]),
        position: faker.helpers.arrayElement(["sale", "marketing"]),
        seniority: faker.number.int({ min: 0, max: 30 }),
        accounts_created: accounts_created,
      };

      employees.push(employee);
    }

    const customers = [];
    const transactions = [];
    while (total_accounts_created.length > 0) {
      const account_customer_number = getRandomNumberInRange(1, 3);
      const account_customers = [];
      const customer_id = faker.string.uuid();

      for (let i = 0; i < account_customer_number; i++) {
        const acc = total_accounts_created.pop();
        const account_customer = {
          account_number: acc.account_number,
          account_name: faker.finance.accountName(),
          account_type: acc.type,
        };

        if (account_customer.account_type === "debit") {
          account_customer.balance = faker.finance.amount(50, 1000, 2);
          (account_customer.min_balance = faker.finance.amount(0, 50, 2)),
            (account_customer.monthly_interest_rate = faker.finance.amount(
              0,
              10,
              2,
            ));
        } else {
          account_customer.credit_limit = faker.finance.amount(100, 1000, 2);
          account_customer.outstanding_balance = faker.finance.amount(
            0,
            100,
            2,
          );
        }

        const number_transaction = getRandomNumberInRange(10, 20);
        const customer_transactions = [];

        for (let i = 0; i < number_transaction; i++) {
          const transaction_customer = {
            transaction_id: faker.string.uuid(),
            amount: faker.finance.amount(-1000, 1000, 2),
            created_at: faker.date.recent().toISOString(),
          };
          customer_transactions.push(transaction_customer);

          // Add more info
          transaction_customer.account_number = account_customer.account_number;
          transaction_customer.employee_id = faker.helpers.arrayElement(
            employees.filter((employee) => employee.position === "sale"),
          ).employee_id;
          transaction_customer.customer_id;
          transaction_customer.type = faker.helpers.arrayElement([
            "deposit",
            "withdrawal",
            "transfer",
            "loan",
            "credit_payment",
            "bill_payment",
            "fee",
            "investment",
          ]);
          transactions.push(transaction_customer);
        }
        account_customer.transactions = customer_transactions;
        account_customers.push(account_customer);
      }

      const customer = {
        customer_id: customer_id,
        identity_card: faker.commerce.isbn(),
        fullname: faker.person.fullName(),
        date_of_birth: faker.date
          .birthdate({ mode: "year", min: 1900, max: 2005 })
          .toISOString(),
        address: faker.location.streetAddress({ useFullAddress: true }),
        accounts: account_customers,
      };
      customers.push(customer);
    }

    try {
      await employeeCollection.insertMany(employees);
      console.log("Employee data inserted");
      await customerCollection.insertMany(customers);
      console.log("Customer data inserted");
      await transactionCollection.insertMany(transactions);
      console.log("Transaction data inserted");
      // Commit the transaction if all operations are successful
      await session.commitTransaction();
      console.log("Transaction committed successfully!");
    } catch (error) {
      await session.abortTransaction();
      console.error("Transaction aborted due to error:", error);
    } finally {
      session.endSession();
    }
  } catch (error) {
    console.log(error);
  } finally {
    await client.close();
  }
}

createDateset();
