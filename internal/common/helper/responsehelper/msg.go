package responsehelper

var MsgFlags = map[SystemCode]string{
	SUCCESS:        "ok",
	ERROR:          "fail",
	INVALID_PARAMS: "请求参数错误",
}

// GetMsg get error information based on Code
func GetMsg(systemCode SystemCode) string {
	msg, ok := MsgFlags[systemCode]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
