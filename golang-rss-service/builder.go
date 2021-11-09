package main

func BuildErrorResponse(errorString string, errorMessage string) *ErrorResponseStruct {

	errorMsg := new(ErrorResponseStruct)
	errorMsg.Error = errorString
	errorMsg.ErrorMsg = errorMessage

	return errorMsg
}
