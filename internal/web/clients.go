package web

import (
	"errors"
	"net/http"

	"github.com/emicklei/go-restful/v3"
	"github.com/weijunji/SA-Project/pkg/db/clientInfo"
)

func NewClientsService() *restful.WebService {
	ws := new(restful.WebService).Path("/client")
	ws.Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON)

	ws.Route(ws.GET("/").To(getClientList).
		Doc("get client list").
		Writes([]clientInfo.ClientInfo{}))

	ws.Route(ws.PUT("/{uuid}").To(putClient).
		Doc("change client status"))
	return ws
}

func getClientList(request *restful.Request, response *restful.Response) {
	clientInfos := clientInfo.GetClientInfos()
	response.WriteAsJson(clientInfos)
}

func putClient(request *restful.Request, response *restful.Response) {
	uuid := request.PathParameter("uuid")
	c := clientInfo.GetClient(uuid)
	if c.UUID == "" {
		response.WriteError(http.StatusNotFound, errors.New("Client not found"))
	} else {
		// TODO
		response.WriteAsJson(c)
	}
}
