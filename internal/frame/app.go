package frame

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi/v5"
	"github.com/pkg/errors"

	"MACKOMUsers/internal/config"
)

type App struct {
	config      *config.Config
	shutdownCtx context.Context
	finish      context.CancelFunc

	publicRouter chi.Router
}

func New(cfg *config.Config) *App {
	shutdownCtx, finish := context.WithCancel(context.Background())

	app := &App{
		config:      cfg,
		shutdownCtx: shutdownCtx,
		finish:      finish,
	}

	return app
}

func (x *App) RegisterHttpHandler(method HttpMethod, pattern string, handler HandlerFn) {
	if x.publicRouter == nil {
		x.initHttpRouter()
	}

	innerHandler := func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json; charset=utf-8")

		data, _ := io.ReadAll(request.Body)
		request.Body = io.NopCloser(bytes.NewReader(data))

		response, err := handler(request)

		request.Body = io.NopCloser(bytes.NewReader(data))

		if response != nil {
			for _, addEntry := range response.Headers.GetAddEntrySlice() {
				writer.Header().Add(addEntry.name, addEntry.value)
			}

			for name, value := range response.Headers.GetSetEntryMap() {
				writer.Header().Set(name, value)
			}
		}

		if err != nil {
			var resultCode int

			if response != nil && response.Code != 0 {
				resultCode = response.Code
			} else {
				resultCode = http.StatusInternalServerError
			}

			setErrorResponse(err, resultCode, writer, request)
			return
		}

		if response.Code < 1 {
			writer.WriteHeader(http.StatusOK)
		} else {
			writer.WriteHeader(response.Code)
		}

		_, err = writer.Write(response.Data)
	}

	switch method {
	case Get:
		x.publicRouter.Get(pattern, innerHandler)
	case Post:
		x.publicRouter.Post(pattern, innerHandler)
	case Head:
		x.publicRouter.Head(pattern, innerHandler)
	case Put:
		x.publicRouter.Put(pattern, innerHandler)
	case Patch:
		x.publicRouter.Patch(pattern, innerHandler)
	case Delete:
		x.publicRouter.Delete(pattern, innerHandler)
	case Options:
		x.publicRouter.Options(pattern, innerHandler)
	default:
	}
}

func (x *App) Run() {
	notifyCtx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	x.runPublicHttpServer()

	<-notifyCtx.Done()
}

func (x *App) runPublicHttpServer() {
	if x.publicRouter == nil {
		x.initHttpRouter()
	}

	httpServer := &http.Server{
		Addr:    fmt.Sprintf(":%d", x.publicHttpPort()),
		Handler: x.publicRouter,
	}

	fmt.Printf("http: starting public server on port %d", x.publicHttpPort())

	go func() {
		serverErr := httpServer.ListenAndServe()
		if serverErr != nil && errors.Is(serverErr, http.ErrServerClosed) {
			x.finish()
			return
		}
	}()
}

func (x *App) initHttpRouter() {
	router := chi.NewRouter()

	x.publicRouter = router
}

func (x *App) publicHttpPort() int {
	if x.config.HTTPServer.Port == 0 {
		return config.HttpPort
	}

	return x.config.HTTPServer.Port
}

func (x *App) GetShutdownContext() context.Context {
	return x.shutdownCtx
}
