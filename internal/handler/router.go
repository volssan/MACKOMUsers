package handler

import (
	"context"

	"github.com/pkg/errors"

	"MACKOMUsers/internal/adapter/store"
	"MACKOMUsers/internal/config"
	"MACKOMUsers/internal/frame"
)

const (
	prefix = "/api/v1"
)

func InitRouter(_ context.Context, app *frame.App, cfg *config.Config) error {
	dbStore, err := store.New(cfg.Database)
	if err != nil {
		return errors.Wrap(err, "store.New")
	}

	userHandler := NewUserHandler(dbStore)

	app.RegisterHttpHandler(frame.Post, prefix+"/user/add", userHandler.AddUser)
	app.RegisterHttpHandler(frame.Get, prefix+"/user/list", userHandler.GetUserList)
	app.RegisterHttpHandler(frame.Post, prefix+"/user/filter", userHandler.GetUserListByFilter)

	return nil
}
