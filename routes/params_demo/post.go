package params_demo

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/nerocho/fastfiber/utils/response"
)

//这里是POST

// POST参数获取方法一
// curl -X POST http://localhost:8080 -d user=john
func PostParams1(c *fiber.Ctx) error {
	// []byte("user=john")
	return response.MakeRes(c).Send(fmt.Sprintf("%s", c.Body()))
}

// POST参数获取方法二
func PostParams2(c *fiber.Ctx) error {
	p := new(PersonPostRequest)
	//B odyParser会自动根据Content-Type解析请求body
	if err := c.BodyParser(p); err != nil {
		return err
	}
	return response.MakeRes(c).Msg("把收到的参数返回去").Send(p)
}

// curl -X POST -H "Content-Type: application/json" --data "{\"name\":\"john\",\"pass\":\"doe\"}" localhost:3000
// curl -X POST -H "Content-Type: application/xml" --data "<login><name>john</name><pass>doe</pass></login>" localhost:3000
// curl -X POST -H "Content-Type: application/x-www-form-urlencoded" --data "name=john&pass=doe" localhost:3000
// curl -X POST -F name=john -F pass=doe http://localhost:3000
// curl -X POST "http://localhost:3000/?name=john&pass=doe"
