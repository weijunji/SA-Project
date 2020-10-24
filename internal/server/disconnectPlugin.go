package server

import (
	"github.com/henrylee2cn/erpc/v6"
	log "github.com/sirupsen/logrus"
)

// NewDisconnectPlugin Returns a disconnectHandle plugin.
func NewDisconnectPlugin() *disconnectHandle {
	return &disconnectHandle{}
}

type disconnectHandle struct{}

var (
	_ erpc.PostDisconnectPlugin = new(disconnectHandle)
)

func (h *disconnectHandle) Name() string {
	return "disconnectHandle"
}

func (h *disconnectHandle) PostDisconnect(sess erpc.BaseSession) *erpc.Status {
	log.Infof("[%s] disconnect", sess.ID())
	return nil
}
