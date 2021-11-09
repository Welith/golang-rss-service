package main

import (
	golang_rss_reader_package "github.com/Welith/golang-rss-reader-package"
	validation "github.com/go-ozzo/ozzo-validation"
)

type ResponseItem struct {
	Items []golang_rss_reader_package.RssItem `json:"items"`
}

type ErrorResponseStruct struct {
	Error    string `json:"error"`
	ErrorMsg string `json:"error_msg"`
}

type Request struct {
	Urls []string `json:"urls"`
}

func (request Request) Validate() error {

	return validation.ValidateStruct(&request,
		validation.Field(&request.Urls, validation.Required),
	)
}
