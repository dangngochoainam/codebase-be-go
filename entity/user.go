package entity

type User struct {
	Id       string `gorm:primaryKey;type:uuid;default:gen_random_uuid()`
	Username string `gorm:column:user_name;type:varchar;not null`
	Password string `gorm:column:password;type:varchar;not null`
}
