package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewServer(port int, h *gin.Engine) *http.Server {
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: h,
	}

	return srv
}
