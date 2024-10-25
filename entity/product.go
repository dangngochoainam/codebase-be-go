package entity

import "example/internal/common/helper/sqlormhelper"

type Product struct {
	sqlormhelper.BaseEntity
	Id   string `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Name string `gorm:"column:user_name;type:varchar;not null"`
}
