package server

import (
	"context"
	web "net/http"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/gorilla/mux"
	"github.com/pronuu/lincoln/internal/config"
	"github.com/pronuu/lincoln/internal/utils"
)

// NewHTTP ...
func NewHTTP(config *config.Config, logger log.Logger, router *mux.Router) *HTTP {
	server := &web.Server{
		Handler:      utils.SetCORSPolicy(router),
		Addr:         config.Address,
		WriteTimeout: config.WriteTimeout,
		ReadTimeout:  config.ReadTimeout,
		IdleTimeout:  config.IdleTimeout,
	}
	return &HTTP{
		Server: server,
		Router: router,
		Logger: logger,
		Config: config,
	}
}

// HTTP ...
type HTTP struct {
	Server *web.Server
	Logger log.Logger
	Router *mux.Router
	Config *config.Config
}

// Start ...
func (h *HTTP) Start(ctx context.Context) error {
	h.Info().Log("msg", "service starting", "address", h.Config.Address)
	return h.Server.ListenAndServe()
}

// Stop ...
func (h *HTTP) Stop(ctx context.Context) error {
	h.Info().Log("msg", "service stopping")
	return h.Server.Shutdown(ctx)
}

// Info ...
func (h *HTTP) Info(messages ...string) log.Logger {
	return level.Info(h.Logger)
}
