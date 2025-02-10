package handler

import (
	"net/http"

	"MACKOMUsers/internal/core"
	"MACKOMUsers/internal/frame"
)

type UserHandler struct {
	userStore core.UserStore
}

func NewUserHandler(userStore core.UserStore) *UserHandler {
	return &UserHandler{
		userStore: userStore,
	}
}

func (h *UserHandler) AddUser(request *http.Request) (*frame.HttpResponse, error) {
	user, status, err := processBody[core.User](request)
	if err != nil {
		return errorResponse(status, err)
	}

	if user.Firstname == "" || user.Lastname == "" || user.Age <= 0 {
		return errorResponse(http.StatusBadRequest, IncorrectBodyError)
	}

	err = h.userStore.AddUser(user)
	if err != nil {
		return errorResponse(http.StatusInternalServerError, InternalServerError)
	}

	return successResponse(SuccessResponse{
		Success: true,
	})
}

func (h *UserHandler) GetUserList(request *http.Request) (*frame.HttpResponse, error) {
	userList, err := h.userStore.GetUserList()
	if err != nil {
		return errorResponse(http.StatusInternalServerError, InternalServerError)
	}

	return successResponse(userList)
}

func (h *UserHandler) GetUserListByFilter(request *http.Request) (*frame.HttpResponse, error) {
	filter, status, err := processBody[GetUsersByFilterRequest](request)
	if err != nil {
		return errorResponse(status, err)
	}

	userList, err := h.userStore.GetUserListByFilter(filter.FromDate, filter.ToDate, filter.MinAge, filter.MaxAge)
	if err != nil {
		return errorResponse(http.StatusInternalServerError, InternalServerError)
	}

	return successResponse(userList)
}
