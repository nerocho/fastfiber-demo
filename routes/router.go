package routes

import (
	"github.com/gofiber/fiber/v2"

	"fastfiber-demo/routes/database_demo"
	"fastfiber-demo/routes/params_demo"
	"fastfiber-demo/routes/response_demo"
	"fastfiber-demo/routes/validate_demo"
)

func SetupRoutes(app *fiber.App) {

	// Group demo
	// Auth
	// auth := api.Group("/auth")
	// auth.Post("/login", handler.Login)

	// v1 := api.Group("/api/v1")
	// v1 ...

	// 下面是所有的接口demo
	api := app.Group("/api")

	// 如何返回数据给前端DEMO
	api.Get("/", response_demo.Hello)
	api.Get("/error", response_demo.Error)
	api.Get("/resdemo1", response_demo.ResDemo1)
	api.Get("/resdemo2", response_demo.ResDemo2)
	api.Get("/resdemo3", response_demo.ResDemo3)
	api.Get("/resdemo4", response_demo.ResDemo4)
	api.Get("/resdemo5", response_demo.ResDemo5)

	// 如何校验前端数据 DEMO
	api.Post("/PostRequestValid", validate_demo.PostRequestValid)

	// 如何获取Get请求的参数
	api.Get("/params1/:value", params_demo.GetParams1)
	api.Get("/params2", params_demo.GetParams2)
	api.Get("/params3", params_demo.GetParams3)

	// 如何获取Post请求的参数
	api.Post("/params4", params_demo.PostParams1)
	api.Post("/params5", params_demo.PostParams2)

	//数据库相关demo
	api.Get("/select1", database_demo.Select1)
	api.Get("/select2", database_demo.Select2)

}
