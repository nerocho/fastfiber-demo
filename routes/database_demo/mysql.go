package database_demo

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nerocho/fastfiber"
	"github.com/nerocho/fastfiber/utils/response"

	"fastfiber-demo/models"
)

func Select1(c *fiber.Ctx) error {
	var users []models.UserModel
	result := fastfiber.Db.Find(&users)
	if result.Error != nil {
		return response.MakeRes(c).SendFail(result.Error.Error())
	}
	return response.MakeRes(c).SendSuccess(users)

	//懒得写了，其他数据库操作，看下gorm的文档fastfiber.Db 为 *gorm.Db指针
	// fastfiber.Db.Create()
}
