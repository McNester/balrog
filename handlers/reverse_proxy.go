package handlers

import (
	"barlog/utils"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

func NewProxy(service string) gin.HandlerFunc {
	return func(c *gin.Context) {

		remote, err := url.Parse(utils.GetEnvVariable(service))

		if err != nil {
			panic("Error during discovery of orders service " + err.Error())
		}

		proxy := httputil.NewSingleHostReverseProxy(remote)
		proxy.Director = func(req *http.Request) {
			req.Header = c.Request.Header
			req.Host = remote.Host
			req.URL.Scheme = remote.Scheme
			req.URL.Host = remote.Host
			req.URL.Path = c.Request.URL.Path
		}

		proxy.ServeHTTP(c.Writer, c.Request)

	}
}
