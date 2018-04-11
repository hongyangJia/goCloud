package common

import (
	"strings"
	"errors"
)

const (
	PING    = "ping"
	PONG    = "pong"
	CONTENT = "ping is null"
	NULL    = ""
	DEFULT  = -1
)

func ReplacePong(s string) (n string, e error) {
	if strings.Contains(s, PING) {
		s := strings.Replace(s, PING, PONG, DEFULT)
		return s, nil
	}
	return NULL, errors.New(CONTENT)
}
