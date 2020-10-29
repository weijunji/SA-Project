package main

import (
	"time"

	"github.com/henrylee2cn/erpc/v6"
	"github.com/henrylee2cn/erpc/v6/plugin/heartbeat"
	"github.com/weijunji/SA-Project/internal/server"
	"github.com/weijunji/SA-Project/pkg/db"
)

func main() {
	defer erpc.FlushLogger()
	erpc.Infof("Starting server...")
	db.GetDB()
	port := server.GetPort()
	srv := erpc.NewPeer(
		erpc.PeerConfig{ListenPort: port},
		server.NewAuthPlugin(),
		heartbeat.NewPong(),
		server.NewDisconnectPlugin(),
	)
	srv.RouteCall(new(Home))
	srv.RoutePush(new(server.Upload))

	go func() {
		time.Sleep(time.Second * 20)
		if sess, b := srv.GetSession("748fb10f-0e83-11eb-89c8-be8385ee77ba"); b {
			sess.Push("/control/lock", "Please lock the door")
			sess.Push("/control/unlock", "Please lock the door")
			sess.Push("/control/stop", "Please stop the lock")
		} else {
			erpc.Errorf("No such client")
		}
	}()
	srv.ListenAndServe()
}

type Home struct {
	erpc.CallCtx
}

func (h *Home) Test(arg *map[string]string) (map[string]interface{}, *erpc.Status) {
	return map[string]interface{}{
		"arg": *arg,
	}, nil
}
