package cha

import (
	glist "github.com/og/x/list"
	gtest "github.com/og/x/test"
	"testing"
)

func testPickString(t *testing.T, seed []string, counter map[string]int) {
	as := gtest.NewAS(t)
	as.True(len(counter) != 0)
	for name,_  := range counter {
		as.True(glist.StringList{seed}.In(name))
	}
}
