package middleware

import "errors"

type AppError struct {
	StatusCode int
	Message    string
	Err        error
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return e.Err.Error()
	}
	return e.Message
}

func NewAppError(statusCode int, message string, err error) *AppError {
	if err == nil {
		err = errors.New(message)
	}
	return &AppError{
		StatusCode: statusCode,
		Message:    message,
		Err:        err,
	}
}

func IsAppError(err error) (*AppError, bool) {
	var appErr *AppError
	ok := errors.As(err, &appErr)
	return appErr, ok
}
