package server

import (
	"github.com/henrylee2cn/erpc/v6"
	"github.com/weijunji/SA-Project/pkg/db/clientInfo"
)

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
	erpc.Infof("[%s] disconnect", sess.ID())
	if val, ok := sess.Swap().Load("info"); ok {
		val.(*clientInfo.ClientInfo).OfflineOp()
	} else {
		erpc.Errorf("Load info failed")
	}
	return nil
}
