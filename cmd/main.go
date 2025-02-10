package main

import (
	"log"

	"github.com/pkg/errors"

	"MACKOMUsers/internal/config"
	"MACKOMUsers/internal/frame"

	"MACKOMUsers/internal/handler"
)

func main() {
	cfg, err := config.CreateConfig()
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to create config"))
	}

	app := frame.New(cfg)
	ctx := app.GetShutdownContext()

	err = handler.InitRouter(ctx, app, cfg)
	if err != nil {
		log.Fatal(errors.Wrap(err, "cant init router"))
	}

	app.Run()
}
