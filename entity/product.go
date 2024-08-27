package entity

type Product struct {
	Id   string `gorm:primaryKey;type:uuid;default:gen_random_uuid()`
	Name string `gorm:column:user_name;type:varchar;not null`
}
