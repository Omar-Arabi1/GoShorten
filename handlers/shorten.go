package handlers

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/Omar-Arabi1/GoShorten/database"
	"github.com/Omar-Arabi1/GoShorten/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (a *App) Shorten(c *gin.Context) {
	longUrl := c.Request.FormValue("long-url")

	if strings.TrimSpace(longUrl) == "" {
		c.AbortWithError(http.StatusNotAcceptable, errors.New("the url is empty"))
		return
	}

	shortUrlKey := utils.ShortenUrl(longUrl)
	err := a.Queries.AddUrl(ctx, database.AddUrlParams{
		ID:              uuid.NewString(),
		LongUrl:         sql.NullString{String: longUrl, Valid: true},
		ShortenedUrlKey: sql.NullString{String: shortUrlKey, Valid: true},
	})

	if err != nil {
		log.Fatal(err)
	}

	c.Redirect(http.StatusSeeOther, fmt.Sprintf("/?shortUrlKey=%s", shortUrlKey))
}
