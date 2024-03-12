package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/ThisIsDaWei/greenlight/internal/data"
)

// Handler for "POST /v1/movies".
func (self *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new movie")
}

// Handler for "GET /v1/movies/:id".
func (self *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	id, err := self.readIDParam(r)
	if err != nil || id < 1 {
		self.notFoundResponse(w, r)
		return
	}

	movie := data.Movie{
		ID:        id,
		CreatedAt: time.Now(),
		Title:     "Casablanca",
		Runtime:   102,
		Genres:    []string{"drama", "romance", "war"},
		Version:   1,
	}

	err = self.writeJSON(w, http.StatusOK, envelope{"movie": movie}, nil)
	if err != nil {
		self.serverErrorResponse(w, r, err)
	}
}
