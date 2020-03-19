# Auth-Service
Authentication service written in Golang with a MongoDB database

## Endpoints
* /SignUp
* /SignIn
* /Validate
  
#### SignUp
Request fields: Username, Password  
Writes username and password (with hash and salt) in the database

#### SignIn
Request fields: Username, Password  
Returns a json with a JWT token and an expiration time

#### Validate
Request fields: JWT Token  
Validates the token in the request

## Usage

#### First
Get dependencies  
``` go get go.mongodb.org/mongo-driver```

#### Build

```make build```

Generates a executable

#### Run

```make run```

Runs the server

