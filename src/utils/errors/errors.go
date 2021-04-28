package errors

import "net/http"

//ApiError implements error interface from builtin
type ApiError struct {
	Reason string `json:"reason,omitempty"`
	Msg    string `json:"message,omitempty"`
	Code   int    `json:"code,omitempty"`
}

func (u *ApiError) Error() string {
	return u.Msg
}

func NewBadRequestError(msg, reason string) *ApiError {
	return &ApiError{
		Reason: reason,
		Msg:    msg,
		Code:   http.StatusBadRequest,
	}
}

func NewNotFoundError(msg, reason string) *ApiError {
	return &ApiError{
		Reason: reason,
		Msg:    msg,
		Code:   http.StatusNotFound,
	}
}

func NewInternalServerError(msg, reason string) *ApiError {
	return &ApiError{
		Reason: reason,
		Msg:    msg,
		Code:   http.StatusInternalServerError,
	}
}
