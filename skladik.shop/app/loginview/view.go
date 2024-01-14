package sklad

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginView(c *gin.Context) {
	c.HTML(http.StatusOK,"login.html",nil)
}