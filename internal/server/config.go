package server

import (
	"io/ioutil"

	"github.com/henrylee2cn/erpc/v6"
	. "github.com/weijunji/SA-Project/pkg/tools"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Port uint16
	Auth string
}

const configFileName string = "sap_server.yaml"

var conf Config

func init() {
	erpc.Infof("Read config file")
	if !FileExist(configFileName) {
		erpc.Fatalf("Cannot find config file")
	} else {
		readConfig()
	}
}

func readConfig() {
	buf, err := ioutil.ReadFile(configFileName)
	HandleErrPanic(err)
	err = yaml.Unmarshal(buf, &conf)
	HandleErrPanic(err)
}

func GetPort() uint16 {
	return conf.Port
}

func GetAuth() string {
	return conf.Auth
}
