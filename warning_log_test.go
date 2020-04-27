package cha_test

import (
	cha "github.com/og/go-chatty"
	gtest "github.com/og/x/test"
	"testing"
)

func TestWarningLog(t *testing.T) {
	as := gtest.AS(t)
	shouldMsg := ""
	cha.WarningLog("some")
	cha.WarningLog = func(msg string) {
		shouldMsg = msg
	}
	cha.WarningLog("a")
	as.Eql(shouldMsg, "a")
}
