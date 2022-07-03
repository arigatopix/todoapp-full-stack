package e

// รับ SUCCESS จาก code.go
var MessageFlags = map[int]string{
	SUCCESS:             "ok",
	ERROR:               "fail",
	INVALID_PARAMS:      "Invalid parameter",
	ERROR_ADD_TODO_FAIL: "Can not add todos",
}

func GetMessage(errCode int) string {
	msg, ok := MessageFlags[errCode]
	if ok {
		return msg
	}

	return MessageFlags[ERROR]
}
