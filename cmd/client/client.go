package main

import (
	"github.com/henrylee2cn/erpc/v6"
	"github.com/henrylee2cn/erpc/v6/plugin/auth"
	"github.com/henrylee2cn/erpc/v6/plugin/heartbeat"
	log "github.com/sirupsen/logrus"
	"github.com/weijunji/SA-Project/internal/client"
)

func main() {
	log.Infof("Connecting to server %s ...", client.GetHost())
	cli := erpc.NewPeer(
		erpc.PeerConfig{PrintDetail: true},
		authBearer,
		heartbeat.NewPing(client.GetHeartbeat(), true),
	)
	sess, stat := cli.Dial(client.GetHost())
	if !stat.OK() {
		log.Fatal(stat)
	}
	log.Info("Connected")
	var result interface{}
	stat = sess.Call("/home/test",
		map[string]string{
			"author": "henrylee2cn",
		},
		&result,
	).Status()
	if !stat.OK() {
		log.Error(stat)
	}
	log.Infof("result:%v", result)

	// hang up
	c := make(chan int)
	<-c
}

var authBearer = auth.NewBearerPlugin(
	func(sess auth.Session, fn auth.SendOnce) (stat *erpc.Status) {
		var ret string
		stat = fn(client.GetAuth()+"%"+client.GetUUID(), &ret)
		if !stat.OK() {
			return
		}
		log.Infof("auth info: %s, result: %s", client.GetAuth(), ret)
		return
	},
	erpc.WithBodyCodec('s'),
)
