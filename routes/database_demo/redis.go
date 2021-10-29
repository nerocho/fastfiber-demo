package database_demo

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/nerocho/fastfiber"
	"github.com/nerocho/fastfiber/utils/response"
)

func Select2(c *fiber.Ctx) error {

	id, err := fastfiber.IdWorker.NextId()
	if err != nil {
		return err //不处理错误，直接让fastfiber 捕获
	}

	//存一个id到redis ，并设置一小时有效期
	fastfiber.RedisPool.SetNX(c.UserContext(), "somerediskey", id, time.Hour)

	return response.MakeRes(c).Msg("redis 都了解吧。。。").Send()
}
