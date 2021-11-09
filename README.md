# Handling Transaction
This project help to produce a handling transaction concept in multiple microservices. There are 2 services: employee and assigment. Each service will communicate via Restful API, then each service will make a gRPC request to Orchestrator to apply a change to database.
Employee service will help to add new employee, then make a request to Assignment service to add a new assigment. If there is any error, system will reverse the changed in database.

For example: When adding a new employee, but a assignment is duplicated. So we can not add new assignment, system will noti to reverse the employee which already added.

## Prerequisites

Make sure you have the development environment matches with these notes below so we can mitigate any problems of version mismatch.

- Backend:
  - Go SDK: 1.16.
    Make sure to set `$GOROOT` and `$GOPATH` correctly.
    You can check those environment variable by typing: `go env`.
  - Postgre: 14.0.1
  - Redis 

- Commons:
  - Install [git](https://git-scm.com/) for manage source code.
  - IDE of your choice, recommended `Goland` or `VS Code`.

## Development

#### 1. Clone code to local

```shell
$ go get -u -v github.com/paphuc/handling-transactions
or
$ cd $GOPATH/src
$ git clone https://github.com/paphuc/handling-transactions.git
```
After this step, source code must be available at `$GOPATH/src/github.com/paphuc/handling-transactions`.

#### 2. Start development environment manually

- Start Postgres service at localhost:5432. The easiest way is to run the Docker as below:

  ```shell
  $ docker run --name postgres-database -e POSTGRES_PASSWORD=123456 -d postgres

  ```

- Start Redis service at port 6379. We can start with Docker as below:

  ```shell
  $ docker run --name redis-instance -p 6379:6379 -d redis

  ```

- Start backend API services:

  ```shell
  $ go run main.go
  # Backend service will start on port : 8081 (Employee service), 8082 (Assignment service), 9080 (Orchestrator)
  ```
