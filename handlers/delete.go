package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a *App) Delete(c *gin.Context) {
	id := c.Param("id")
	err := a.Queries.DeleteUrl(ctx, id)
	if err != nil {
		log.Fatal(err)
	}

	c.Redirect(http.StatusSeeOther, "/view_urls")
}
