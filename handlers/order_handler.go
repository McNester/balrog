package handlers

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
}

func NewOrderHandler() *OrderHandler {
	return &OrderHandler{}
}

func (h *OrderHandler) Proxy(c *gin.Context) {
	remote, err := url.Parse("http://localhost:8083")

	if err != nil {
		panic("Error during discovery of orders service " + err.Error())
	}

	proxy := httputil.NewSingleHostReverseProxy(remote)
	proxy.Director = func(req *http.Request) {
		req.Header = c.Request.Header
		req.Host = remote.Host
		req.URL.Scheme = remote.Scheme
		req.URL.Host = remote.Host
		req.URL.Path = c.Param("orders")
	}

	proxy.ServeHTTP(c.Writer, c.Request)

}
