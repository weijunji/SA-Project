package main

import (
	"github.com/henrylee2cn/erpc/v6"
	"github.com/henrylee2cn/erpc/v6/plugin/auth"
	log "github.com/sirupsen/logrus"
	"github.com/weijunji/SA-Project/internal/server"
)

func main() {
	log.Info("Starting server...")
	port := server.GetPort()
	srv := erpc.NewPeer(
		erpc.PeerConfig{ListenPort: port},
		authChecker,
	)
	srv.RouteCall(new(Home))
	srv.ListenAndServe()
}

var authChecker = auth.NewCheckerPlugin(
	func(sess auth.Session, fn auth.RecvOnce) (ret interface{}, stat *erpc.Status) {
		var authInfo string
		stat = fn(&authInfo)
		if !stat.OK() {
			return
		}
		log.Infof("auth info: %v", authInfo)
		if server.GetAuth() != authInfo {
			return nil, erpc.NewStatus(403, "auth fail", "auth fail detail")
		}
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
