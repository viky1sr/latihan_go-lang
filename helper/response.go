package helper

import "strings"

type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Error   interface{} `json:"errors"`
	Data    interface{} `json:"data"`
}

// EmptyObj object is used when data doesn't want to be null on json
type EmptyObj struct{}

// BuildResponse method to inject response success
func BuildResponse(status bool, message string, data interface{}) Response {
	res := Response{
		Status:  status,
		Message: message,
		Error:   nil,
		Data:    data,
	}
	return res
}

// BuildErrorResponse method to inject response failed
func BuildErrorResponse(message string, err string, data interface{}) Response {
	splittedError := strings.Split(err, "\n")
	res := Response{
		Status:  false,
		Message: message,
		Error:   splittedError,
		Data:    data,
	}
	return res
}
