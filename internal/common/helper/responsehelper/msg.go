package responsehelper

var MsgFlags = map[SystemCode]string{
	SUCCESS:        "ok",
	ERROR:          "fail",
	INVALID_PARAMS: "invalid params",
}

// GetMsg get error information based on Code
func GetMsg(systemCode SystemCode) string {
	msg, ok := MsgFlags[systemCode]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
