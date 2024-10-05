#!/bin/bash

# MongoDB connection details
MONGO_USER="ithadevnguyen"
MONGO_PASSWORD=""

# Database and collection names
DATABASE_NAME="bank-manager"

# Sample Employee data
mongosh "mongodb+srv://bank-manager.z3mb2.mongodb.net/" --apiVersion 1 --username $MONGO_USER --password $MONGO_PASSWORD <<EOF
use $DATABASE_NAME

db.employee.insertMany([
  {
    "_id": ObjectId("64e243268a041259a0264255"),
    "employee_id": "e4308544-9a03-4e4f-b64b-372645e53d08",
    "identity_card": "978-0-487-32207-0",
    "fullname": "Nguyễn Văn A",
    "date_of_birth": "1990-01-15T00:00:00.000Z",
    "address": "123 Lê Duẩn, Hanoi, Vietnam",
    "job_rank": "senior",
    "position": "sales",
    "seniority": 5,
    "accounts_created": [
      {
        "account_number": "1234567890",
        "type": "credit",
        "initial_deposit": 1000,
        "created_at": "2023-04-24T12:53:54.853Z"
      },
      {
        "account_number": "9876543210",
        "type": "debit",
        "initial_deposit": 500,
        "created_at": "2023-05-10T12:53:54.853Z"
      },
      {
        "account_number": "1112223333",
        "type": "credit",
        "initial_deposit": 2000,
        "created_at": "2023-06-15T12:53:54.853Z"
      }
    ]
  },
  {
    "_id": ObjectId("64e243268a041259a0264256"),
    "employee_id": "d7825a5f-5265-466a-b8c5-5721a3b38848",
    "identity_card": "978-0-487-32208-1",
    "fullname": "Trần Thị B",
    "date_of_birth": "1985-07-20T00:00:00.000Z",
    "address": "456 Nguyễn Trãi, Hanoi, Vietnam",
    "job_rank": "junior",
    "position": "marketing",
    "seniority": 2,
    "accounts_created": [
      {
        "account_number": "4567890123",
        "type": "debit",
        "initial_deposit": 750,
        "created_at": "2023-07-01T12:53:54.853Z"
      },
      {
        "account_number": "3210987654",
        "type": "credit",
        "initial_deposit": 1500,
        "created_at": "2023-08-10T12:53:54.853Z"
      },
      {
        "account_number": "7890123456",
        "type": "debit",
        "initial_deposit": 1000,
        "created_at": "2023-09-20T12:53:54.853Z"
      },
      {
        "account_number": "0123456789",
        "type": "credit",
        "initial_deposit": 2500,
        "created_at": "2023-10-05T12:53:54.853Z"
      }
    ]
  },
  {
    "_id": ObjectId("64e243268a041259a0264257"),
    "employee_id": "f873a87b-478e-47b7-a15c-1412d0b8e244",
    "identity_card": "978-0-487-32209-2",
    "fullname": "Lê Hồng C",
    "date_of_birth": "1992-03-08T00:00:00.000Z",
    "address": "789 Hoàng Diệu, Hanoi, Vietnam",
    "job_rank": "mid",
    "position": "management",
    "seniority": 3,
    "accounts_created": [
      {
        "account_number": "2345678901",
        "type": "credit",
        "initial_deposit": 1200,
        "created_at": "2023-10-20T12:53:54.853Z"
      },
      {
        "account_number": "5678901234",
        "type": "debit",
        "initial_deposit": 800,
        "created_at": "2023-11-01T12:53:54.853Z"
      },
      {
        "account_number": "8901234567",
        "type": "credit",
        "initial_deposit": 1800,
        "created_at": "2023-12-10T12:53:54.853Z"
      },
      {
        "account_number": "0123456789",
        "type": "debit",
        "initial_deposit": 1500,
        "created_at": "2024-01-05T12:53:54.853Z"
      },
      {
        "account_number": "9012345678",
        "type": "credit",
        "initial_deposit": 2200,
        "created_at": "2024-02-15T12:53:54.853Z"
      }
    ]
  },
  {
    "_id": ObjectId("64e243268a041259a0264258"),
    "employee_id": "21a04e97-a921-4664-b79d-8f3163c70d96",
    "identity_card": "978-0-487-32210-8",
    "fullname": "Đặng Minh D",
    "date_of_birth": "1988-05-05T00:00:00.000Z",
    "address": "1011 Trần Hưng Đạo, Hanoi, Vietnam",
    "job_rank": "junior",
    "position": "sales",
    "seniority": 1,
    "accounts_created": [
      {
        "account_number": "3456789012",
        "type": "debit",
        "initial_deposit": 900,
        "created_at": "2024-03-01T12:53:54.853Z"
      },
      {
        "account_number": "6789012345",
        "type": "credit",
        "initial_deposit": 1700,
        "created_at": "2024-03-15T12:53:54.853Z"
      },
      {
        "account_number": "9012345678",
        "type": "debit",
        "initial_deposit": 1300,
        "created_at": "2024-04-10T12:53:54.853Z"
      },
      {
        "account_number": "1234567890",
        "type": "credit",
        "initial_deposit": 2000,
        "created_at": "2024-04-24T12:53:54.853Z"
      },
      {
        "account_number": "4567890123",
        "type": "debit",
        "initial_deposit": 1600,
        "created_at": "2024-05-05T12:53:54.853Z"
      },
      {
        "account_number": "7890123456",
        "type": "credit",
        "initial_deposit": 2500,
        "created_at": "2024-05-20T12:53:54.853Z"
      }
    ]
  }
]);
EOF

