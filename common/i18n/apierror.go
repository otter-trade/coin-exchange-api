package i18n

import (
	"google.golang.org/grpc/status"
	"net/http"
)

type SimpleMsg struct {
	Msg string `json:"msg"`
}

type ApiError struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
}

func (e *ApiError) Error() string {
	return e.Msg
}

func (e *ApiError) GetCode() int64 {
	return e.Code
}

func NewApiError(code int64) error {
	return &ApiError{Code: code, Msg: ""}
}

func NewApiErrorWithMsg(code int64, msg string) error {
	return &ApiError{Code: code, Msg: msg}
}

func NewApiErrorWithoutMsg(code int64) error {
	return &ApiError{Code: code, Msg: ""}
}

func NewApiInternalError(msg string) error {
	return &ApiError{Code: http.StatusInternalServerError, Msg: msg}
}

func NewApiBadRequestError(msg string) error {
	return &ApiError{Code: http.StatusBadRequest, Msg: msg}
}

func NewApiUnauthorizedError(msg string) error {
	return &ApiError{Code: http.StatusUnauthorized, Msg: msg}
}

func NewApiForbiddenError(msg string) error {
	return &ApiError{Code: http.StatusForbidden, Msg: msg}
}

func NewApiNotFoundError(msg string) error {
	return &ApiError{Code: http.StatusNotFound, Msg: msg}
}

func NewApiBadGatewayError(msg string) error {
	return &ApiError{Code: http.StatusBadGateway, Msg: msg}
}

func NewRpcError(err error) error {
	s, ok := status.FromError(err)
	var code int64 = ServerError
	if ok {
		code = int64(s.Code())
	}
	return &ApiError{Code: code, Msg: CodeToMsg(code)}
}
