package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	//. "requestRouting/uuid"
	//. "requestRouting/middlew"
	. "requestRouting/route"

	//. "requestRouting/rpc"
	//. "requestRouting/tests"
	. "requestRouting/app"

	"github.com/julienschmidt/httprouter"
)

func httpRouter() {
	router := httprouter.New()
	HttpRouterCmd(router)
	StaticServ(router)
	http.ListenAndServe(":8080", router)
}
func MainFunc() {
	var cfg Config
	flag.IntVar(&cfg.Port, "port", 4000, "API server port")
	flag.StringVar(&cfg.Env, "env", "development", "Environment (development|staging|production")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	app := &Application{
		Config: cfg,
		Logger: logger,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/v1/healthcheck", app.HealthcheckHandler)
	mux.HandleFunc("/v1/movies", app.ShowMovieHandler)

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Port),
		Handler:      mux,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Printf("starting %s server on %s", cfg.Env, srv.Addr)
	err := srv.ListenAndServe()
	logger.Fatal(err)
}

func main() {
	//mux := &UUID{}
	//mux := NewSrvMux()
	//http.ListenAndServe(":8080", mux)
	//GorillaSrv()
	//MMain()
	//TimeServerActivate()
	//RunSortAndTotal()
	//EmickleiRestfulService()
	//GinJson()
	MainFunc()
}
