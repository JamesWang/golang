package main

import (
	. "EasyDI/domain"

	. "EasyDI/log"
	. "EasyDI/write"

	//. "EasyDI/pdf"
	//. "EasyDI/reg"
	//. "EasyDI/read"
	//. "EasyDI/gor"
	//. "EasyDI/http"
	//. "EasyDI/rw"
	//. "EasyDI/algo"
	//. "EasyDI/read"
	. "EasyDI/ds"
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
	//DoMutex()
	//DoRWMutex()
	//arr := []int{5, 2, 7, 3, 4, 8, 6}
	//Insertion_Sort(arr)
	//Bubble_Sort(arr)
	//fmt.Println(arr)
	//fmt.Printf("config=%v\n", Config)
	//WriteData()
	WriteJson()
}
