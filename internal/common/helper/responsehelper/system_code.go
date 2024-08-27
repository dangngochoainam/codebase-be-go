package responsehelper

type SystemCode int

const (
	SUCCESS        SystemCode = 200
	ERROR          SystemCode = 500
	INVALID_PARAMS SystemCode = 400
)
