# hr-system

```
hr-system/
├── api-gateway/
│   ├── main.go              # Gin RESTful API Gateway
│   ├── routes/              # API 路由定義
│   │   └── routes.go
│   └── middlewares/         # 通用中間件 (如 JWT 驗證)
│       └── jwt_middleware.go
├── auth-service/
│   ├── main.go              # 驗證服務入口
│   ├── proto/               # gRPC 定義檔案
│   │   ├── auth.proto
│   │   └── auth.pb.go
│   ├── internal/            # 核心業務邏輯
│   │   ├── handler.go
│   │   └── service.go
│   ├── migrations/          # GORM 資料庫遷移
│   │   └── migration.go
│   ├── seeds/               # GORM 資料庫初始化
│       └── seed.go
├── user-service/
│   ├── main.go
│   ├── proto/
│   ├── internal/
│   ├── migrations/
│   ├── seeds/
│   └── models/
├── common/                  # 共用模組 (工具、配置等)
│   ├── config/              # 配置管理
│   │   └── config.go
│   ├── logger/              # 日誌集成 (Graylog)
│   │   └── graylog.go
│   └── utils/               # 工具函數 (如加密、驗證)
├── docker/                  # Docker Compose 與 Dockerfile
│   ├── docker-compose.yml
│   ├── Dockerfile.api-gateway
│   ├── Dockerfile.auth-service
│   └── ...
├── makefile                 # Build 與 Deploy 指令
└── README.md
｀``




```
curl -X 'GET' \
  'http://localhost:8081/api/employees/1' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzQ1OTkzMDQsInVzZXJfaWQiOjd9.dYhpJ8WmoM2BDrdNza7KYJwiCU_UJJg-V1wgDfGQ9Qg'

```