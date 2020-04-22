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