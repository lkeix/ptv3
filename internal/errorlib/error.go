package errorlib

import "errors"

// ErrorStruct is error code & messsage struct
var ErrorStruct map[string]error = map[string]error{
	"400": errors.New("400 Bad Request. your request is bad request."),
	"403": errors.New("403 Forbidden. you don't have certification to access this page."),
	"404": errors.New("404 Not Found. this page doesn't exist."),
	"500": errors.New("500 Internal Error."),
}
