#!/bin/bash

# MongoDB connection details
MONGO_USER="ithadevnguyen"
MONGO_PASSWORD=""

# Database and collection names
DATABASE_NAME="bank-manager"

# Connect to MongoDB
mongosh "mongodb+srv://bank-manager.z3mb2.mongodb.net/" --apiVersion 1 --username $MONGO_USER --password $MONGO_PASSWORD <<EOF
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
                    "bsonType": "string"
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
                                "bsonType": "string"
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
                            "debit_limit": {
                                "bsonType": "number",
                                "minimum": 0
                            },
                            "outstanding_debt": {
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
                                            "bsonType": "number",
                                            "minimum": 0
                                        },
                                        "date": {
                                            "bsonType": "string"
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
                    "bsonType": "number",
                    "minimum": 0
                },
                "date": {
                    "bsonType": "string"
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
EOF

echo "Database '$DATABASE_NAME' and collection created successfully."
