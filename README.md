# Golang Microservices Boilerplate - Clean Architecture

[![issues](https://img.shields.io/github/issues/phoenixTW/microservices-go)](https://github.com/phoenixTW/microservices-go/tree/master/.github/ISSUE_TEMPLATE)
[![license](https://img.shields.io/github/license/phoenixTW/microservices-go)](https://github.com/phoenixTW/microservices-go/tree/master/LICENSE)
[![CodeFactor](https://www.codefactor.io/repository/github/phoenixtw/microservices-go/badge)](https://www.codefactor.io/repository/github/phoenixtw/microservices-go)
[![Maintainability](https://api.codeclimate.com/v1/badges/9a323c190f2758377150/maintainability)](https://codeclimate.com/github/phoenixTW/microservices-go/maintainability)

Example structure to start a microservices project with golang. Using a MySQL databaseSQL. Using a Hexagonal
Architecture tha is a Clean Architecture.

## Manual Installation

If you would still prefer to do the installation manually, follow these steps:

Clone the repo:

```bash
git clone https://github.com/phoenixTW/microservices-go
```

If you need, configure the environment variables in file config.json, if you use docker-compose leave the variables set
in the file config.json.example

```bash 
cp config.json.example config.json
```

**TL;DR command list**

    git clone https://github.com/phoenixTW/microservices-go
    cd microservices-go
    cp config.json.example config.json
    docker-compose up  --build  -d

## Table of Contents

- [Features](#features)
- [Commands](#commands)
- [Environment Variables](#environment-variables)
- [Project Structure](#project-structure)
- [API Documentation](#api-documentation)
- [Error Handling](#error-handling)
- [Validation](#validation)
- [Linting](#linting)

## Features

- **Golang v1.21**: Stable version of go
- **Framework**: A stable version of [gin-go](https://github.com/gin-gonic/gin)
- **Token Security**: with [JWT](https://jwt.io)
- **SQL databaseSQL**: [MariaDB](https://mariadb.org/) using internal sql package of
  go [sql](https://golang.org/pkg/databaseSQL/sql/)
- **Testing**: unit and integration tests using package of go [testing](https://golang.org/pkg/testing/)
- **API documentation**: with [swaggo](https://github.com/swaggo/swag) @latest version that is a go implementation
  of [swagger](https://swagger.io/)
- **Dependency management**: with [go modules](https://golang.org/ref/mod)
- **Environment variables**: using [viper](https://github.com/spf13/viper)
- **Docker support**
- **Code quality**: with [CodeFactor](https://www.codefactor.io/) and [Codacy](https://www.codacy.com/)
- **Linting**: with [golangci-lint](https://golangci-lint.run/usage/install/) an implementation of a Golang linter

## Commands

### Build and run image of docker

```bash
docker-compose up  --build  -d
```

### Swagger Implementation

```bash
swag init -g infrastructure/rest/routes/routes.go
```

To visualize the swagger documentation on local use

http://localhost:8080/v1/swagger/index.html

### Unit test command

```bash
# run recursive test
go test  ./test/unit/...
# clean go test results in cache
go clean -testcache
```

### Lint inspection of go

```bash
golangci-lint run ./...
```



