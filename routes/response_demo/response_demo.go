package response_demo

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/nerocho/fastfiber"
	"github.com/nerocho/fastfiber/utils/response"
)

func Hello(c *fiber.Ctx) error {
	// 原生response
	// fastfiber.Logger.Info().Str("a", "b").Msg("")    //这个是日志demo，不带message
	fastfiber.Logger.Info().Str("a", "b").Msg("这个是日志demo")

	return c.SendString("hello world!")

	// return c.JSON(fiber.Map{
	// 	"method": c.Method(),
	// 	"path":   c.Params("*"),
	// 	"data": fiber.Map{
	// 		"嵌套字段1": "value1",
	// 		"嵌套字段2": "value2",
	// 	},
	// })
}

func Error(c *fiber.Ctx) error {

	//直接抛出错误，让框架(fastfiber)自动处理,
	return errors.New("我就跑了个错误")

	//常见使用方式如下：
	// someresult,err := somefunc()
	// if err {
	// 	return err
	// }
	// do domeresult
}

//下面是封装后的response

func ResDemo1(c *fiber.Ctx) error {
	// return res.MakeRes(c).Msg("这是一段msg提示").Send()
	return response.MakeRes(c).Code(400).Msg("这是一段msg提示").Send() //基础用法
	//使用简写返回
	// return res.MakeRes(c).Send() // 返回一个成功提示
}

func ResDemo2(c *fiber.Ctx) error {
	// 根据struct 返回json
	type User struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	u := &User{Name: "nike", Age: 20}

	return response.MakeRes(c).Msg("success").Send(u)
}

func ResDemo3(c *fiber.Ctx) error {
	//根据map 返回json
	// return res.MakeRes(c).Msg("success").Send(fiber.Map{
	// 	"name": "nike",
	// 	"age":  20,
	// })
	// 也可以这样简写（默认为200和success）
	return response.MakeRes(c).Send(fiber.Map{
		"name":  "nike",
		"age":   20,
		"phone": "188888888888",
	})
}

func ResDemo4(c *fiber.Ctx) error {
	//使用简写返回
	return response.MakeRes(c).SendSuccess(fiber.Map{
		"name": "nike",
		"age":  20,
		"from": c.Path(),
	})
}

func ResDemo5(c *fiber.Ctx) error {
	//使用简写返回
	return response.MakeRes(c).SendSuccess("覆盖默认消息提示", fiber.Map{
		"name": "nike",
		"age":  20,
		"from": c.Path(),
	})

	// 这两个也同理
	// SendFail
	// SendWarn
}
