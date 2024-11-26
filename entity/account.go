package entity

import "example/internal/common/helper/sqlormhelper"

type Account struct {
	sqlormhelper.BaseEntity
	Id             string  `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	AccountName    string  `gorm:"column:acct_name;type:varchar;not null"`
	AccountNo      string  `gorm:"column:acct_no;type:varchar;not null"`
	AccountBalance float64 `gorm:"column:acct_bal;type:decimal(10,2);not null"`
}
