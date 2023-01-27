package app

import (
	"encoding/json"
	"net/http"
	"requestRouting/data"
)

func (app *Application) HealthcheckHandler(w http.ResponseWriter, r *http.Request) {
	//js := `{"status": "available", "environment": %q, "version": %q}`
	//js = fmt.Sprintf(js, app.config.env, version)
	data := map[string]string{
		"status":      "available",
		"environment": app.Config.Env,
		"version":     Version,
	}

	js, err := json.Marshal(data)
	if err != nil {
		app.Logger.Print(err)
		http.Error(w, "The server encountered a problem and could not process your requests", http.StatusInternalServerError)
		return
	}
	js = append(js, '\n')
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func (app *Application) WriteJson(w http.ResponseWriter, status int, data any, headers http.Header) error {
	js, err := json.Marshal(data)
	if err != nil {
		return err
	}
	js = append(js, '\n')
	for key, value := range headers {
		w.Header()[key] = value
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)
	return nil
}

func (app *Application) ShowMovieHandler(w http.ResponseWriter, r *http.Request) {
	movie := data.SampleData(1)
	err := app.WriteJson(w, http.StatusOK, movie, nil)
	if err != nil {
		app.Logger.Print(err)
		http.Error(w, "The server encountered a problem and could not process your requests", http.StatusInternalServerError)
	}
}
