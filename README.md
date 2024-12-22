# hr-system


--

```
hr-system/
.
├── LICENSE
├── README.md
├── api-gateway
│   ├── client
│   │   ├── auth_client.go
│   │   └── employee_client.go
│   ├── handler
│   │   ├── auth.go
│   │   ├── employee.go
│   │   ├── err.go
│   │   └── health.go
│   ├── main.go
│   ├── middlewares
│   │   └── jwt.go
│   ├── routes
│   │   └── routes.go
│   └── service
│       ├── auth_service.go
│       ├── base.go
│       └── employee_service.go
├── auth-service
│   ├── internal
│   │   └── handler.go
│   └── main.go
├── common
│   ├── config
│   │   └── config.go
│   ├── dao
│   │   ├── models
│   │   │   ├── departments.go
│   │   │   ├── employees.go
│   │   │   ├── positions.go
│   │   │   └── users.go
│   │   └── query
│   │       ├── departments.gen.go
│   │       ├── employees.gen.go
│   │       ├── gen.go
│   │       ├── positions.gen.go
│   │       └── users.gen.go
│   ├── database
│   │   ├── gorm-gen
│   │   │   └── gen.go
│   │   ├── mysql.go
│   │   └── plugin_redis.go
│   ├── jwt
│   │   └── jwt.go
│   ├── logger
│   │   └── logger.go
│   ├── makefile
│   └── proto
│       ├── auth.pb.go
│       ├── auth.proto
│       ├── auth_grpc.pb.go
│       ├── employee.pb.go
│       ├── employee.proto
│       └── employee_grpc.pb.go
├── docker
│   ├── Dockerfile.api-gateway
│   ├── Dockerfile.auth-service
│   ├── Dockerfile.employee-service
│   ├── Dockerfile.migrations
│   ├── docker-compose-db-only.yaml
│   └── docker-compose.yaml
├── docs
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── employee-service
│   ├── internal
│   │   ├── handler.go
│   │   └── handler_test.go
│   └── main.go
├── go.mod
├── go.sum
├── makefile
├── migrations
│   └── main.go
└── seeddata
    └── seed_data.go
｀``

# System Architecture Diagram (micro service)

                   +----------------------+
                   |      Client          |
                   |  (Web/Mobile App)    |
                   +----------+-----------+
                              |
                              v
                   +----------+-----------+
                   |      API Gateway      |
                   | (Routing & Auth Layer)|
                   +----------+-----------+
                              |
      +-----------------------+-----------------------+
      |                                               |
      v                                               v
+----------------+                          +------------------+
|  Auth Service  |                          | Employee Service |
|  (User Auth)   |                          | (Employee Mgmt)  |
+----------------+                          +------------------+
                              |
                              v
                   +----------+-----------+
                   |    Common Modules     |
                   | (Shared Logic, Tools) |
                   +----------+-----------+
                              |
                              v
         +-----------------------------------------+
         |     Database Layer (MySQL & Redis)      |
         |  (Structured Data & Caching Layer)      |
         +-----------------------------------------+
```


# System Description

## 1. Client (Frontend Application)
- Includes Web or Mobile Apps that interact with backend services via the API Gateway.

---

## 2. API Gateway
### Features
- **Routing Management**: Directs requests to appropriate services like `auth-service` and `employee-service`.
- **Business Logic Layer**: Handles interactions between multiple services.
- **Authentication**: Handles JWT authentication (via `middlewares/jwt.go`).

### Relevant Files
- `api-gateway/handler/*`: Handles specific routing logic.
- `api-gateway/middlewares/jwt.go`: JWT authentication middleware.
- `api-gateway/routes/routes.go`: Routing configuration.

---

## 3. Auth Service
### Features
- Manages user login, registration, and other related operations.
- Provides APIs related to user management.

### Relevant Files
- `auth-service/internal/handler.go`: Handles specific business logic.
- `auth-service/main.go`: Service entry point.

---

## 4. Employee Service
### Features
- Manages operations related to employee management.
- Interacts with the database to manage employee data.

### Relevant Files
- `employee-service/internal/handler.go`: Handles specific business logic.
- `employee-service/main.go`: Service entry point.

---

## 5. Common Modules
### Features
- Encapsulates shared logic across services:
  - **Configuration Management**: `common/config/`
  - **Data Access Objects (DAO)**: `common/dao/models` and `common/dao/query`
  - **Database Initialization and Management**: `common/database/mysql.go` and `gorm-gen`
  - **JWT Utilities**: `common/jwt/jwt.go`
  - **Logging Utilities**: `common/logger/logger.go`
  - **gRPC and Proto Definitions**: `common/proto/`

---

## 6. Data Storage
- **MySQL**: Stores structured data like employee information, user accounts, departments, and positions.
- **Redis**: Used for caching, potentially speeding up JWT validation or other queries.

---

## 7. Containerized Deployment (Docker)
- Defines Docker images for different services via multiple `Dockerfile` configurations.
- Uses `docker-compose.yaml` for service startup and network configuration.

---

## 8. Documentation and Other Modules
- **Documentation**: `docs` contains Swagger documentation for API reference.
- **Database Migrations**: `migrations` manages database schema migrations.
- **Seed Data**: `seeddata` is used to generate seed data for development or testing.

---



---
# Makefile Description
- Run unit tests (ensure database is up before running):
  ```
  make test
  ```
- Manage dependencies:
  ```
  make mod
  ```
- Start the entire service stack:
  ```
  make docker-compose-up
  ```
- Start only the database service:
  ```
  make docker-compose-up-db-only
  ```
- Generate Swagger documentation:
  ```
  make swag
  ```
- Generate GORM model code:
  ```
  make gorm-gen
  ```

---
# Quick Start

```
Start all services

> make docker-compose-up
```

```
Login And Get Token 

> curl -X 'POST' \
  'http://localhost:8081/api/login' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{}'

{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzQ5MjI5OTEsInVzZXJfaWQiOjd9.8ny3-GIKAjfq9mhHPSmn4sO1L928QPBsGut6NFvwlU0"
}


Get Employee #1 Data

> curl -X 'GET' \
  'http://localhost:8081/api/employees/1' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzQ5MjI5OTEsInVzZXJfaWQiOjd9.8ny3-GIKAjfq9mhHPSmn4sO1L928QPBsGut6NFvwlU0'


{"id":1,"name":"Alice","email":"alice@example.com","phone":"111222333","department_id":1,"position_id":1,"hire_date":{},"salary":50000,"department":{"id":1,"name":"Engineering"},"position":{"id":1,"title":"Junior Developer","level":"Junior"}}%   %                                                                         

```
