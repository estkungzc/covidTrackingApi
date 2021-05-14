package app

import (
	"context"
	"github.com/estkungzc/covidTrackingApi/handler"
	"github.com/estkungzc/covidTrackingApi/internal/config"
	"github.com/estkungzc/covidTrackingApi/internal/server"
	"github.com/estkungzc/covidTrackingApi/model"
	"github.com/estkungzc/covidTrackingApi/router"
	"github.com/estkungzc/covidTrackingApi/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type App struct {
	router *http.Server
	cfg    *model.Config
}

func NewApp() *App {
	cfg := model.Config{}

	config.LoadConfig(&cfg, "config")
	gin.SetMode(cfg.Env)

	covidTrackerService := service.NewCovidTrackerServiceImp()

	handlers := handler.NewHandler(covidTrackerService)
	httpRouter := router.NewHTTPRouter(handlers)
	httpServer := server.NewServer(cfg.Http.Port, httpRouter)

	return &App{
		router: httpServer,
		cfg:    &cfg,
	}
}

func (app *App) Start() error {
	log.Printf("Start service: %s (%s) \n", app.cfg.ServiceName, app.cfg.Env)
	return app.router.ListenAndServe()
}

func (app *App) Shutdown(ctx context.Context) error {
	return app.router.Shutdown(ctx)
}
