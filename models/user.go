package models

import (
	"github.com/nerocho/fastfiber/utils/orm"
)

type UserModel struct {
	ID         int64         `gorm:"primaryKey;column:id;type:bigint;not null" json:"id"`
	Number     string        `gorm:"unique;column:number;type:varchar(255);not null;default:''" json:"number"` // 学号
	Name       string        `gorm:"column:name;type:varchar(255);not null;default:''" json:"name"`            // 用户名称
	Password   string        `gorm:"column:password;type:varchar(255);not null;default:''" json:"password"`    // 用户密码
	Gender     string        `gorm:"column:gender;type:char(5);not null" json:"gender"`                        // 男｜女｜未公开
	CreateTime orm.LocalTime `gorm:"column:create_time;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"createTime"`
	UpdateTime orm.LocalTime `gorm:"column:update_time;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"updateTime"`
}

// 自定义表名
func (c *UserModel) TableName() string {
	return "user"
}
