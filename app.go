package main

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/nerocho/fastfiber"
	"github.com/nerocho/fastfiber/middleware"
	"github.com/nerocho/fastfiber/utils/response"

	"fastfiber-demo/routes"
)

func main() {
	//初始化数据库等资源，绑定配置
	fastfiber.Bootstrap()
	//fiber
	app := fiber.New(fiber.Config{
		Prefork:      false,
		ProxyHeader:  "X-Forwarded-For", //如果有前置代理，则使用该字段覆盖c.IP()
		AppName:      fastfiber.Conf.GetString("System.AppName"),
		ErrorHandler: fastfiber.ErrorHandler,
	})
	// recover
	app.Use(recover.New())
	// 限流
	if fastfiber.Conf.GetBool("Limiter.Enable") {
		whiteIpList := fastfiber.Conf.GetStringSlice("Limiter.IpWhiteList")
		app.Use(limiter.New(limiter.Config{
			Max:        100,             //请求上限
			Expiration: 1 * time.Minute, // 缓存时间也是限制时间
			Next: func(c *fiber.Ctx) bool {
				f := false
				cur := c.IP()
				for i := 0; i < len(whiteIpList); i++ {
					if cur == whiteIpList[i] {
						f = true
						break
					}
				}
				return f
			},
			LimitReached: func(c *fiber.Ctx) error {
				return response.MakeRes(c).Code(429).Msg("请求频率过高").Send()
			},
		}))
	}
	// 跨域配置
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*", // 测试
	}))
	//访问日志
	app.Use(middleware.Access(&middleware.AcccessOptions{
		LogResponse: fastfiber.Conf.GetBool("System.LogResponseBody"),
	}))

	//缓存
	if fastfiber.Conf.GetBool("ApiCache.Enable") {
		app.Use(cache.New(cache.Config{
			Next: func(c *fiber.Ctx) bool {
				return c.Query("refresh") == "true"
			},
			Expiration: 120 * time.Second,
			// KeyGenerator: func(c *fiber.Ctx) string
			Storage: nil,
		}))
	}

	//绑定路由
	routes.SetupRoutes(app)
	// 处理404
	app.Use(middleware.NotFound())
	// 启动应用
	fastfiber.GraceRun(app, 10)
}
