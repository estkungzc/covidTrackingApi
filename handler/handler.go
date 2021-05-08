package handler

import (
	"github.com/estkungzc/covidTrackingApi/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	covidTrackerService *service.CovidTrackerService
}

func NewHandler(covidTrackerService *service.CovidTrackerService) *Handler {
	return &Handler{covidTrackerService: covidTrackerService}
}

func (h *Handler) CovidSummary(c *gin.Context) {
	response := h.covidTrackerService.GetCovidSummary()

	c.JSON(http.StatusOK, response)
}
