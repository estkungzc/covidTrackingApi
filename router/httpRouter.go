package router

import (
	"github.com/estkungzc/covidTrackingApi/handler"
	"github.com/gin-gonic/gin"
)

func NewHTTPRouter(h *handler.Handler) *gin.Engine {
	r := gin.Default()

	covidGroup := r.Group("covid")
	{
		covidGroup.GET("summary", h.CovidSummary)
	}

	return r
}
