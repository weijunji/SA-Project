package main

import (
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
	srv.RoutePush(new(server.Upload))
	srv.RouteCall(new(server.WebCall))

	srv.ListenAndServe()
}
