package validate_demo

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/nerocho/fastfiber"
	"github.com/nerocho/fastfiber/utils/response"
)

type UserForm struct {
	Name     string    `json:"name" form:"name" validate:"required|minLen:4"`
	Email    string    `json:"email" form:"email" validate:"email" message:"email is invalid"`
	Age      int       `json:"age" form:"age" validate:"required|int|min:1|max:99" message:"int:年龄必须为整数| min: 年龄最小值为1"` //自定义消息
	CreateAt int       `json:"createAt" form:"createAt" validate:"min:1"`
	Safe     int       `json:"safe" form:"safe" validate:"-"` //!不验证
	UpdateAt time.Time `json:"updateAt" form:"updateAt" validate:"required"`
	Code     string    `json:"code" form:"code" validate:"customValidator|required" message:"code必须为4位"` // !使用自定义验证器
	// 结构体嵌套
	ExtInfo struct {
		Homepage string `json:"homePage" form:"homePage" validate:"required"`
		CityName string `json:"cityName" form:"cityName" validate:"required"`
	} `json:"extInfo" form:"extInfo" validate:"required"`
}

// CustomValidator 定义在结构体中的自定义验证器
func (f UserForm) CustomValidator(val string) bool {
	return len(val) == 4
}

// 字段翻译demo
func (f UserForm) Translates() map[string]string {
	return map[string]string{
		"Name":     "姓名",
		"Email":    "邮箱",
		"Age":      "年龄",
		"CreateAt": "创建时间",
		"UpdateAt": "更新时间",
		"Code":     "验证码",
		"ExtInfo":  "其他",
		"Homepage": "主页",
		"CityName": "城市",
	}
}

// 前端参数校验demo
// 使用 validator.BindAndAllErr 绑定并验证，BindAndAllErr会自动根据请求类型，进行参数绑定
func PostRequestValid(c *fiber.Ctx) error {

	u := UserForm{}

	// 绑定并返回错误信息
	// err := validator.BindAndOneErr(c, &u) // 返回一个错误，错误无顺序（随机）
	err := fastfiber.BindAndAllErr(c, &u) // 返回所有错误

	// 返回验证错误信息
	if err != nil {
		return err
	}

	//返回验证过的数据
	response.MakeRes(c).SendSuccess(u)

	return nil
}
