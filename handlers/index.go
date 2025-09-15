package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a *App) Index(c *gin.Context) {
	shortUrlKey := c.Query("shortUrlKey")
	c.HTML(http.StatusOK, "index.html", shortUrlKey)
}
