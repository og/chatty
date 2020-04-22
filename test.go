package cha

import (
	glist "github.com/og/x/list"
	gis "github.com/og/x/test"
	"testing"
)

func testPickString(t *testing.T, seed []string, counter map[string]int) {
	is := gis.New(t)
	is.True(len(counter) != 0)
	for name,_  := range counter {
		is.True(glist.StringList{seed}.In(name))
	}

}