# Sample Customer data
mongosh "mongodb+srv://bank-manager.z3mb2.mongodb.net/" --apiVersion 1 --username $MONGO_USER --password $MONGO_PASSWORD <<EOF
use $DATABASE_NAME

db.customer.insertMany([
  {
    "_id": ObjectId("6504896395e32c653869831c"),
    "customer_id": "6270228c-9d4c-4059-b3f7-15b159890552",
    "identity_card": "299-0-8785-6769-0",
    "fullname": "Nguyễn Văn A",
    "date_of_birth": "1990-03-15T04:16:48.554Z",
    "address": "123 Đường Láng, Hà Nội, Việt Nam",
    "accounts": [
      {
        "account_number": "1234567890",
        "account_name": "Savings Account",
        "account_type": "debit",
        "balance": 100,
        "debit_limit": 500,
        "outstanding_debt": 0,
        "transactions": [
          {
            "transaction_id": "a1b2c3d4-e5f6-7890-1234-567890abcdef",
            "amount": 50,
            "date": "2023-10-26T07:10:35.306Z"
          },
          {
            "transaction_id": "f7890123-4567-8901-2345-67890abcdef",
            "amount": 100,
            "date": "2023-10-27T07:10:35.306Z"
          },
          {
            "transaction_id": "c3d4e5f6-7890-1234-5678-90abcdef12",
            "amount": 200,
            "date": "2023-10-28T07:10:35.306Z"
          }
        ]
      },
      {
        "account_number": "9876543210",
        "account_name": "Checking Account",
        "account_type": "credit",
        "balance": 100,
        "min_balance": 50,
        "monthly_interest_rate": 0.05,
        "transactions": [
          {
            "transaction_id": "f1e2d3c4-b5a6-9780-1234-567890abcdef",
            "amount": 100,
            "date": "2023-10-26T07:10:35.306Z"
          },
          {
            "transaction_id": "c9876543-210f-e1d2-3456-7890abcdef",
            "amount": 150,
            "date": "2023-10-27T07:10:35.306Z"
          },
          {
            "transaction_id": "b4c5d6e7-8901-2345-6789-0abcdef12",
            "amount": 200,
            "date": "2023-10-28T07:10:35.306Z"
          },
          {
            "transaction_id": "a5b6c7d8-9012-3456-7890-abcdef12",
            "amount": 250,
            "date": "2023-10-29T07:10:35.306Z"
          }
        ]
      }
    ]
  },
  {
    "_id": ObjectId("6504896395e32c653869831d"),
    "customer_id": "9c636f56-017a-46d9-9562-9269b65b673c",
    "identity_card": "299-0-8785-6769-0",
    "fullname": "Trần Thị B",
    "date_of_birth": "1992-05-20T04:16:48.554Z",
    "address": "123 Đường Nguyễn Trãi, Hà Nội, Việt Nam",
    "accounts": [
      {
        "account_number": "9876543210",
        "account_name": "Checking Account",
        "account_type": "credit",
        "balance": 100,
        "min_balance": 50,
        "monthly_interest_rate": 0.05,
        "transactions": [
          {
            "transaction_id": "f1e2d3c4-b5a6-9780-1234-567890abcdef",
            "amount": 100,
            "date": "2023-10-26T07:10:35.306Z"
          },
          {
            "transaction_id": "c9876543-210f-e1d2-3456-7890abcdef",
            "amount": 150,
            "date": "2023-10-27T07:10:35.306Z"
          },
          {
            "transaction_id": "b4c5d6e7-8901-2345-6789-0abcdef12",
            "amount": 200,
            "date": "2023-10-28T07:10:35.306Z"
          },
          {
            "transaction_id": "a5b6c7d8-9012-3456-7890-abcdef12",
            "amount": 250,
            "date": "2023-10-29T07:10:35.306Z"
          }
        ]
      },
      {
        "account_number": "0123456789",
        "account_name": "Savings Account",
        "account_type": "debit",
        "balance": 100,
        "debit_limit": 500,
        "outstanding_debt": 0,
        "transactions": [
          {
            "transaction_id": "a1b2c3d4-e5f6-7890-1234-567890abcdef",
            "amount": 50,
            "date": "2023-10-26T07:10:35.306Z"
          },
          {
            "transaction_id": "f7890123-4567-8901-2345-67890abcdef",
            "amount": 100,
            "date": "2023-10-27T07:10:35.306Z"
          },
          {
            "transaction_id": "c3d4e5f6-7890-1234-5678-90abcdef12",
            "amount": 200,
            "date": "2023-10-28T07:10:35.306Z"
          }
        ]
      }
    ]
  },
  {
    "_id": ObjectId("6504896395e32c653869831e"),
    "customer_id": "86e47b58-3739-438c-a9a9-49730589700d",
    "identity_card": "599-0-6659-5868-7",
    "fullname": "Lê Văn C",
    "date_of_birth": "1993-07-10T04:16:48.554Z",
    "address": "123 Đường Giải Phóng, Hà Nội, Việt Nam",
    "accounts": [
      {
        "account_number": "1011121314",
        "account_name": "Savings Account",
        "account_type": "debit",
        "balance": 100,
        "debit_limit": 500,
        "outstanding_debt": 0,
        "transactions": [
          {
            "transaction_id": "a1b2c3d4-e5f6-7890-1234-567890abcdef",
            "amount": 50,
            "date": "2023-10-26T07:10:35.306Z"
          },
          {
            "transaction_id": "f7890123-4567-8901-2345-67890abcdef",
            "amount": 100,
            "date": "2023-10-27T07:10:35.306Z"
          },
          {
            "transaction_id": "c3d4e5f6-7890-1234-5678-90abcdef12",
            "amount": 200,
            "date": "2023-10-28T07:10:35.306Z"
          }
        ]
      },
      {
        "account_number": "1516171819",
        "account_name": "Checking Account",
        "account_type": "credit",
        "balance": 100,
        "min_balance": 50,
        "monthly_interest_rate": 0.05,
        "transactions": [
          {
            "transaction_id": "f1e2d3c4-b5a6-9780-1234-567890abcdef",
            "amount": 100,
            "date": "2023-10-26T07:10:35.306Z"
          },
          {
            "transaction_id": "c9876543-210f-e1d2-3456-7890abcdef",
            "amount": 150,
            "date": "2023-10-27T07:10:35.306Z"
          },
          {
            "transaction_id": "b4c5d6e7-8901-2345-6789-0abcdef12",
            "amount": 200,
            "date": "2023-10-28T07:10:35.306Z"
          },
          {
            "transaction_id": "a5b6c7d8-9012-3456-7890-abcdef12",
            "amount": 250,
            "date": "2023-10-29T07:10:35.306Z"
          }
        ]
      },
      {
        "account_number": "2021222324",
        "account_name": "Credit Card",
        "account_type": "credit",
        "balance": 100,
        "min_balance": 50,
        "monthly_interest_rate": 0.05,
        "transactions": [
          {
            "transaction_id": "f1e2d3c4-b5a6-9780-1234-567890abcdef",
            "amount": 100,
            "date": "2023-10-26T07:10:35.306Z"
          },
          {
            "transaction_id": "c9876543-210f-e1d2-3456-7890abcdef",
            "amount": 150,
            "date": "2023-10-27T07:10:35.306Z"
          },
          {
            "transaction_id": "b4c5d6e7-8901-2345-6789-0abcdef12",
            "amount": 200,
            "date": "2023-10-28T07:10:35.306Z"
          },
          {
            "transaction_id": "a5b6c7d8-9012-3456-7890-abcdef12",
            "amount": 250,
            "date": "2023-10-29T07:10:35.306Z"
          }
        ]
      }
    ]
  },
  {
    "_id": ObjectId("6504896395e32c653869831f"),
    "customer_id": "6270228c-9d4c-4059-b3f7-15b159890553",
    "identity_card": "499-0-9766-1979-8",
    "fullname": "Phạm Thị D",
    "date_of_birth": "1994-09-05T04:16:48.554Z",
    "address": "123 Đường Tây Sơn, Hà Nội, Việt Nam",
    "accounts": [
      {
        "account_number": "2526272829",
        "account_name": "Savings Account",
        "account_type": "debit",
        "balance": 100,
        "debit_limit": 500,
        "outstanding_debt": 0,
        "transactions": [
          {
            "transaction_id": "a1b2c3d4-e5f6-7890-1234-567890abcdef",
            "amount": 50,
            "date": "2023-10-26T07:10:35.306Z"
          },
          {
            "transaction_id": "f7890123-4567-8901-2345-67890abcdef",
            "amount": 100,
            "date": "2023-10-27T07:10:35.306Z"
          },
          {
            "transaction_id": "c3d4e5f6-7890-1234-5678-90abcdef12",
            "amount": 200,
            "date": "2023-10-28T07:10:35.306Z"
          }
        ]
      },
      {
        "account_number": "3031323334",
        "account_name": "Checking Account",
        "account_type": "credit",
        "balance": 100,
        "min_balance": 50,
        "monthly_interest_rate": 0.05,
        "transactions": [
          {
            "transaction_id": "f1e2d3c4-b5a6-9780-1234-567890abcdef",
            "amount": 100,
            "date": "2023-10-26T07:10:35.306Z"
          },
          {
            "transaction_id": "c9876543-210f-e1d2-3456-7890abcdef",
            "amount": 150,
            "date": "2023-10-27T07:10:35.306Z"
          },
          {
            "transaction_id": "b4c5d6e7-8901-2345-6789-0abcdef12",
            "amount": 200,
            "date": "2023-10-28T07:10:35.306Z"
          },
          {
            "transaction_id": "a5b6c7d8-9012-3456-7890-abcdef12",
            "amount": 250,
            "date": "2023-10-29T07:10:35.306Z"
          }
        ]
      },
      {
        "account_number": "3536373839",
        "account_name": "Credit Card",
        "account_type": "credit",
        "balance": 100,
        "min_balance": 50,
        "monthly_interest_rate": 0.05,
        "transactions": [
          {
            "transaction_id": "f1e2d3c4-b5a6-9780-1234-567890abcdef",
            "amount": 100,
            "date": "2023-10-26T07:10:35.306Z"
          },
          {
            "transaction_id": "c9876543-210f-e1d2-3456-7890abcdef",
            "amount": 150,
            "date": "2023-10-27T07:10:35.306Z"
          },
          {
            "transaction_id": "b4c5d6e7-8901-2345-6789-0abcdef12",
            "amount": 200,
            "date": "2023-10-28T07:10:35.306Z"
          },
          {
            "transaction_id": "a5b6c7d8-9012-3456-7890-abcdef12",
            "amount": 250,
            "date": "2023-10-29T07:10:35.306Z"
          }
        ]
      }
    ]
  },
  {
    "_id": ObjectId("6504896395e32c6538698320"),
    "customer_id": "051a3d2b-f9b2-4744-8d5d-5f55239c2044",
    "identity_card": "299-0-8785-6769-0",
    "fullname": "Vũ Thị E",
    "date_of_birth": "1995-11-12T04:16:48.554Z",
    "address": "123 Đường Hoàng Quốc Việt, Hà Nội, Việt Nam",
    "accounts": [
      {
        "account_number": "4041424344",
        "account_name": "Savings Account",
        "account_type": "debit",
        "balance": 100,
        "debit_limit": 500,
        "outstanding_debt": 0,
        "transactions": [
          {
            "transaction_id": "a1b2c3d4-e5f6-7890-1234-567890abcdef",
            "amount": 50,
            "date": "2023-10-26T07:10:35.306Z"
          },
          {
            "transaction_id": "f7890123-4567-8901-2345-67890abcdef",
            "amount": 100,
            "date": "2023-10-27T07:10:35.306Z"
          },
          {
            "transaction_id": "c3d4e5f6-7890-1234-5678-90abcdef12",
            "amount": 200,
            "date": "2023-10-28T07:10:35.306Z"
          }
        ]
      },
      {
        "account_number": "4546474849",
        "account_name": "Checking Account",
        "account_type": "credit",
        "balance": 100,
        "min_balance": 50,
        "monthly_interest_rate": 0.05,
        "transactions": [
          {
            "transaction_id": "f1e2d3c4-b5a6-9780-1234-567890abcdef",
            "amount": 100,
            "date": "2023-10-26T07:10:35.306Z"
          },
          {
            "transaction_id": "c9876543-210f-e1d2-3456-7890abcdef",
            "amount": 150,
            "date": "2023-10-27T07:10:35.306Z"
          },
          {
            "transaction_id": "b4c5d6e7-8901-2345-6789-0abcdef12",
            "amount": 200,
            "date": "2023-10-28T07:10:35.306Z"
          },
          {
            "transaction_id": "a5b6c7d8-9012-3456-7890-abcdef12",
            "amount": 250,
            "date": "2023-10-29T07:10:35.306Z"
          }
        ]
      },
      {
        "account_number": "5051525354",
        "account_name": "Credit Card",
        "account_type": "credit",
        "balance": 100,
        "min_balance": 50,
        "monthly_interest_rate": 0.05,
        "transactions": [
          {
            "transaction_id": "f1e2d3c4-b5a6-9780-1234-567890abcdef",
            "amount": 100,
            "date": "2023-10-26T07:10:35.306Z"
          },
          {
            "transaction_id": "c9876543-210f-e1d2-3456-7890abcdef",
            "amount": 150,
            "date": "2023-10-27T07:10:35.306Z"
          },
          {
            "transaction_id": "b4c5d6e7-8901-2345-6789-0abcdef12",
            "amount": 200,
            "date": "2023-10-28T07:10:35.306Z"
          },
          {
            "transaction_id": "a5b6c7d8-9012-3456-7890-abcdef12",
            "amount": 250,
            "date": "2023-10-29T07:10:35.306Z"
          }
        ]
      }
    ]
  }
]);
EOF

