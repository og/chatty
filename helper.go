package cha

import (
	ge "github.com/og/x/error"
	"math/big"
)
import "crypto/rand"

func PickString(list []string) string {
	index, err := rand.Int(rand.Reader, big.NewInt(int64(len(list)))) ; ge.Check(err)
	return list[index.Int64()]
}
func Run(n int, fn func(i int) (_break bool) ) {
	for i:=0; i<n; i++ {
		if fn(i) {
			break
		}
	}
}
