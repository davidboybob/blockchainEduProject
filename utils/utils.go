package utils

import "log"

var logFn = log.Panic

func HandleErr(err error) {
	if err != nil {
		// log.Panic(err)
		logFn(err)
	}
}
