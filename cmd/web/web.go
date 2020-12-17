package main

import (
	"log"
	"net/http"

	restful "github.com/emicklei/go-restful/v3"
	"github.com/henrylee2cn/erpc/v6"
	"github.com/henrylee2cn/erpc/v6/plugin/heartbeat"
	"github.com/weijunji/SA-Project/internal/web"
)

func main() {
	restful.Add(web.NewClientsService())
	cors := restful.CrossOriginResourceSharing{
		ExposeHeaders:  []string{"*"},
		AllowedHeaders: []string{"content-type"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "HEAD"},
		CookiesAllowed: false,
		Container:      restful.DefaultContainer}
	restful.Filter(cors.Filter)
	go func() {
		erpc.Infof("Start server on localhost:80")
		http.ListenAndServe(":80", nil)
	}()

	peer := erpc.NewPeer(
		erpc.PeerConfig{PrintDetail: true},
		web.NewAuthPlugin(),
		heartbeat.NewPing(web.GetHeartbeat(), true),
	)
	sess, stat := peer.Dial(web.GetHost())
	web.Sess = sess
	if !stat.OK() {
		log.Fatal(stat)
	}
	// hang up
	c := make(chan int)
	<-c
}
