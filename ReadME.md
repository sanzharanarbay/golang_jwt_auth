# Golang Restful API using GORM ORM (PostgreSQL), Gorilla Mux, JWT

## Getting Started

### Folder Structure
This is my folder structure under my `$GOPATH` or `$HOME/your_username/go`.
```
.
|-- bin
+-- src
|   +-- github.com
|   |   +-- sanzharanarbay
|   |   |   +-- jwt_auth_golang
|   |   |   |   |-- .env
|   |   |   |   |-- main.go
|   |   |   |   |-- .gitignore.go
|   |   |   |   |-- README.md
|   |   |   |   |   +-- controllers
|   |   |   |   |   |   | users
|   |   |   |   |   |   | +-- UsersController.go
|   |   |   |   |   +-- auth
|   |   |   |   |   |   |-- AuthController.go
|   |   |   |   +-- models
|   |   |   |   |   |-- users.go
|   |   |   |   +-- middleware
|   |   |   |   |   |-- middleware.go
|   |   |   |   +-- handlers
|   |   |   |   |   |-- authHandler.go
|   |   |   |   +-- modules
|   |   |   |   |   |   +-- auth
|   |   |   |   |   |   |   |-- auth.go
|   |   |   |   |   |   |   |-- token.go
|   |   |   |   +-- routes
|   |   |   |   |   |-- api.go
|   |   |   |   +-- utils
|   |   |   |   |   |-- utils.go
```

## Download the packages used to create this rest API
Run the following Golang commands to install all the necessary packages. These packages will help you set up a web server, ORM for interacting with your db, mysql driver for db connection, load your environment variables from the .env file and generate JWT tokens.

```
go get -u github.com/gorilla/mux
go get -u github.com/jinzhu/gorm
go get -u github.com/joho/godotenv
go get -u github.com/dgrijalva/jwt-go
go get -u github.com/gin-gonic/gin
go get -u github.com/go-redis/redis/v7
```

### Running documentation locally (Only documentation of packages your have installed)
set port APP_PORT on .env file , then run and Visit `http://localhost:{APP_PORT}`. Note that you can change the port to your preferred port number.

## Setting configuration file
Create a .env file in the root of the project and set the following parameters

```
REDIS_HOST=127.0.0.1                    # Redis host
REDIS_PORT=6379                         # Redis post
REDIS_PASSWORD=                         # Redis password (default null)
ACCESS_SECRET=nbqzUMf90s56n7Lry5KeXJWMXrIsn4Mv      # Access_token secret key
REFRESH_SECRET=RQPfAOgu63vGVtLxDcdNppSQiI2RbUvM      # refresh_token secret key
APP_PORT=9001                                   # Application port
APP_HOST=localhost                              # Application host
GIN_MODE=debug                                  # Application MODE (release, debug, test)
ROUTE_PREFIX=/api/v1                            # API prefix
DB_HOST=127.0.0.1                               # database host
DB_USER=postgres                                # database user
DB_PASSWORD=postgres                            # database password
DB_NAME=golang_jwt_auth                         # database name
DB_PORT=5432                                    # database port
DB_TYPE=postgres                                # database type
ACCESS_TOKEN_EXPIRE=30                          #  in minutes
REFRESH_TOKEN_EXPIRE=7                          # in hours

```

## Running the project

`go run main.go`

## Database Table Creation Statement
Use the following DDL (Data Definition Language) to create the users table.

``` SQL
CREATE TABLE `users` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `username` varchar(255) DEFAULT NULL,
  `password` varchar(255) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=19 DEFAULT CHARSET=utf8;
```

## API Endpoints & Usage

To be able to login, you need to use the create new user endpoint to set up a user account.

* POST    `/login` login with username and password in your db to get token back
* POST    `/register` create a new account with a username, password and name
* POST    `/refresh` refresh tokens
* POST    `/logout` logout from current account
* GET     `api/v1/dashboard/users` retrieve all users
* GET     `api/v1/dashboard/users/1` retrieve user with id = 1
* POST    `api/v1/dashboard/users` create a new user
* PUT     `api/v1/dashboard/users/1` update the record with id = 1
* DELETE  `api/v1/dashboard/users/1` delete the user with id = 1

### To create a new user

1. POST `api/v1/register`

```
{
	"Name": "Joe Bloke",
	"Username": "joe.bloke@fake-domain.com",
    "Password": "secret"
}
```

*** Output ***

Note that the current implementation still returns the encrypted password, this needs to be removed from the response.

```
{
    "message": "success",
    "status": true,
    "user": {
        "ID": 1,
        "CreatedAt": "2019-05-06T00:54:22.09382+01:00",
        "UpdatedAt": "2019-05-06T00:54:22.09382+01:00",
        "DeletedAt": null,
        "Name": "Joe Bloke",
        "Username": "joe.bloke@fake-domain.com"
    }
}
```

2. POST `api/v1/login`



*** Input ***
```
{
  "Username": "joe.bloke@fake-domain.com"
  "Password": "secret"
}
```

*** Output ***

```
{
    "data": {
        "access-token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjkwMDAwMDAwMDAwMCwicGFzc3dvcmQiOiJzZWNyZXQiLCJ1c2VybmFtZSI6ImFtYXZpQHh5ei5jb20ifQ.WJ5VMnH5ijHQOZhUlrrnrh7NCYfFpww3jBz26EkRsHQ"
    },
    "message": "success",
    "status": true
}
```
