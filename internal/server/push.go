package server

import (
	"github.com/henrylee2cn/erpc/v6"
	"github.com/weijunji/SA-Project/pkg/db/clientInfo"
)

type Upload struct {
	erpc.PushCtx
}

func (u *Upload) Status(locked *bool) *erpc.Status {
	if *locked {
		erpc.Infof("Client is locked")
		if val, ok := u.Swap().Load("info"); ok {
			val.(*clientInfo.ClientInfo).LockOp()
		}
	} else {
		erpc.Infof("Client is unlock")
		if val, ok := u.Swap().Load("info"); ok {
			val.(*clientInfo.ClientInfo).UnlockOp()
		}
	}
	return nil
}
