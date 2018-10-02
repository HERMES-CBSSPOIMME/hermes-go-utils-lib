package customhttpresponse

import (
	"hermes-go-utils-lib/logger"
	http "net/http"
)

const (
	SUCCESS_AUTH_CODE = "SUCCESS-AUTH"
	SUCCESS_AUTH_MSG  = "Authentication Successful !"
)

var (
	mapping = map[string]CustomHTTPResponse{

		SUCCESS_AUTH_CODE: CustomHTTPResponse{
			StatusCode: http.StatusOK,
			Headers: []http.Header{
				http.Header{"Content-Type", "application/json"},
			},
			Body: CustomHTTPResponseBody{
				LoggingID: "",
				Code:      SUCCESS_AUTH_CODE,
				Message:   SUCCESS_AUTH_MSG,
			},
		},
	}
)

// ProxyRequest : Received Request from Client
type CustomHTTPResponse struct {
	StatusCode int
	Headers    []http.Header
	Body       CustomHTTPResponseBody
}

type CustomHTTPResponseBody struct {
	LogInfos logger.LogInfos `json:"logInfos"`
	Code     string          `json:"code"`
	Message  string          `json:"message"`
	Content  interface{}     `json:"content"`
}
