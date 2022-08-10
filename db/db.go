package db

import (
	"rest-module/models"
)

// set up a database dummy
var (
	Moviedb = make(map[string]models.Movie)
)