# Sample Employee data
mongosh "mongodb+srv://bank-manager.z3mb2.mongodb.net/" --apiVersion 1 --username $MONGO_USER --password $MONGO_PASSWORD <<EOF
use $DATABASE_NAME

db.transaction.insertMany([
  {
    "_id": ObjectId("646d4551405d0301468a5854"),
    "transaction_id": "7521d257-a26f-42f9-9818-c034e8115d17",
    "account_number": "424293682",
    "customer_id": "49654291-2737-4a5d-9516-4c2887922a11",
    "employee_id": "a853a240-f965-44a9-8c73-a58754136743",
    "amount": 965.96,
    "date": "2024-05-23T15:19:04.165Z",
    "type": "credit_payment"
  },
  {
    "_id": ObjectId("646d4551405d0301468a5855"),
    "transaction_id": "2a095785-2572-4412-9923-848750638455",
    "account_number": "658142484",
    "customer_id": "e70c8323-4811-4384-a279-981249896413",
    "employee_id": "7874325b-4774-40c5-94d9-74100684a589",
    "amount": 633.35,
    "date": "2024-05-23T10:26:49.818Z",
    "type": "debit"
  },
  {
    "_id": ObjectId("646d4551405d0301468a5856"),
    "transaction_id": "e35894d2-5703-4e31-b951-3a3513a711b2",
    "account_number": "434134514",
    "customer_id": "72082738-7b99-4411-9d40-e2155129a406",
    "employee_id": "2c295092-8673-4d35-8841-928c52362e58",
    "amount": 382.29,
    "date": "2024-05-22T03:47:42.018Z",
    "type": "debit"
  },
  {
    "_id": ObjectId("646d4551405d0301468a5857"),
    "transaction_id": "521054e9-4430-4c35-a8f9-957254b9c463",
    "account_number": "884561899",
    "customer_id": "f6b30b92-5300-45d6-98a1-a64942d07871",
    "employee_id": "320877d3-5a44-414e-8479-3d21272e3413",
    "amount": 521.64,
    "date": "2024-05-23T15:19:44.552Z",
    "type": "transfer"
  },
  {
    "_id": ObjectId("646d4551405d0301468a5858"),
    "transaction_id": "0526e929-8760-4349-883b-c9238840427c",
    "account_number": "542665177",
    "customer_id": "85755781-1286-461c-b430-170f4d2783a7",
    "employee_id": "45a6c1c6-e443-42f7-96d8-079d5d6e9012",
    "amount": 101.38,
    "date": "2024-05-22T22:32:16.154Z",
    "type": "debit"
  },
  {
    "_id": ObjectId("646d4551405d0301468a5859"),
    "transaction_id": "d81a642a-717a-4b7d-886a-78216131d64c",
    "account_number": "288199871",
    "customer_id": "97475799-577d-423c-8039-b764521a0522",
    "employee_id": "75b9568e-9567-47e8-9796-878f60e72597",
    "amount": 268.85,
    "date": "2024-05-22T01:42:39.386Z",
    "type": "credit_payment"
  },
  {
    "_id": ObjectId("646d4551405d0301468a585a"),
    "transaction_id": "c986f91b-1422-4c24-b87f-241b26752369",
    "account_number": "999938966",
    "customer_id": "7d324866-6668-4923-a8e6-d1e369a32d96",
    "employee_id": "21877404-a222-4282-a632-4a0f8f4d7e63",
    "amount": 464.82,
    "date": "2024-05-22T13:17:33.482Z",
    "type": "credit_payment"
  },
  {
    "_id": ObjectId("646d4551405d0301468a585b"),
    "transaction_id": "09408240-f548-4654-92f8-14744740c702",
    "account_number": "245747244",
    "customer_id": "16034403-9b39-44f2-a39a-37335551109d",
    "employee_id": "8c7176d2-0335-4264-97c4-3054f7f85e65",
    "amount": 771.78,
    "date": "2024-05-23T14:23:28.387Z",
    "type": "debit"
  },
  {
    "_id": ObjectId("646d4551405d0301468a585c"),
    "transaction_id": "4844e72a-a903-425c-b783-5b8984a0704e",
    "account_number": "493524825",
    "customer_id": "85755781-1286-461c-b430-170f4d2783a7",
    "employee_id": "320877d3-5a44-414e-8479-3d21272e3413",
    "amount": 533.54,
    "date": "2024-05-23T13:44:24.154Z",
    "type": "transfer"
  },
  {
    "_id": ObjectId("646d4551405d0301468a585d"),
    "transaction_id": "b4667a10-0a3e-407a-b450-024a0f024179",
    "account_number": "454171526",
    "customer_id": "49654291-2737-4a5d-9516-4c2887922a11",
    "employee_id": "a853a240-f965-44a9-8c73-a58754136743",
    "amount": 950.86,
    "date": "2024-05-22T21:22:48.887Z",
    "type": "debit"
  },
  {
    "_id": ObjectId("646d4551405d0301468a585e"),
    "transaction_id": "a00554a4-3000-4278-a320-37654d004276",
    "account_number": "434983853",
    "customer_id": "e70c8323-4811-4384-a279-981249896413",
    "employee_id": "7874325b-4774-40c5-94d9-74100684a589",
    "amount": 800.22,
    "date": "2024-05-23T11:25:29.262Z",
    "type": "transfer"
  },
  {
    "_id": ObjectId("646d4551405d0301468a585f"),
    "transaction_id": "8d609210-1847-4539-962e-5206c70920e7",
    "account_number": "227788148",
    "customer_id": "72082738-7b99-4411-9d40-e2155129a406",
    "employee_id": "2c295092-8673-4d35-8841-928c52362e58",
    "amount": 221.16,
    "date": "2024-05-23T10:19:17.279Z",
    "type": "credit_payment"
  },
  {
    "_id": ObjectId("646d4551405d0301468a5860"),
    "transaction_id": "844712f2-f183-485d-9a3d-3b3259097141",
    "account_number": "248959412",
    "customer_id": "f6b30b92-5300-45d6-98a1-a64942d07871",
    "employee_id": "320877d3-5a44-414e-8479-3d21272e3413",
    "amount": 844.62,
    "date": "2024-05-22T23:06:57.533Z",
    "type": "debit"
  },
  {
    "_id": ObjectId("646d4551405d0301468a5861"),
    "transaction_id": "4315938a-a826-4e72-9794-0598a499b634",
    "account_number": "693448967",
    "customer_id": "85755781-1286-461c-b430-170f4d2783a7",
    "employee_id": "45a6c1c6-e443-42f7-96d8-079d5d6e9012",
    "amount": 333.88,
    "date": "2024-05-22T21:21:56.377Z",
    "type": "credit_payment"
  },
  {
    "_id": ObjectId("646d4551405d0301468a5862"),
    "transaction_id": "91010a09-0d2f-4520-922a-08179d509d38",
    "account_number": "812561844",
    "customer_id": "97475799-577d-423c-8039-b764521a0522",
    "employee_id": "75b9568e-9567-47e8-9796-878f60e72597",
    "amount": 680.87,
    "date": "2024-05-22T09:21:52.153Z",
    "type": "debit"
  },
  {
    "_id": ObjectId("646d4551405d0301468a5863"),
    "transaction_id": "09443163-3a09-46b6-94c1-311b874a8431",
    "account_number": "925578821",
    "customer_id": "7d324866-6668-4923-a8e6-d1e369a32d96",
    "employee_id": "21877404-a222-4282-a632-4a0f8f4d7e63",
    "amount": 931.31,
    "date": "2024-05-22T04:07:32.144Z",
    "type": "transfer"
  },
  {
    "_id": ObjectId("646d4551405d0301468a5864"),
    "transaction_id": "878909a4-351c-41a4-86d9-e3f9596b0086",
    "account_number": "761572968",
    "customer_id": "16034403-9b39-44f2-a39a-37335551109d",
    "employee_id": "8c7176d2-0335-4264-97c4-3054f7f85e65",
    "amount": 791.89,
    "date": "2024-05-22T14:01:23.327Z",
    "type": "credit_payment"
  },
  {
    "_id": ObjectId("646d4551405d0301468a5865"),
    "transaction_id": "f95d35d4-8a59-414a-b845-27f3d04e239d",
    "account_number": "913635599",
    "customer_id": "85755781-1286-461c-b430-170f4d2783a7",
    "employee_id": "320877d3-5a44-414e-8479-3d21272e3413",
    "amount": 177.68,
    "date": "2024-05-23T17:16:59.268Z",
    "type": "debit"
  },
  {
    "_id": ObjectId("646d4551405d0301468a5866"),
    "transaction_id": "45131406-2d3e-4064-a6d3-3c6d479d7986",
    "account_number": "957156362",
    "customer_id": "49654291-2737-4a5d-9516-4c2887922a11",
    "employee_id": "a853a240-f965-44a9-8c73-a58754136743",
    "amount": 478.93,
    "date": "2024-05-23T05:34:16.671Z",
    "type": "credit_payment"
  },
  {
    "_id": ObjectId("646d4551405d0301468a5867"),
    "transaction_id": "06c86c64-4752-44c4-9d52-8920b2458c6b",
    "account_number": "473318991",
    "customer_id": "e70c8323-4811-4384-a279-981249896413",
    "employee_id": "7874325b-4774-40c5-94d9-74100684a589",
    "amount": 204.77,
    "date": "2024-05-22T11:09:36.264Z",
    "type": "transfer"
  },
  {
    "_id": ObjectId("646d4551405d0301468a5868"),
    "transaction_id": "392c73d9-0516-4736-981a-03420563d77a",
    "account_number": "552132141",
    "customer_id": "72082738-7b99-4411-9d40-e2155129a406",
    "employee_id": "2c295092-8673-4d35-8841-928c52362e58",
    "amount": 945.84,
    "date": "2024-05-23T11:01:48.115Z",
    "type": "debit"
  },
  {
    "_id": ObjectId("646d4551405d0301468a5869"),
    "transaction_id": "b949d907-7b3c-4e6c-a20f-1526c9508558",
    "account_number": "524236385",
    "customer_id": "f6b30b92-5300-45d6-98a1-a64942d07871",
    "employee_id": "320877d3-5a44-414e-8479-3d21272e3413",
    "amount": 960.71,
    "date": "2024-05-22T04:01:59.858Z",
    "type": "transfer"
  },
  {
    "_id": ObjectId("646d4551405d0301468a586a"),
    "transaction_id": "e12b7705-4b76-49f4-a123-f7f10e33d407",
    "account_number": "582348822",
    "customer_id": "85755781-1286-461c-b430-170f4d2783a7",
    "employee_id": "45a6c1c6-e443-42f7-96d8-079d5d6e9012",
    "amount": 854.51,
    "date": "2024-05-23T10:28:55.658Z",
    "type": "credit_payment"
  },
  {
    "_id": ObjectId("646d4551405d0301468a586b"),
    "transaction_id": "59e21221-641c-423c-a787-364200700d0f",
    "account_number": "162151886",
    "customer_id": "97475799-577d-423c-8039-b764521a0522",
    "employee_id": "75b9568e-9567-47e8-9796-878f60e72597",
    "amount": 766.13,
    "date": "2024-05-22T12:12:19.847Z",
    "type": "debit"
  },
  {
    "_id": ObjectId("646d4551405d0301468a586c"),
    "transaction_id": "1c522078-4597-4466-91f5-433a0790f034",
    "account_number": "233972139",
    "customer_id": "7d324866-6668-4923-a8e6-d1e369a32d96",
    "employee_id": "21877404-a222-4282-a632-4a0f8f4d7e63",
    "amount": 141.15,
    "date": "2024-05-22T18:37:53.286Z",
    "type": "transfer"
  },
  {
    "_id": ObjectId("646d4551405d0301468a586d"),
    "transaction_id": "a5d6c5f4-26c9-44a3-87a2-b2a0b6427076",
    "account_number": "552181556",
    "customer_id": "16034403-9b39-44f2-a39a-37335551109d",
    "employee_id": "8c7176d2-0335-4264-97c4-3054f7f85e65",
    "amount": 500.34,
    "date": "2024-05-23T15:05:13.944Z",
    "type": "credit_payment"
  }
]);

EOF

echo "Load sample dataset for database '$DATABASE_NAME' successfully."
