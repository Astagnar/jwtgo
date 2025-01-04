# jwtgo

A Go (Golang) Backend Clean Architecture project with Gin, MongoDB and JWT Authentication Middleware.

The project was created for educational purposes and can be used in your projects as needed.

## Project Architecture
The architecture of a web application consists of 3 main layers:
- Controller
- Service
- Repository

![Image](https://raw.githubusercontent.com/Astagnar/jwtgo/refs/heads/main/assets/architecture.png)

## Major Packages used in project
- **[Gin](https://pkg.go.dev/github.com/gin-gonic/gin)**: Gin is a HTTP web framework written in Go (Golang). It features a Martini-like API with much better performance -- up to 40 times faster. If you need smashing performance, get yourself some Gin. 
- **[MongoDB](https://pkg.go.dev/go.mongodb.org/mongo-driver)**: The Official Golang driver for MongoDB.
- **[JWT](https://pkg.go.dev/github.com/golang-jwt/jwt/v5)**: Go implementation of JSON Web Tokens (JWT).
- **[CleanENV](https://pkg.go.dev/github.com/ilyakaznacheev/cleanenv)**: Clean and minimalistic environment configuration reader for Golang.
- **[Bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt)**: Package bcrypt implements Provos and Mazières's bcrypt adaptive hashing algorithm.
- **[Logrus](https://pkg.go.dev/github.com/sirupsen/logrus)**: Structured, pluggable logging for Go.

## Request Flow without JWT Authentication Middleware
![Image](https://raw.githubusercontent.com/Astagnar/jwtgo/refs/heads/main/assets/without-jwt.png)

## Request Flow with JWT Authentication Middleware
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
- Navigate to the `configs` folder and create a `config.yaml` file, similar to `config.yaml.sample`, in this directory.
- Install the `go` if it is not installed on your computer.
- Install the `MongoDB` if it is not installed on your computer.
- Fill in the `config.yaml` file with your data.
- Run `go run cmd/app/main.go`.
- Access API using http://127.0.0.1:8000.

### Run with Docker
- Coming soon.