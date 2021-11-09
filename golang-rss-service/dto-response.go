package main

type ResponseItem struct {
	Items []interface{} `json:"items,omitempty"`
}

type ErrorResponseStruct struct {
	Error    string `json:"error"`
	ErrorMsg string `json:"error_msg"`
}
