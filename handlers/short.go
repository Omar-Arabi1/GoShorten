package handlers

import (
	"database/sql"
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a *App) Short(c *gin.Context) {
	shortUrlKey := c.Param("shortUrlKey")
	url, err := a.Queries.QueryUrlByShortUrl(ctx, sql.NullString{String: shortUrlKey, Valid: true})

	// here we check if the ID is empty it is impossible to be empty if it actually exists so if it is that means
	// that the shorturl key we recieved doens't exist in the database thus we raise an error (the 404 template)
	if url.ID == "" {
		c.HTML(http.StatusNotFound, "not_found.html", errors.New("could not find a url with the given short url key"))
		return
	}

	if err != nil {
		log.Fatal(err)
	}

	c.Redirect(http.StatusMovedPermanently, url.LongUrl.String)
}
