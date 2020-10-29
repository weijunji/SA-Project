package client

import (
	"os"

	"github.com/henrylee2cn/erpc/v6"
)

type Control struct {
	erpc.PushCtx
}

func (c *Control) Lock(arg *string) *erpc.Status {
	erpc.Infof("%s", *arg)
	LockTheDoor()
	return nil
}

func (c *Control) Unlock(arg *string) *erpc.Status {
	erpc.Infof("%s", *arg)
	UnlockTheDoor()
	return nil
}

func (c *Control) Stop(arg *string) *erpc.Status {
	erpc.Infof("%s", *arg)
	erpc.Infof("Stop the client because of the server command.")
	erpc.FlushLogger()
	os.Exit(1)
	return nil
}
