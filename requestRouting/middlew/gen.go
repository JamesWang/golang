package middlew

import (
	"fmt"
	"net/http"
)

func Generator() func() int {
	var i = 0
	return func() int {
		i++
		return i
	}
}

func MiddleWare(origHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Executing middleware before request phase!")
		origHandler.ServeHTTP(w, r)
		fmt.Println("Executing middleware after response phase")
	})
}

func HandleIt(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Executing mainHandler")
	w.Write([]byte("OK"))
}

func MMain() {
	origHandler := http.HandlerFunc(HandleIt)
	http.Handle("/", MiddleWare(origHandler))
	http.ListenAndServe(":8080", nil)
}
