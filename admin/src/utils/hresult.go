package utils

type HttpResult struct {
	Code    int
	Message string
	Err     error
}

func NewHttpResult(code int, message string, err error) HttpResult {
	return HttpResult{
		Code:    code,
		Message: message,
		Err:     err,
	}
}

func (result HttpResult) Error() string {
	return result.Err.Error()
}