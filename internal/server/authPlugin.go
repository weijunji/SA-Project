package server

import (
	"strings"

	"github.com/henrylee2cn/erpc/v6"
	"github.com/henrylee2cn/erpc/v6/plugin/auth"
	"github.com/weijunji/SA-Project/pkg/db/clientInfo"
)

func NewAuthPlugin() erpc.Plugin {
	return auth.NewCheckerPlugin(
		func(sess auth.Session, fn auth.RecvOnce) (ret interface{}, stat *erpc.Status) {
			var authInfo string
			stat = fn(&authInfo)
			if !stat.OK() {
				return
			}
			auths := strings.Split(authInfo, "%")
			if len(auths) == 1 {
				return nil, erpc.NewStatus(403, "auth fail", "no client uuid")
			}
			if GetAuth() != auths[0] {
				return nil, erpc.NewStatus(403, "auth fail", "wrong auth code")
			}
			sess.SetID(auths[1])
			c := clientInfo.NewClientInfo(auths[1], true, true)
			sess.Swap().Store("info", c)
			return "pass", nil
		},
		erpc.WithBodyCodec('s'),
	)
}
