package handlers

import (
	"context"
	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"

	"github.com/Omar-Arabi1/GoShorten/database"
)

type App struct {
	Queries *database.Queries
}

var ctx = context.Background()

func setupDbConnection(dbName string, schema string) (*App, error) {
	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		return &App{}, err
	}

	// schema variable here comes from the main.go file
	if _, err = db.ExecContext(ctx, schema); err != nil {
		return &App{}, err
	}

	queries := database.New(db)
	return &App{Queries: queries}, nil
}

func New(dbName string, schema string) (*gin.Engine, error) {
	app, err := setupDbConnection(dbName, schema)
	if err != nil {
		return &gin.Engine{}, err
	}

	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	// this first route is used to display the page where the user could shorten their url and a link
	// to view all their urls pointing to route 3 here view_urls
	router.GET("/", app.Index)
	router.POST("/shorten", app.Shorten)
	router.GET("/view_urls", app.ViewUrls)
	router.DELETE("/delete_url/:id", app.Delete)
	router.PUT("/edit_url/:id", app.Edit)

	// this endpoint will be the endpoint that will redirect the user when they use the generated shortened url
	router.GET("/short/:shortUrlKey", app.Short)

	return router, nil
}
