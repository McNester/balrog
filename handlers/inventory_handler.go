package handlers

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)


type InventoryHandler struct{
    
}

func NewInvenotryHandler() *InventoryHandler{
    return &InventoryHandler{}
}


func (h *InventoryHandler) Proxy(c *gin.Context){
    remote, err := url.Parse("http://localhost:8082")

    if err != nil {
        panic("Error during discovery of products service " + err.Error())
    }

    proxy := httputil.NewSingleHostReverseProxy(remote)
    proxy.Director = func(r *http.Request) {
        r.Header = c.Request.Header
        r.Host = remote.Host
        r.URL.Scheme = remote.Scheme
        r.URL.Host = remote.Host
        r.URL.Path = c.Param("products")
    }

    proxy.ServeHTTP(c.Writer, c.Request)

}
