package handler

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/pkg/errors"

	"MACKOMUsers/internal/frame"
)

var (
	InternalServerError        = errors.New("Внутренняя ошибка сервера")
	IncorrectBodyError         = errors.New("Переданы некорректные параметры запроса")
	RequestBodyIsRequiredError = errors.New("Параметры запроса не переданы")
)

func processBody[T any](request *http.Request) (T, int, error) {
	var result T

	body, err := io.ReadAll(request.Body)
	if err != nil {
		return result, http.StatusInternalServerError, InternalServerError
	}

	if len(body) <= 0 {
		return result, http.StatusBadRequest, RequestBodyIsRequiredError
	}

	if len(body) > 0 {
		err = json.Unmarshal(body, &result)
		if err != nil {
			return result, http.StatusBadRequest, IncorrectBodyError
		}
	}

	return result, http.StatusOK, nil
}

func errorResponse(statusCode int, err error) (*frame.HttpResponse, error) {
	return &frame.HttpResponse{Code: statusCode}, err
}

func successResponse(data any) (*frame.HttpResponse, error) {
	marshaled, err := json.Marshal(data)
	if err != nil {
		return nil, errors.Wrap(err, "Marshal")
	}

	return &frame.HttpResponse{
		Data: marshaled,
	}, nil
}
