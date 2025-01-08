# jwtgo

A Go (Golang) Backend Clean Architecture project with Gin, MongoDB and JWT Authentication Middleware.

The project was created for educational purposes and can be used in your projects as needed.

## Project architecture
The architecture of a web application consists of 3 main layers:
- Controller
- Service
- Repository

![Image](https://raw.githubusercontent.com/Astagnar/jwtgo/refs/heads/main/assets/architecture.png)

## Major packages used in project
- **[Gin](https://pkg.go.dev/github.com/gin-gonic/gin)**: Gin is a HTTP web framework written in Go (Golang). It features a Martini-like API with much better performance -- up to 40 times faster. If you need smashing performance, get yourself some Gin. 
- **[MongoDB](https://pkg.go.dev/go.mongodb.org/mongo-driver)**: The Official Golang driver for MongoDB.
- **[JWT](https://pkg.go.dev/github.com/golang-jwt/jwt/v5)**: Go implementation of JSON Web Tokens (JWT).
- **[Cleanenv](https://pkg.go.dev/github.com/ilyakaznacheev/cleanenv)**: Clean and minimalistic environment configuration reader for Golang.
- **[Bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt)**: Package bcrypt implements Provos and Mazières's bcrypt adaptive hashing algorithm.
- **[Logrus](https://pkg.go.dev/github.com/sirupsen/logrus)**: Structured, pluggable logging for Go.
- **[Validator](https://pkg.go.dev/github.com/go-playground/validator/v10)**: Go Struct and Field validation, including Cross Field, Cross Struct, Map, Slice and Array diving.

## Request flow without JWT authentication middleware
![Image](https://raw.githubusercontent.com/Astagnar/jwtgo/refs/heads/main/assets/without-jwt.png)

## Request flow with JWT authentication middleware
![Image](https://raw.githubusercontent.com/Astagnar/jwtgo/refs/heads/main/assets/with-jwt.png)

## How to run the project?
First, download it and navigate to the root directory:
```bash
# Move to your workspace
cd your-workspace

# Clone the project into your workspace
git clone https://github.com/Astagnar/jwtgo.git

# Move to the project root directory
cd jwtgo
```

### Run without Docker
- Create a `.env` file, similar to `.env.sample`.
- Install the `go` if it is not installed on your computer.
- Install the `MongoDB` if it is not installed on your computer.
- Fill in the `.env` file with your data.
- Download the dependencies with the following command:
  ```bash
  go mod download
  ```
- Run the application with the following command:
  ```bash
  go run cmd/app/main.go
  ```
- Access API using http://127.0.0.1:8001.

### Run with Docker
- Create a `.env` file, similar to `.env.sample`.
- Install the `Docker` if it is not installed on your computer.
- Fill in the `.env` file with your data.
- Run the application build with the following command:
  ```bash
  docker-compose -f deployments/docker-compose.yaml --env-file=".env" up -d
  ```
- Access API using http://127.0.0.1:8001.

## Examples of API requests and responses
### SignUp endpoint
- Request:
  ```
  curl --location 'http://localhost:8001/auth/signup' \
  --header 'Content-Type: application/json' \
  --data-raw '{
    "email": "test@gmail.com",
    "password": "securepassword"
  }'
  ```
  
- Response:
  ```
  HTTP/1.1 200 OK
  Content-Type: application/json
  ```
  ```json
  {
    "message": "User successfully registered"
  }
  ```

### SignIn endpoint
- Request:
  ```
  curl --location 'http://localhost:8001/auth/signin' \
  --header 'Content-Type: application/json' \
  --data-raw '{
    "email": "test@gmail.com",
    "password": "securepassword"
  }'
  ```

- Response:
  ```
  HTTP/1.1 200 OK
  Content-Type: application/json
  Set-Cookie: access_token=access_token; Path=/; HttpOnly; SameSite=Strict
  Set-Cookie: refresh_token=refresh_token; Path=/; HttpOnly; SameSite=Strict
  ```
  ```json
  {
    "message": "User successfully logged in"
  }
  ```

### Refresh endpoint
- Request:
  ```
  curl --location 'http://localhost:8001/auth/refresh' \
  --header 'Content-Type: application/json' \
  -b 'access_token=access_token; refresh_token=refresh_token'
  ```

- Response:
  ```
  HTTP/1.1 200 OK
  Content-Type: application/json
  Set-Cookie: access_token=access_token; Path=/; HttpOnly; SameSite=Strict
  Set-Cookie: refresh_token=refresh_token; Path=/; HttpOnly; SameSite=Strict
  ```
  ```json
  {
    "message": "Tokens successfully updated"
  }
  ```

## Complete project folder structure
```
├─── .env
├─── build
│   └─── package
│       └─── Dockerfile
├─── cmd
│   └─── app
│       └─── main.go
├─── deployments
│   └─── docker-compose.yaml
├─── internal
│   ├─── app
│   │   ├─── main.go
│   │   ├─── adapter
│   │   │   └─── mongodb
│   │   │       ├─── entity
│   │   │       │   └─── user.go
│   │   │       ├─── mapper
│   │   │       │   └─── user.go
│   │   │       └─── repository
│   │   │           └─── user.go
│   │   ├─── config
│   │   │   └─── config.go
│   │   ├─── controller
│   │   │   └─── http
│   │   │       ├─── dto
│   │   │       │   └─── user.go
│   │   │       ├─── mapper
│   │   │       │   └─── user.go
│   │   │       ├─── middleware
│   │   │       │   ├─── security.go
│   │   │       │   └─── validation.go
│   │   │       └─── v1
│   │   │           └─── auth.go
│   │   ├─── entity
│   │   │   └─── user.go
│   │   ├─── error
│   │   │   ├─── auth.go
│   │   │   ├─── jwt.go
│   │   │   └─── server.go
│   │   ├─── interface
│   │   │   ├─── repository
│   │   │   │   └─── user.go
│   │   │   └─── service
│   │   │       ├─── auth.go
│   │   │       ├─── jwt.go
│   │   │       └─── password.go
│   │   ├─── schema
│   │   │   └─── jwt.go
│   │   └─── service
│   │       ├─── auth.go
│   │       ├─── jwt.go
│   │       └─── password.go
│   └─── pkg
│       └─── request
│           ├─── response.go
│           └─── schema
│               └─── response.go
└─── pkg
    ├─── client
    │   └─── mongodb.go
    └─── logging
        └─── logger.go
```