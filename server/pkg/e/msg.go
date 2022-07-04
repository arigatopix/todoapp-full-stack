package e

// รับ SUCCESS จาก code.go
var MessageFlags = map[int]string{
	SUCCESS:                "ok",
	ERROR:                  "fail",
	INVALID_PARAMS:         "Invalid parameter",
	ERROR_ADD_TODO_FAIL:    "Can not add todos",
	ERROR_UPDATE_TODO:      "Fail to update todo",
	ERROR_TODO_NOT_EXIST:   "Todo does not exist",
	ERROR_DELETE_TODO_FAIL: "Delete todo fail",
}

func GetMessage(errCode int) string {
	msg, ok := MessageFlags[errCode]
	if ok {
		return msg
	}

	return MessageFlags[ERROR]
}
