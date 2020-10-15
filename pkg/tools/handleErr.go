package tools

import (
	log "github.com/sirupsen/logrus"
)

func HandleErrPanic(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func HandleErr(err error) {
	if err != nil {
		log.Error(err)
	}
}
