
# Player Service

Use command: `git clone https://github.com/tarkovskynik/player-service.git`

## Dependencies:
Docker

Docker image PostgreSQL

golang-migrate (brew install golang-migrate)

go

## Build & Run:
Use command for create Data Base: `docker run --name=players-db -e POSTGRES_PASSWORD='qwerty' -p 5432:5432 -d postgres`

Use: `make migrate` for migrates

Build the app: `go build main.go` in terminal or in IDE 


## Operations:

POST - `"/user/create"` - create user

GET - `"/user/get"` - get user

POST `"/deposit"` - accrues money to the user for a completed deposit

PUT `"/transaction"` - making a game transaction

Example of creating a user:
In Postman you choose "Body" menu, POST `localhost:8080/user/create` and type for example:

{
"id": 1,
"balance": 0,
"token": "testtask"
}

#### Token - with which token the user was created with such a token it will be possible to perform operations on this ID

Example of get user:
In Postman you choose "Body" menu, POST `localhost:8080/user/get` and type for example:

{"id": 1,
"token": "testtask"
}

Example of deposit:
In Postman you choose "Body" menu, POST `localhost:8080/deposit` and type for example:

{"user_id": 1,
"deposit_id": 1,
"amount": 500,
"token": "testtask"
}

Example of transaction:
In Postman you choose "Body" menu, POST `localhost:8080/transaction` and type for example:

{"user_id": 1,
"transaction_id": 1,
"type": "Win",
"amount": 1000,
"token": "testtask"
}
