package uuid

import (
	"crypto/rand"
	"fmt"
	"net/http"

	"go.uber.org/zap"
)

type UUID struct {
}

func (p *UUID) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		giveRandomUUID(w, r)
		return
	}
	http.NotFound(w, r)
}

func giveRandomUUID(w http.ResponseWriter, r *http.Request) {
	c := 10
	logger, _ := zap.NewProduction()

	defer logger.Sync()
	b := make([]byte, c)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	logger.Info("generating UUID")
	fmt.Fprintf(w, fmt.Sprintf("%x", b))
}
