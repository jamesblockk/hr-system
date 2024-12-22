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

# 系統架構圖

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


# 系統描述

## 1. Client (前端應用)
- 包括 Web 或 Mobile App，透過 API Gateway 與後端服務交互。

---

## 2. API Gateway
### 功能
- **路由管理**：分發請求到對應服務，例如 `auth-service` 和 `employee-service`。
- **中介層業務邏輯**：處理多服務交互。
- **驗證**：如處理 JWT 驗證（`middlewares/jwt.go`）。

### 相關檔案
- `api-gateway/handler/*`：處理具體路由邏輯。
- `api-gateway/middlewares/jwt.go`：JWT 驗證中間件。
- `api-gateway/routes/routes.go`：路由設定。

---

## 3. Auth Service (認證服務)
### 功能
- 管理使用者登入、註冊等操作。
- 提供與使用者相關的 API。

### 相關檔案
- `auth-service/internal/handler.go`：處理具體業務邏輯。
- `auth-service/main.go`：啟動服務入口。

---

## 4. Employee Service (員工管理服務)
### 功能
- 管理員工相關的操作。
- 與數據庫交互管理員工的資料。

### 相關檔案
- `employee-service/internal/handler.go`：處理具體業務邏輯。
- `employee-service/main.go`：啟動服務入口。

---

## 5. Common Modules (共享模組)
### 功能
- 封裝跨服務的公共邏輯：
  - **配置管理**：`common/config/`
  - **資料訪問物件（DAO）**：`common/dao/models` 和 `common/dao/query`
  - **資料庫初始化與管理**：`common/database/mysql.go` 和 `gorm-gen`
  - **JWT 工具**：`common/jwt/jwt.go`
  - **日誌工具**：`common/logger/logger.go`
  - **gRPC 與 Proto 定義**：`common/proto/`

---

## 6. 數據存儲
- **MySQL**：存儲員工資料、使用者帳號、部門與職位等結構化數據。
- **Redis**：用於緩存，可能加速 JWT 驗證或其他查詢。

---

## 7. 容器化部署 (Docker)
- 透過多個 `Dockerfile` 為不同服務提供鏡像定義。
- 使用 `docker-compose.yaml` 管理服務間的啟動與網路配置。

---

## 8. 文件與其他模組
- **文件**：`docs` 包含 Swagger 文件，用於 API 說明。
- **數據遷移**：`migrations` 用於管理數據庫遷移。
- **種子數據**：`seeddata` 用於開發或測試時生成種子數據。



---
# Makefile 說明
- 執行單元測試(執行前資料庫須先開啟）：
  ```
  make test
  ```
- 整理依賴模組：
  ```
  make mod
  ```
- 啟動完整服務：
  ```
  make docker-compose-up
  ```
- 僅啟動資料庫：
  ```
  make docker-compose-up-db-only
  ```
- 生成 Swagger 文檔：
  ```
  make swag
  ```
- 生成 GORM 模型代碼：
  ```
  make gorm-gen
  ```

---
# 快速啟動

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
