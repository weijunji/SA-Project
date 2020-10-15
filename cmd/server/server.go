package main

import (
	"strings"

	"github.com/henrylee2cn/erpc/v6"
	"github.com/henrylee2cn/erpc/v6/plugin/auth"
	"github.com/henrylee2cn/erpc/v6/plugin/heartbeat"
	log "github.com/sirupsen/logrus"
	"github.com/weijunji/SA-Project/internal/server"
)

func main() {
	log.Info("Starting server...")
	port := server.GetPort()
	srv := erpc.NewPeer(
		erpc.PeerConfig{ListenPort: port},
		authChecker,
		heartbeat.NewPong(),
	)
	srv.RouteCall(new(Home))
	srv.ListenAndServe()
	// _, b := srv.GetSession("748fb10f-0e83-11eb-89c8-be8385ee77ba")
}

var authChecker = auth.NewCheckerPlugin(
	func(sess auth.Session, fn auth.RecvOnce) (ret interface{}, stat *erpc.Status) {
		var authInfo string
		stat = fn(&authInfo)
		if !stat.OK() {
			return
		}
		auths := strings.Split(authInfo, "%")
		if len(auths) == 1 {
			return nil, erpc.NewStatus(403, "auth fail", "no client uuid")
		}
		if server.GetAuth() != auths[0] {
			return nil, erpc.NewStatus(403, "auth fail", "wrong auth code")
		}
		sess.SetID(auths[1])
		return "pass", nil
	},
	erpc.WithBodyCodec('s'),
)

type Home struct {
	erpc.CallCtx
}

func (h *Home) Test(arg *map[string]string) (map[string]interface{}, *erpc.Status) {
	return map[string]interface{}{
		"arg": *arg,
	}, nil
}
