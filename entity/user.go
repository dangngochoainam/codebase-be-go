package entity

import "example/internal/common/helper/sqlormhelper"

type User struct {
	sqlormhelper.BaseEntity
	Id       string `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Username string `gorm:"column:user_name;type:varchar;not null"`
	Password string `gorm:"column:password;type:varchar;not null"`
	Email    string `gorm:"column:email;type:varchar"`
	Age      int    `gorm:"column:age;type:int;not null"`
	IsActive bool   `gorm:"column:is_active;type:tinyint(1);not null;default:0"`
}
