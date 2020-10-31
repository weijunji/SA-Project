package web

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/weijunji/SA-Project/pkg/db/clientInfo"
)

func NewClientsService() *restful.WebService {
	ws := new(restful.WebService).Path("/clients")
	ws.Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON)

	ws.Route(ws.GET("/").To(GetClientList).
		Doc("get client list").
		Writes([]clientInfo.ClientInfo{}))
	return ws
}

func GetClientList(request *restful.Request, response *restful.Response) {

}
