package server

import (
	"github.com/henrylee2cn/erpc/v6"
	"github.com/weijunji/SA-Project/pkg/db/clientInfo"
)

type WebCall struct {
	erpc.CallCtx
}

func (c *WebCall) Locked(locked *bool) (string, *erpc.Status) {
	if sess, b := c.Peer().GetSession(string(c.PeekMeta("uuid"))); b {
		if val, ok := sess.Swap().Load("info"); ok {
			erpc.Debugf("%+v", val.(*clientInfo.ClientInfo))
			erpc.Debugf("%+v", *locked)
			if *locked {
				if !val.(*clientInfo.ClientInfo).Locked {
					sess.Push("/control/lock", "Please lock the door")
					val.(*clientInfo.ClientInfo).LockOp()
				}
			} else {
				if val.(*clientInfo.ClientInfo).Locked {
					sess.Push("/control/unlock", "Please unlock the door")
					val.(*clientInfo.ClientInfo).UnlockOp()
				}
			}
		}
	} else {
		return "No such client", nil
	}
	return "success", nil
}

func (c *WebCall) Offline(uuid *string) (string, *erpc.Status) {
	if sess, b := c.Peer().GetSession(*uuid); b {
		sess.Push("/control/stop", "Please stop the lock")
		if val, ok := sess.Swap().Load("info"); ok {
			val.(*clientInfo.ClientInfo).OfflineOp()
		}
	}
	return "success", nil
}
