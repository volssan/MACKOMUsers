package frame

import (
	"net/http"

	"github.com/go-chi/render"
)

type HttpResponse struct {
	Headers Headers
	Data    []byte
	Code    int
}

type errorResponse struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

func setErrorResponse(err error, code int, writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(code)

	render.JSON(writer, request, errorResponse{
		Code:  code,
		Error: err.Error(),
	})
}
