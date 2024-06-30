package rest_err

import (
	"net/http"
)

// Anexa a função Error implementada abaixo
type RestErr struct {
	Message string   `json:"message"`
	Err     string   `json:"err"`
	Code    int      `json:"code"`
	Causes  []Causes `json:"causes"`
}

type Causes struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// Função anexada ao RestErr
// A variável Error do Go é do tipo interface e conseguimos implementar utilizando um objeto nosso
func (r *RestErr) Error() string {
	return r.Message
}

// func NewBadRequestValidationError(message string, causes ...Causes) *RestErr {
// 	return &RestErr{
// 		Message: message,
// 		Err:     "bad_request",
// 		Code:    http.StatusBadRequest,
// 		Causes:  nil,
// 	}
// }

func NewBadRequestError(message string, causes ...Causes) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "bad_request",
		Code:    http.StatusBadRequest,
		Causes:  causes,
	}
}

func NewInternalServerError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "internal_server",
		Code:    http.StatusInternalServerError,
		Causes:  nil,
	}
}

func NewNotFoundErrorError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "not_found",
		Code:    http.StatusNotFound,
		Causes:  nil,
	}
}
