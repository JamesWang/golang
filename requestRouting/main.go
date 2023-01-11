package main

import (
	"net/http"
	//. "requestRouting/uuid"
	. "requestRouting/middlew"
	. "requestRouting/route"

	//. "requestRouting/rpc"
	//. "requestRouting/tests"
	"github.com/julienschmidt/httprouter"
)

func httpRouter() {
	router := httprouter.New()
	HttpRouterCmd(router)
	StaticServ(router)
	http.ListenAndServe(":8080", router)
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
	GinJson()
}
