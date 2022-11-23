package main

import (
	"net/http"

	"github.com/karokojnr/logger-service/cmd/data"
)

type JSONPayload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

func (app *Config) WriteLog(w http.ResponseWriter, r *http.Request) {
	// read json into var
	var requestPayload JSONPayload
	_ = app.readJSON(w, r, &requestPayload)

	// inserData

	event := data.LogEntry{
		Name: requestPayload.Name,
		Data: requestPayload.Data,
	}

	err := app.Models.LogEntry.Insert(event)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	resp := jsonResponse{
		Error:   false,
		Message: "logged",
	}

	app.wrtieJSON(w, http.StatusAccepted, resp)
}
