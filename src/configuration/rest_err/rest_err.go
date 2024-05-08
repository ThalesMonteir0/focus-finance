package rest_err

import "net/http"

type RestErr struct {
	Message string  `json:"message"`
	Code    int     `json:"code"`
	Err     string  `json:"err"`
	Causes  []Cause `json:"causes,omitempty"`
}

type Cause struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (e *RestErr) Error() string {
	return e.Message
}

func NewBadRequestError(msg string) *RestErr {
	return &RestErr{
		Message: msg,
		Code:    http.StatusBadRequest,
		Err:     "bad_request",
	}
}

func NewBadRequestValidationError(message string, causes []Cause) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "bad_request",
		Code:    http.StatusBadRequest,
		Causes:  causes,
	}
}

func NewInternalServerError(msg string) *RestErr {
	return &RestErr{
		Message: msg,
		Code:    http.StatusInternalServerError,
		Err:     "bad_request",
	}
}

func NewUnauthorizedRequestError(msg string) *RestErr {
	return &RestErr{
		Message: msg,
		Code:    http.StatusUnauthorized,
		Err:     "bad_request",
	}
}

func NewNotFoundError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "not_found",
		Code:    http.StatusNotFound,
	}
}
