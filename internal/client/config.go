package client

import (
	"bufio"
	"io/ioutil"
	"os"

	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
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
	if !fileExist(configFileName) {
		log.Fatal("No config file")
	} else {
		readConfig()
	}
}

func handleErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func initUUID() {
	conf.UUID = generateUUID()
	log.Info("UUID not exist, generate uuid: ", conf.UUID)

	confFile, err := os.OpenFile(configFileName, os.O_WRONLY, 0666)
	handleErr(err)
	defer confFile.Close()

	d, err := yaml.Marshal(&conf)
	handleErr(err)
	writer := bufio.NewWriter(confFile)
	writer.WriteString(string(d))
	writer.Flush()
}

func readConfig() {
	buf, err := ioutil.ReadFile(configFileName)
	handleErr(err)
	err = yaml.Unmarshal(buf, &conf)
	handleErr(err)
	_, err = uuid.FromString(conf.UUID)
	if err != nil {
		initUUID()
	}
}

func fileExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
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
