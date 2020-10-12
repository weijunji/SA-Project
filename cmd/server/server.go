package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/weijunji/SA-Project/internal/server"
)

func main() {
	log.Info("Start server")
	log.Info(server.Add(2, 4))
}
