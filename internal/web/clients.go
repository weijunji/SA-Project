package web

import (
	"errors"
	"net/http"

	"github.com/emicklei/go-restful/v3"
	"github.com/henrylee2cn/erpc/v6"
	"github.com/weijunji/SA-Project/pkg/db/clientInfo"
)

var Sess erpc.Session

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

type ClientPutParams struct {
	Name   string
	Locked *bool
	Online *bool
}

func putClient(request *restful.Request, response *restful.Response) {
	uuid := request.PathParameter("uuid")
	c := clientInfo.GetClient(uuid)
	if c.UUID == "" {
		response.WriteError(http.StatusNotFound, errors.New("Client not found"))
	} else {
		params := new(ClientPutParams)
		err := request.ReadEntity(&params)
		if err != nil {
			response.WriteError(http.StatusInternalServerError, err)
		} else {
			if params.Name != "" {
				c.SetName(params.Name)
			}
			if params.Locked != nil && *params.Locked != c.Locked {
				var result string
				stat := Sess.Call("/web_call/locked",
					*params.Locked,
					&result,
					erpc.WithAddMeta("uuid", uuid)).Status()
				erpc.Debugf("%+v", stat)
				if stat.OK() {
					erpc.Debugf("%+v", result)
				}
			}
			if params.Online != nil && c.Online && !*params.Online {
				var result string
				stat := Sess.Call("/web_call/offline", uuid, &result).Status()
				if stat.OK() {
					erpc.Debugf("%+v", result)
				}
			}
		}
		response.WriteAsJson("Success")
	}
}
