package middlew

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/emicklei/go-restful"
	"github.com/gin-gonic/gin"
)

func PingTime(req *restful.Request, resp *restful.Response) {
	io.WriteString(resp, fmt.Sprintf("%s", time.Now()))
}

func EmickleiRestfulService() {
	webservice := new(restful.WebService)
	webservice.Route(webservice.GET("/ping").To(PingTime))
	restful.Add(webservice)
	http.ListenAndServe(":8000", nil)
}

func GinJson() {
	r := gin.Default()
	r.GET("pingTime", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"serverTime": time.Now().UTC(),
		})
	})
	r.Run(":8000")
}
