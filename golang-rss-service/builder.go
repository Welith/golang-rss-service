package main

//BuildErrorResponse to build identical error responses
func BuildErrorResponse(errorString string, errorMessage string) *ErrorResponseStruct {

	errorMsg := new(ErrorResponseStruct)
	errorMsg.Error = errorString
	errorMsg.ErrorMsg = errorMessage

	return errorMsg
}
