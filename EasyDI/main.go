package main

import (
	. "EasyDI/domain"
	. "EasyDI/ds"
	. "EasyDI/log"

	//. "EasyDI/pdf"
	//. "EasyDI/reg"
	//. "EasyDI/read"
	. "EasyDI/gor"
	//. "EasyDI/http"
	"net/http"
)

func main1() {
	l := LoggerAdapter(LogOutput)
	ds := NewSimpleDataStore()
	logic := NewSimpleLogic(l, ds)
	controller := NewController(l, logic)
	http.HandleFunc("/hello", controller.SayHello)
	http.ListenAndServe(":8088", nil)
}

func main() {
	//Extract()
	//Reg()
	//MySplit()
	//Display()
	//ProcessData()
	//products := RequestData("http://localhost:3500/api/products")
	//converter := func(s string) []string { return strings.Split(s, "\n") }
	//music := RequestData("http://localhost:8088/admin/list", converter)
	//for _, m := range music {
	//		fmt.Println(m)
	//}
	//DoWait()
	DoMutex()
}
