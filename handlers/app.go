package handlers

import (
	"context"
	"database/sql"

	_ "github.com/mattn/go-sqlite3"

	"github.com/Omar-Arabi1/GoShorten/database"
)

type App struct {
	Queries *database.Queries
}

var ctx = context.Background()

func Setup(dbName string) (*App, error) {
	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		return &App{}, err
	}

	// the second string is for query string which just means the query we are doing I think so since I am
	// not doing any queries I am not passing anything in the reference I am using at
	// https://docs.sqlc.dev/en/latest/tutorials/getting-started-sqlite.html they use a variable called dll at the top
	// which I don't understand but from the documentation I see for this ExecContext here it takes in a second query param
	// and even if I am wrong the only functionality of this is to create the table as said in the docs and that is what they did
	// so it is what it is...
	if _, err = db.ExecContext(ctx, ""); err != nil {
		return &App{}, err
	}

	queries := database.New(db)
	return &App{Queries: queries}, nil
}
