# fastfiber-demo

fastfiber-demo 是[fastfiber](https://github.com/nerocho/fastfiber) 的 demo 项目

## 项目结构

项目结构参照`expressjs`组织

```bash
.
├── README.md
├── app.go // 入口文件
├── config.development.yml // 配置文件
├── go.mod
├── go.sum
├── models // ORM 模型文件
│   └── user.go
└── routes // 路由文件
    ├── database_demo // 数据库demo
    │   ├── mysql.go
    │   └── redis.go
    ├── params_demo // 请求及参数demo
    │   ├── get.go
    │   ├── post.go
    │   └── request.go
    ├── response_demo // response 示例
    │   └── response_demo.go
    ├── router.go // 路由绑定及入口
    └── validate_demo // 参数自动校验demo
        └── validate_demo.go

```
