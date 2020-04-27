package cha

import (
	"log"
)

var WarningLog = func(msg string) {
	log.Print("go-chatty: " + msg)
}
