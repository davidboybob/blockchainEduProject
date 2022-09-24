package utils

import "log"

var logFn = log.Panic

func HandleErr(err error) {
	if err != nil {
		// log.Panic(err)
		logFn(err)
	}
}

// dictionary 생성자
type dict struct {
	data map[string]string
}

func newDict() *dict {
	d := dict{}
	d.data = map[string]string{}
	return &d
}
