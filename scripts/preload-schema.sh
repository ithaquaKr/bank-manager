#!/bin/bash

# MongoDB connection details
#MONGO_USER="ithadevnguyen"
MONGO_USER="admin"
#MONGO_PASSWORD=""
MONGO_PASSWORD="secretpassword"
#MONGO_URI="mongodb+srv://bank-manager.z3mb2.mongodb.net/"
MONGO_URI="mongodb://127.0.0.1:27017/"

# Database and collection names
DATABASE_NAME="bank-manager"

# Connect to MongoDB
mongosh $MONGO_URI --apiVersion 1 --username $MONGO_USER --password $MONGO_PASSWORD <<EOF
use $DATABASE_NAME

db.createCollection("employee", {
    "capped": false,
    "validator": {
        "$jsonSchema": {
            "bsonType": "object",
            "title": "employee",
            "properties": {
                "_id": {
                    "bsonType": "objectId"
                },
                "employee_id": {
                    "bsonType": "string"
                },
                "identity_card": {
                    "bsonType": "string"
                },
                "fullname": {
                    "bsonType": "string"
                },
                "date_of_birth": {
                    "bsonType": "date"
                },
                "address": {
                    "bsonType": "string"
                },
                "job_rank": {
                    "bsonType": "string"
                },
                "position": {
                    "bsonType": "string"
                },
                "seniority": {
                    "bsonType": "number",
                    "minimum": 0
                },
                "accounts_created": {
                    "bsonType": "array",
                    "additionalItems": true,
                    "items": {
                        "bsonType": "object",
                        "properties": {
                            "account_number": {
                                "bsonType": "string"
                            },
                            "type": {
                                "bsonType": "string"
                            },
                            "initial_deposit": {
                                "bsonType": "number",
                                "minimum": 0
                            },
                            "created_at": {
                                "bsonType": "date"
                            }
                        },
                        "additionalProperties": false
                    }
                }
            },
            "additionalProperties": false
        }
    },
    "validationLevel": "off",
    "validationAction": "warn"
});

db.createCollection("customer", {
    "capped": false,
    "validator": {
        "$jsonSchema": {
            "bsonType": "object",
            "title": "customer",
            "properties": {
                "_id": {
                    "bsonType": "objectId"
                },
                "customer_id": {
                    "bsonType": "string"
                },
                "identity_card": {
                    "bsonType": "string"
                },
                "fullname": {
                    "bsonType": "string"
                },
                "date_of_birth": {
                    "bsonType": "string"
                },
                "address": {
                    "bsonType": "string"
                },
                "accounts": {
                    "bsonType": "array",
                    "additionalItems": true,
                    "items": {
                        "bsonType": "object",
                        "properties": {
                            "account_number": {
                                "bsonType": "string"
                            },
                            "account_name": {
                                "bsonType": "string"
                            },
                            "account_type": {
                                "bsonType": "string"
                            },
                            "balance": {
                                "bsonType": "number",
                                "minimum": 0
                            },
                            "credit_limit": {
                                "bsonType": "number",
                                "minimum": 0
                            },
                            "outstanding_balance": {
                                "bsonType": "number",
                                "minimum": 0
                            },
                            "min_balance": {
                                "bsonType": "number",
                                "minimum": 0
                            },
                            "monthly_interest_rate": {
                                "bsonType": "number",
                                "minimum": 0
                            },
                            "transactions": {
                                "bsonType": "array",
                                "additionalItems": true,
                                "items": {
                                    "bsonType": "object",
                                    "properties": {
                                        "transaction_id": {
                                            "bsonType": "string"
                                        },
                                        "amount": {
                                            "bsonType": "number"
                                        },
                                        "created_at": {
                                            "bsonType": "date"
                                        }
                                    },
                                    "additionalProperties": false
                                }
                            }
                        },
                        "additionalProperties": false
                    }
                }
            },
            "additionalProperties": false
        }
    },
    "validationLevel": "off",
    "validationAction": "warn"
});

db.createCollection("transaction", {
    "capped": false,
    "validator": {
        "$jsonSchema": {
            "bsonType": "object",
            "title": "transaction",
            "properties": {
                "_id": {
                    "bsonType": "objectId"
                },
                "transaction_id": {
                    "bsonType": "string"
                },
                "account_number": {
                    "bsonType": "string"
                },
                "customer_id": {
                    "bsonType": "string"
                },
                "employee_id": {
                    "bsonType": "string"
                },
                "amount": {
                    "bsonType": "number"
                },
                "created_at": {
                    "bsonType": "date"
                },
                "type": {
                    "bsonType": "string"
                }
            },
            "additionalProperties": false
        }
    },
    "validationLevel": "off",
    "validationAction": "warn"
});

db.employee.createIndex({ "employee_id": 1 }, { unique: true });
db.employee.createIndex({ "identity_card": 1 }, { unique: true });
db.customer.createIndex({ "customer_id": 1 }, { unique: true });
db.customer.createIndex({ "identity_card":1 }, { unique: true });
db.customer.createIndex({ "accounts.account_number": 1 }, { unique: true });
db.customer.createIndex({ "accounts.transactions.transaction_id": 1 }, { unique: true });
db.transaction.createIndex({ "transaction_id": 1 }, { unique: true });
EOF

echo "Database '$DATABASE_NAME' and collection created successfully."
