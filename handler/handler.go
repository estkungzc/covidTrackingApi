package handler

import (
	"github.com/estkungzc/covidTrackingApi/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	CovidTrackerServiceImp *service.CovidTrackerServiceImp
}

func NewHandler(CovidTrackerServiceImp *service.CovidTrackerServiceImp) *Handler {
	return &Handler{CovidTrackerServiceImp: CovidTrackerServiceImp}
}

func (h *Handler) CovidSummary(c *gin.Context) {

	covidStat, _ := h.CovidTrackerServiceImp.GetCovidStat()

	response := h.CovidTrackerServiceImp.GetCovidSummary(covidStat)

	c.JSON(http.StatusOK, response)
}
