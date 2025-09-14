package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a *App) ViewUrls(c *gin.Context) {
	// ctx variable here comes from the app.go file
	urls, err := a.Queries.GetUrls(ctx)
	if err != nil {
		log.Fatal(err)
	}

	c.HTML(http.StatusOK, "view_urls.html", urls)
}
