# Go Web Starter

一个**可直接运行、可持续扩展**的 Go Web 项目基础模板，适合你这种：
> 有其他语言经验 → 想最小成本上 Go → 先上线再慢慢完善

当前已集成：
- Gin（HTTP / 路由 / 中间件）
- GORM（MySQL ORM）
- Validator（参数校验）
- JWT（登录鉴权）
- Zap（结构化日志）
- Swagger（接口文档）

项目目标不是“炫技”，而是：**工程清晰、能跑、好扩展**。

---

## 一、项目结构说明

```text
go-web-starter
├── main.go                # 程序入口
├── go.mod
├── config
│   └── config.yaml        # 配置文件（端口 / DB / JWT）
├── internal               # 业务代码（不对外暴露）
│   ├── api                # HTTP Handler（只做参数 + 返回）
│   ├── service            # 业务逻辑层
│   ├── dao                # 数据访问层（GORM）
│   ├── model              # 数据模型
│   └── middleware         # Gin 中间件（JWT / 日志）
├── pkg                    # 基础设施 / 工具包
│   ├── db                 # 数据库初始化
│   ├── jwt                # JWT 生成
│   ├── logger             # Zap 日志初始化
│   └── response            # 统一响应结构
├── docs                   # Swagger 自动生成文件
└── README.md
```

### 分层原则

- **api**：只处理 HTTP 相关（参数校验 / 返回 JSON）
- **service**：业务逻辑，不依赖 Gin
- **dao**：数据库操作，只关心数据
- **middleware**：鉴权 / 日志等横切逻辑
- **pkg**：基础能力，尽量与业务无关

---

## 二、环境要求

- Go >= 1.20（推荐 1.21+）
- MySQL >= 5.7 / 8.x
- 已开启 Go Modules（默认）

---

## 三、数据库准备

### 1️⃣ 创建数据库和表

数据库名：`go_web_demo`

```sql
CREATE DATABASE IF NOT EXISTS go_web_demo DEFAULT CHARSET utf8mb4;
USE go_web_demo;

CREATE TABLE user (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    username VARCHAR(64) NOT NULL UNIQUE,
    password VARCHAR(128) NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO user (username, password)
VALUES ('admin', '123456');
```

> 说明：
> - 当前为演示用明文密码
> - 实际项目中建议改为 bcrypt（后续可扩展）

---

## 四、配置说明

### `config/config.yaml`

```yaml
server:
  port: 9001

db:
  dsn: root:root123@tcp(raspberrypi:3306)/go_web_demo?charset=utf8mb4&parseTime=True&loc=UTC

jwt:
  secret: go-web-secret
  expire: 7200   # token 有效期（秒）
```

请根据你的环境修改：
- MySQL 地址
- 用户名 / 密码

---

## 五、安装依赖

在项目根目录执行：

```bash
go mod tidy
```

如果你是第一次使用 Swagger：

```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

生成接口文档：

```bash
swag init
```

---

## 六、启动项目

```bash
go run main.go
```

启动成功后，你会看到：
- Zap JSON 日志输出
- Gin 服务监听端口

默认地址：

```
http://localhost:9001/api
```

构建二进制：
```bash
GOOS=linux GOARCH=arm64 go build -o ./bin/go-web-starter main.go
```

---

## 七、接口说明与测试

### 1️⃣ Swagger 文档

浏览器访问：

```
http://localhost:9001/swagger/index.html
```

可以直接在线调试接口。

---

### 2️⃣ 登录接口（获取 JWT）

**POST** `/api/login`

请求示例：

```json
{
  "username": "admin",
  "password": "123456"
}
```

响应示例：

```json
{
  "code": 0,
  "msg": "ok",
  "data": {
    "token": "xxxxx.yyyyy.zzzzz"
  }
}
```

---

### 3️⃣ 用户列表（需要 JWT）

**GET** `/api/users`

请求头：

```
Authorization: <token>
```

---

### 4️⃣ 创建用户（需要 JWT）

**POST** `/api/users`

```json
{
  "username": "test",
  "password": "123456"
}
```

---

## 八、日志说明

- 使用 Zap 输出 JSON 日志
- 每个 HTTP 请求都会记录：
  - method
  - path
  - status
  - 耗时
  - IP

适合直接接入：
- ELK
- Loki
- 云日志系统

---

## 九、后续可扩展方向（推荐顺序）

1. 登录密码改为 bcrypt
2. JWT 中解析 userID 写入 context
3. 增加 request_id / trace_id
4. 用户权限 / RBAC
5. 分页 / 错误码体系

---

## 十、一句设计理念

> 这个项目不是“完美模板”，
> 而是一个**你能真正长期用、不断往上加东西的起点**。

如果你在此基础上继续演进，
它会自然成长为一个成熟的 Go Web 服务。

