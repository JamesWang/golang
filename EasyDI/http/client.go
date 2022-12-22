package http

import (
	"fmt"
	"io"
	"net/http"
)

func RequestData(url string, converter func(string) []string) []string {
	response, err := http.Get(url)
	if err == nil && response.StatusCode == http.StatusOK {
		data, err := io.ReadAll(response.Body)
		if err == nil {
			defer response.Body.Close()
			return converter(string(data[:]))
		}
	} else {
		fmt.Printf("Error: %v, Status Code:%v\n", err.Error(), response.StatusCode)
	}
	return []string{}
}
