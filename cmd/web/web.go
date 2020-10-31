package main

import (
	"net/http"

	restful "github.com/emicklei/go-restful/v3"
	"github.com/henrylee2cn/erpc/v6"
	"github.com/weijunji/SA-Project/internal/web"
)

func main() {
	ws := new(restful.WebService).Path("/test")
	restful.Add(ws)
	restful.Add(web.NewClientsService())
	cors := restful.CrossOriginResourceSharing{
		ExposeHeaders:  []string{"X-Custom-Header"},
		AllowedHeaders: []string{"X-Custom-Header", "X-Additional-Header"},
		CookiesAllowed: true,
		Container:      restful.DefaultContainer}
	restful.Filter(cors.Filter)
	restful.Filter(web.BasicAuthenticate)
	erpc.Infof("Start server on localhost:80")
	http.ListenAndServe(":80", nil)
}
