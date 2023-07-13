package i18n

const defaultCode = 3

type CodeError struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
}

type CodeErrorResponse struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
}

func NewCodeError(code int64) error {
	return &CodeError{Code: code, Msg: ""}
}

func NewCodeErrorWithMsg(code int64, msg string) error {
	return &CodeError{Code: code, Msg: msg}
}

func NewCodeCanceledError(msg string) error {
	return &CodeError{Code: 1, Msg: msg}
}

func NewCodeInvalidArgumentError(msg string) error {
	return &CodeError{Code: 3, Msg: msg}
}

func NewCodeNotFoundError(msg string) error {
	return &CodeError{Code: 5, Msg: msg}
}

func NewCodeAlreadyExistsError(msg string) error {
	return &CodeError{Code: 6, Msg: msg}
}

func NewCodeAbortedError(msg string) error {
	return &CodeError{Code: 10, Msg: msg}
}

func NewCodeInternalError(msg string) error {
	return &CodeError{Code: 13, Msg: msg}
}

func NewCodeUnavailableError(msg string) error {
	return &CodeError{Code: 14, Msg: msg}
}

func NewDefaultError(msg string) error {
	return NewCodeErrorWithMsg(defaultCode, msg)
}

func (e *CodeError) Error() string {
	return e.Msg
}

func (e *CodeError) GetCode() int64 {
	return e.Code
}

func (e *CodeError) Data() *CodeErrorResponse {
	return &CodeErrorResponse{
		Code: e.Code,
		Msg:  e.Msg,
	}
}
