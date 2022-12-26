package route

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os/exec"

	"github.com/julienschmidt/httprouter"
)

func NewSrvMux() *http.ServeMux {
	newMux := http.NewServeMux()

	newMux.HandleFunc("/randomFloat", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, rand.Float64())
	})

	newMux.HandleFunc("/randomInt", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, rand.Intn(100))
	})
	return newMux
}

func HttpRouterCmd(router *httprouter.Router) {
	router.GET("/api/v1/go-version", goVersion)
	router.GET("/api/v1/show-file/:name", getFileContent)
	return
}

func getCommandOutput(command string, arguments ...string) string {
	fmt.Printf("command: %v, arguments: %v\n", command, arguments)
	out, _ := exec.Command(command, arguments...).Output()
	fmt.Println(string(out))
	return string(out)
}

func goVersion(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	response := getCommandOutput("go", "version")
	io.WriteString(w, response)
}

func getFileContent(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	fmt.Fprintf(w, getCommandOutput("cat", params.ByName("name")))
}

func StaticServ(router *httprouter.Router) {
	router.ServeFiles("/static/*filepath", http.Dir("./static"))
}
