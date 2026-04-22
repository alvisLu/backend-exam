package httpx

import "net/http"

type HTTPError struct {
	Status  int
	Message string
}

func (e *HTTPError) Error() string { return e.Message }

func BadRequest(msg string) *HTTPError { return &HTTPError{Status: http.StatusBadRequest, Message: msg} }
func NotFound(msg string) *HTTPError   { return &HTTPError{Status: http.StatusNotFound, Message: msg} }
