## 一个史诗级论坛
### 技能清单
##### 1. 雪花算法
##### 2. gin框架
##### 3. zap日志库
##### 4. Go语言操作MySQL **(sqlx)**
##### 5. Go语言操作Redis **(go-redis)**
##### 6. Viper配置管理
##### 7. JWT认证
##### 8. swagger生成文档
##### 9. 令牌桶限流
##### 10. Dockerfile docker-compose

### 后端结构树
```bash
.
├── Dockerfile
├── Makefile
├── config
│   └── config.yaml
├── controller
│   ├── code.go
│   ├── community.go
│   ├── docs_models.go
│   ├── post.go
│   ├── post_test.go
│   ├── request.go
│   ├── response.go
│   ├── user.go
│   ├── validator.go
│   └── vote.go
├── dao
│   ├── mysql
│   │   ├── community.go
│   │   ├── err_code.go
│   │   ├── mysql.go
│   │   ├── post.go
│   │   ├── post_test.go
│   │   └── user.go
│   └── redis
│       ├── err_code.go
│       ├── keys.go
│       ├── post.go
│       ├── redis.go
│       └── vote.go
├── docker-compose.yaml
├── docs
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── go.mod
├── go.sum
├── init.sql
├── log
│   └── bluebell.log
├── logger
│   └── logger.go
├── main.go
├── middleware
│   ├── auth.go
│   └── ratelimit.go
├── model
│   ├── community.go
│   ├── param.go
│   ├── post.go
│   └── user.go
├── pkg
│   ├── crypto
│   │   └── crypto.go
│   ├── jwt
│   │   └── jwt.go
│   └── snowflake
│       └── gen_id.go
├── router
│   └── router.go
├── service
│   ├── community.go
│   ├── post.go
│   ├── user.go
│   └── vote.go
├── setting
│   └── setting.go
├── static
│   ├── css
│   │   └── app.f4cf2413.css
│   ├── favicon.ico
│   ├── img
│   │   ├── avatar.7b0a9835.png
│   │   ├── iconfont.cdbe38a0.svg
│   │   ├── logo.da56125f.png
│   │   └── search.8e85063d.png
│   └── js
│       ├── app.1deeed5e.js
│       ├── app.1deeed5e.js.map
│       ├── chunk-vendors.47138da1.js
│       └── chunk-vendors.47138da1.js.map
├── templates
│   └── index.html
└── wait-for.sh
```
### 前端略有错误
