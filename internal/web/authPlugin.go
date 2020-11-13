package web

import (
	"github.com/henrylee2cn/erpc/v6"
	"github.com/henrylee2cn/erpc/v6/plugin/auth"
)

func NewAuthPlugin() erpc.Plugin {
	return auth.NewBearerPlugin(
		func(sess auth.Session, fn auth.SendOnce) (stat *erpc.Status) {
			var ret string
			stat = fn(GetAuth()+"%"+GetUUID(), &ret)
			if !stat.OK() {
				return
			}
			erpc.Infof("auth info: %s, result: %s", GetAuth(), ret)
			return
		},
		erpc.WithBodyCodec('s'),
	)
}
