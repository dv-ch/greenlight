package main

import (
	"net/http"
)

// Respond with app status, operating env and version in JSON.
func (self *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	env := envelope{
		"status": "available",
		"system_info": map[string]string{
			"environment": self.config.env,
			"version":     version,
		},
	}

	err := self.writeJSON(w, http.StatusOK, env, nil)
	if err != nil {
		self.serverErrorResponse(w, r, err)
	}
}
