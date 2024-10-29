package commonhelper

type HeaderKeyType string

const (
	HeaderKeyType_TraceId HeaderKeyType = "TraceId"
	HeaderKeyType_TimeMs  HeaderKeyType = "TimeMs"
)

type ContextKey string

const (
	ContextKey_User ContextKey = "User"
	ContextKey_Key  ContextKey = "Key"
)
