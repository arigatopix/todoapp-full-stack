package e

// รับ SUCCESS จาก code.go
var MessageFlags = map[int]string{
	SUCCESS:                "ok",
	ERROR:                  "fail",
	INVALID_PARAMS:         "Invalid field or value",
	ERROR_ADD_TODO_FAIL:    "Can not add todos",
	ERROR_UPDATE_TODO:      "Fail to update todo",
	ERROR_TODO_NOT_EXIST:   "Todo does not exist",
	ERROR_DELETE_TODO_FAIL: "Delete todo fail",

	ERROR_REGISTER_USER:  "Can not register user with email",
	ERROR_USER_NOT_EXIST: "User does not exist",

	ERROR_USER_EXISTED:             "Email used. Please enter new email or log in instaed",
	ERROR_AUTH_TOKEN:               "Can not generate token",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Please login again",
	ERROR_AUTH_CHECK_TOKEN_FAIL:    "Token does not match unauthorized",
}

func GetMessage(errCode int) string {
	msg, ok := MessageFlags[errCode]
	if ok {
		return msg
	}

	return MessageFlags[ERROR]
}
