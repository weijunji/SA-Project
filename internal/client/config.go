package client

import (
	"bufio"
	"io/ioutil"
	"os"

	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	. "github.com/weijunji/SA-Project/pkg/tools"
	"gopkg.in/yaml.v2"
)

type Config struct {
	UUID      string
	HeartBeat int
}

const configFileName string = "sap_client.yaml"

var conf Config = Config{"", 2}

func init() {
	log.Info("Read config file")
	if !FileExist(configFileName) {
		log.Fatal("No config file")
	} else {
		readConfig()
	}
}

func initUUID() {
	conf.UUID = generateUUID()
	log.Info("UUID not exist, generate uuid: ", conf.UUID)

	confFile, err := os.OpenFile(configFileName, os.O_WRONLY, 0666)
	HandleErrPanic(err)
	defer confFile.Close()

	d, err := yaml.Marshal(&conf)
	HandleErrPanic(err)
	writer := bufio.NewWriter(confFile)
	writer.WriteString(string(d))
	writer.Flush()
}

func readConfig() {
	buf, err := ioutil.ReadFile(configFileName)
	HandleErrPanic(err)
	err = yaml.Unmarshal(buf, &conf)
	HandleErrPanic(err)
	_, err = uuid.FromString(conf.UUID)
	if err != nil {
		initUUID()
	}
}

func generateUUID() string {
	uuid := uuid.NewV1()
	return string(uuid.String())
}

func GetUUID() string {
	return conf.UUID
}

func GetHeartbeat() int {
	return conf.HeartBeat
}
