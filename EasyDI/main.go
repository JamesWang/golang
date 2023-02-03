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
	"fmt"
	"net/http"
	"os"
	"time"
)

func main1() {
	l := LoggerAdapter(LogOutput)
	ds := NewSimpleDataStore()
	logic := NewSimpleLogic(l, ds)
	controller := NewController(l, logic)
	http.HandleFunc("/hello", controller.SayHello)
	http.ListenAndServe(":8088", nil)
}

func main2() {
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

type DateFormat string

const (
	DD_MMM_YYYY       = DateFormat("02 January 2006")
	DD_MMM_YYYY_HH_MM = DateFormat("02 January 2006 15:30")
	DD_MM_YYYY_HH_MM  = DateFormat("02-01-2006 15:04")
	HH_MM             = DateFormat("15:04")
)

func show_datetime(format DateFormat, dateString string) {
	d, err := time.Parse(string(format), dateString)
	output := make([]string, 2, 3)
	if err == nil {
		output[0] = fmt.Sprintf("Full: %v", d)
		switch format {
		case DD_MMM_YYYY:
		case DD_MMM_YYYY_HH_MM:
			fallthrough
		case DD_MM_YYYY_HH_MM:
			output[1] = fmt.Sprintf("Date: %d %d %d", d.Day(), d.Month(), d.Year())
			output[2] = fmt.Sprintf("Time: %d %d", d.Hour(), d.Minute())
		case HH_MM:
			output[1] = fmt.Sprintf("Time: %d %d", d.Hour(), d.Minute())
		}
	}
}
func main() {
	start := time.Now()
	if len(os.Args) != 2 {
		fmt.Println("Usage: dates parse_string")
		return
	}
	dateString := os.Args[1]

	show_datetime(DD_MMM_YYYY, dateString)
	show_datetime(DD_MMM_YYYY_HH_MM, dateString)
	show_datetime(DD_MM_YYYY_HH_MM, dateString)
	show_datetime(HH_MM, dateString)

	t := time.Now().Unix()
	fmt.Println("Epoch time:", t)
	d := time.Unix(t, 0)
	fmt.Println("Date:", d.Day(), d.Month(), d.Year())
	fmt.Printf("Time: %d:%d\n", d.Hour(), d.Minute())

	duration := time.Since(start)
	fmt.Println("Execution time:", duration)
}
