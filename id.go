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
	value int
	lock *sync.RWMutex
}
func (incrID *incrementID) Int() int {
	value := incrID.value
	incrID.lock.Lock()
	incrID.value +=1
	incrID.lock.Unlock()
	return value
}
func (incrID *incrementID) String() string {
	return gconv.IntString(incrID.Int())
}
func IncrID() incrementID {
	return incrementID{
		value: 1,
		lock: new(sync.RWMutex),
	}
}
var nameIncrIDMap = struct {
	hash map[string/* name */]*incrementID
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
type Model interface {
	TableName () string
}
func DBIncrID (model Model) string {
	return NameIncrID(model.TableName())
}