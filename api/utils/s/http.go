package s

import (
	"net/http"

	"api/utils/r"
)

type Result struct {
	Code         int         `json:"code"`
	Data         interface{} `json:"data,omitempty"`
	ErrorMessage string      `json:"error_message,omitempty"`
}

func NewResultData(data interface{}) *Result {
	return &Result{
		Code:         http.StatusOK,
		Data:         data,
		ErrorMessage: "",
	}
}

func NewResultMessage(message string) *Result {
	return &Result{
		Code:         http.StatusOK,
		Data:         message,
		ErrorMessage: "",
	}
}

func NewResultError(err error) *Result {
	code := http.StatusInternalServerError
	if eR, ok := err.(r.Ex); ok {
		if eR.IsDataNotFound() {
			code = 512
		} else if eR.IsDatabaseError() {
			code = 513
		} else if eR.IsRepositoryError() {
			code = 514
		} else if eR.IsUseCaseError() {
			code = 515
		}
	}
	return &Result{
		Code:         code,
		Data:         nil,
		ErrorMessage: err.Error(),
	}
}
