package cha

import (
	ge "github.com/og/x/error"
	gjson "github.com/og/x/json"
	glist "github.com/og/x/list"
	"github.com/pkg/errors"
	"reflect"
	"regexp"
	"strings"
)
const (
	fnUUID = "UUID"
	fnNameIncrID = "NameIncrID"
)
func fillByFnName(fnName string, args []interface{}, value reflect.Value, valueType reflect.StructField) {
	switch fnName {
	case fnUUID:
		value.SetString(UUID())
		break
	case fnNameIncrID:
		id := NameIncrID(reflect.ValueOf(args[0]).String())
		value.SetString(id)
	default:
		tag := mockTag + `:"` + valueType.Tag.Get(mockTag) + `"`

		panic(errors.New("chatty: "+ tag + "\r\n" + fnName + " can not found"))
	}
}
const mockTag = "cha"
func coreUnsafeMock(valuePtr reflect.Value) {
	value := valuePtr.Elem()
	valueType := value.Type()
	glist.Run(valueType.NumField(), func(typeIndex int) (_break bool) {
		fieldValue := value.Field(typeIndex)
		fieldType := valueType.Field(typeIndex)
		tag, has := fieldType.Tag.Lookup(mockTag)
		if fieldValue.Kind() == reflect.Struct {
			coreUnsafeMock(fieldValue.Addr())
			return
		}
		if !has { return }
		fnCall := strings.Split(tag, "()")
		switch len(fnCall) {
		case 1:
			fnName := strings.Split(tag, "(")[0]
			prefixRe, err:= regexp.Compile(`^.*\(`); ge.Check(err)
			argsString := tag
			argsString = prefixRe.ReplaceAllString(argsString, "")
			suffixRe, err:= regexp.Compile(`\)$`); ge.Check(err)
			argsString = suffixRe.ReplaceAllString(argsString, "")
			anyList := []interface{}{}
			gjson.Parse(`[` + argsString +`]`, &anyList)
			fillByFnName(fnName, anyList, fieldValue, fieldType)
		case 2:
			fnName := fnCall[0]
			fillByFnName(fnName, nil, fieldValue, fieldType)
		default:
			panic(errors.New("go-chatty: tag can not be empty string. " + fieldType.Name))
		}
		return
	})
}
func UnsafeMock(ptr interface{}) {
	coreUnsafeMock(reflect.ValueOf(ptr))
}
type Data interface {
	Chatty()
}
const mockDataChattyName = "Chatty"
func coreSafeMock(valuePtr reflect.Value) {
	value := valuePtr.Elem()
	valuePtrType := valuePtr.Type()
	valueType := value.Type()
	glist.Run(valuePtrType.NumMethod(), func(i int) (_break bool) {
			method := valuePtrType.Method(i)
			if method.Name == mockDataChattyName {
				method.Func.Call([]reflect.Value{valuePtr})
			}
			return
	})
	glist.Run(valueType.NumField(), func(typeIndex int) (_break bool) {
		fieldValue := value.Field(typeIndex)
		if fieldValue.Kind() == reflect.Struct {
			coreSafeMock(fieldValue.Addr())
			return
		}
		return
	})
}
func Mock(data Data) {
	value := reflect.ValueOf(data)
	coreSafeMock(value)
}
