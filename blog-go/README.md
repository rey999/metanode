# blog-go 项目文档

> 一个基于 Gin 框架和 GORM 的博客系统后端项目，支持用户登录、JWT 身份验证和 SQLite 数据库持久化。

## 📦 项目概述

本项目是一个轻量级博客系统后端服务，使用 Go 语言开发，采用 Gin 框架处理 HTTP 请求，GORM 作为 ORM 工具，并通过 JWT 实现身份验证机制。适用于学习 Go Web 开发、RESTful API 构建及微服务架构实践。

### 主要功能
- 用户注册与登录
- JWT Token 生成与解析（HS256 算法）
- 使用 SQLite 进行数据持久化
- MVC 架构设计

## 🧰 技术栈

| 技术 | 版本 | 说明 |
|------|------|------|
| Go | 1.24.4 | 项目主语言 |
| Gin | v1.10.1 | Web 框架 |
| GORM | v1.30.0 | ORM 数据库操作 |
| SQLite | - | 默认数据库驱动 |
| JWT | github.com/golang-jwt/jwt/v5 | 替代已废弃的 jwt-go |
| JSON 序列化 | github.com/goccy/go-json | 提升 JSON 性能 |
| 表单验证 | github.com/go-playground/validator/v10 | 验证输入合法性 |

## 📁 目录结构

```bash
.
├── common        # 公共工具类（如数据库连接）
├── config        # 配置管理
├── controller    # 控制器层，处理 HTTP 请求
├── entity        # 数据模型定义
├── middleware    # 中间件（如 JWT 验证）
├── route         # 路由定义
├── server        # 启动逻辑
├── service       # 业务逻辑层
├── utils         # 工具函数（如 Token 解析）
├── config.json   # 配置文件
└── main.go       # 应用入口
```

## 🔐 安全规范

- 所有 Token 使用 HS256 箾名算法生成。
- 密钥长度推荐为 256 bits，避免硬编码，建议通过环境变量配置。
- 不在 Token payload 中存储敏感信息。

## 🚀 安装与启动

### 安装依赖

```bash
go mod tidy
```

### 启动服务

```bash
go run main.go
```

### 构建可执行文件

```bash
go build -o blog-go main.go
```

## 🧪 示例请求

**登录接口示例：**

```http
POST /login
Content-Type: application/json

{
  "username": "admin",
  "password": "password"
}
```

**响应示例：**

```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.xxxxx",
  "user": {
    "id": 1,
    "username": "admin"
  }
}
```

## 📌 注意事项

- SQLite 在高并发场景下性能有限，生产环境建议替换为 MySQL 或 PostgreSQL。
- JWT 库已从 `dgrijalva/jwt-go` 迁移至官方维护的 `golang-jwt/jwt/v5`。
- 所有密钥应统一配置，避免多个地方使用不同值导致签名不一致。

## 🤝 贡献指南

欢迎提交 PR 和 Issue，请确保代码风格与项目保持一致，并遵循 Go 最佳实践。

## 📜 License

MIT License