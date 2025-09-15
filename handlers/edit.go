package handlers

import (
	"database/sql"
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/Omar-Arabi1/GoShorten/database"
	"github.com/gin-gonic/gin"
)

func (a *App) Edit(c *gin.Context) {
	id := c.Param("id")
	updatedLongUrl := c.Request.FormValue("updated-long-url")

	if strings.TrimSpace(updatedLongUrl) == "" {
		c.AbortWithError(http.StatusNotAcceptable, errors.New("the url is empty"))
		return
	}

	err := a.Queries.UpdateUrl(ctx, database.UpdateUrlParams{
		LongUrl: sql.NullString{String: updatedLongUrl, Valid: true},
		ID:      id,
	})

	if err != nil {
		log.Fatal(err)
	}

	c.Redirect(http.StatusSeeOther, "/view_urls")
}
