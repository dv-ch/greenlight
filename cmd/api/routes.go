package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (self *application) routes() *httprouter.Router {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", self.healthcheckHandler)
	router.HandlerFunc(http.MethodPost, "/v1/movies", self.createMovieHandler)
	router.HandlerFunc(http.MethodGet, "/v1/movies/:id", self.showMovieHandler)

	return router
}
