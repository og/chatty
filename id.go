package cha

import (
	"github.com/google/uuid"
	gconv "github.com/og/x/conv"
	"sync"
)

func UUID() string {
	return uuid.New().String()
}

type incrementID struct {
	Value int
	lock *sync.RWMutex
}
func (incrID *incrementID) Int() int {
	value := incrID.Value
	incrID.lock.Lock()
	incrID.Value +=1
	incrID.lock.Unlock()
	return value
}
func (incrID *incrementID) String() string {
	return gconv.IntString(incrID.Int())
}
func IncrID() incrementID {
	return incrementID{
		Value: 1,
		lock: new(sync.RWMutex),
	}
}
var nameIncrIDMap = struct {
	hash map[string/* mame */]*incrementID
	lock *sync.RWMutex
}{

}
func init () {
	nameIncrIDMap.hash = map[string]*incrementID{}
	nameIncrIDMap.lock = new(sync.RWMutex)
}
func NameIncrID(name string) string {
	nameIncrIDMap.lock.Lock()
	defer func() {
		nameIncrIDMap.lock.Unlock()
	}()
	incrID, has := nameIncrIDMap.hash[name]
	if !has {
		incrIDElem := IncrID()
		incrID = &incrIDElem
		nameIncrIDMap.hash[name] = incrID
	}
	return incrID.String()
}