package web

import (
	"io/ioutil"

	log "github.com/sirupsen/logrus"
	. "github.com/weijunji/SA-Project/pkg/tools"
	"gopkg.in/yaml.v2"
)

type Config struct {
	UUID      string
	HeartBeat int
	Auth      string
	Host      string
}

const configFileName string = "sap_web.yaml"

var conf Config

func init() {
	log.Info("Read config file")
	if !FileExist(configFileName) {
		log.Fatal("No config file")
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

func GetUUID() string {
	return conf.UUID
}

func GetHeartbeat() int {
	return conf.HeartBeat
}

func GetAuth() string {
	return conf.Auth
}

func GetHost() string {
	return conf.Host
}
