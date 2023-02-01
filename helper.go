package fiberutils

type Map map[string]interface{}

func NewRequestError(statusCode int, message string, err error) *RequestError {
	return &RequestError{
		StatusCode: statusCode,
		Message:    message,
		Err:        err,
	}
}
