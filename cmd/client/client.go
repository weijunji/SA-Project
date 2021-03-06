package main

import (
	"github.com/henrylee2cn/erpc/v6"
	"github.com/henrylee2cn/erpc/v6/plugin/heartbeat"
	log "github.com/sirupsen/logrus"
	"github.com/weijunji/SA-Project/internal/client"
)

func main() {
	defer erpc.FlushLogger()
	log.Infof("Connecting to server %s ...", client.GetHost())
	cli := erpc.NewPeer(
		erpc.PeerConfig{PrintDetail: true},
		client.NewAuthPlugin(),
		heartbeat.NewPing(client.GetHeartbeat(), true),
	)
	cli.RoutePush(new(client.Control))
	sess, stat := cli.Dial(client.GetHost())
	if !stat.OK() {
		log.Fatal(stat)
	}
	log.Info("Connected")

	// push locked status to server
	sess.Push("/upload/status", true)

	// hang up
	c := make(chan int)
	<-c
}
