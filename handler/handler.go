package handler

import (
	"net/http"

	"github.com/estkungzc/covidTrackingApi/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	CovidTrackerService service.CovidTrackerService
}

func NewHandler(CovidTrackerService service.CovidTrackerService) *Handler {
	return &Handler{CovidTrackerService: CovidTrackerService}
}

func (h *Handler) CovidSummary(c *gin.Context) {

	covidStat, _ := h.CovidTrackerService.GetCovidStat()

	response := h.CovidTrackerService.GetCovidSummary(covidStat)

	c.JSON(http.StatusOK, response)
}
