package params_demo

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nerocho/fastfiber/utils/response"
)

// 路由参数demo
// "/:value",
func GetParams1(c *fiber.Ctx) error {
	return response.MakeRes(c).Send("value: " + c.Params("value"))
}

// GET参数获取方法一 Query
// ?order=desc&brand=nike
func GetParams2(c *fiber.Ctx) error {
	return response.MakeRes(c).Send(fiber.Map{
		"order": c.Query("order"),         // "desc"
		"brand": c.Query("brand"),         // "nike"
		"empty": c.Query("empty", "nike"), // 注意这里，empty没有传参时，可以设置默认值
	})

}

// GET参数获取方法二 QueryParser !推荐
// ?name=john&pass=doe&products=shoe,hat"
func GetParams3(c *fiber.Ctx) error {
	p := new(PersonGetRequest)
	if err := c.QueryParser(p); err != nil {
		return err
	}
	return response.MakeRes(c).Msg("把收到的参数返回去").Send(p)
}
