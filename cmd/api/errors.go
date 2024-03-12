package main

import (
	"fmt"
	"net/http"
)

// Helper to log err msg.
func (self *application) logError(r *http.Request, err error) {
	self.logger.Println(err)
}

// Helper to send JSON-formatted err msg to client with status code.
func (self *application) errorResponse(w http.ResponseWriter, r *http.Request, status int, message any) {
	env := envelope{"error": message}

	err := self.writeJSON(w, status, env, nil)
	if err != nil {
		self.logError(r, err)
		w.WriteHeader(500)
	}
}

// Helper to log err msg when app encounters problem at runtime, and calls errorResponse().
func (self *application) serverErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	self.logError(r, err)
	message := "the server encountered a problem and could not process your request"
	self.errorResponse(w, r, http.StatusInternalServerError, message)
}

// Helper to send 404 Not Found status code and JSON response to client.
func (self *application) notFoundResponse(w http.ResponseWriter, r *http.Request) {
	message := "the requested resource could not be found"
	self.errorResponse(w, r, http.StatusNotFound, message)
}

// Helper to send 405 Method Not Allowed status code and JSON response to the client.
func (self *application) methodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("the %s method is not supported for this resource", r.Method)
	self.errorResponse(w, r, http.StatusMethodNotAllowed, message)
}
