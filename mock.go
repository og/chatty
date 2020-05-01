package cha

import (
	ge "github.com/og/x/error"
	gjson "github.com/og/x/json"
	"github.com/pkg/errors"
	"reflect"
	"regexp"
	"strings"
)
func dict () (dict struct{
	FN struct{
		UUID string
		NameIncrID  string
		Letter string
		CapitalLetter string
		FirstName string
		LastName string
		Name string
		FullName string
		CFirstName string
		CLastName string
		CName string
		Int string
		Bool string
		TrueLikelihood string
	}
}) {
	dict.FN.UUID = "UUID"
	dict.FN.NameIncrID = "NameIncrID"
	dict.FN.Letter = "Letter"
	dict.FN.CapitalLetter = "CapitalLetter"
	dict.FN.FirstName = "FirstName"
	dict.FN.LastName = "LastName"
	dict.FN.Name = "Name"
	dict.FN.FullName = "FullName"
	dict.FN.CFirstName = "CFirstName"
	dict.FN.CLastName = "CLastName"
	dict.FN.CName = "CName"
	dict.FN.Int = "Int"
	dict.FN.Bool = "Bool"
	dict.FN.TrueLikelihood = "TrueLikelihood"
	return
}

func fillByFnName(funcName string, args []interface{}, value reflect.Value, valueType reflect.StructField) {
	switch funcName {
	case dict().FN.UUID:
		value.SetString(UUID())
	case dict().FN.NameIncrID:
		id := NameIncrID(reflect.ValueOf(args[0]).String())
		value.SetString(id)
	case dict().FN.Letter:
		floatValue := reflect.ValueOf(args[0]).Float()
		value.SetString(Letter(int(floatValue)))
	case dict().FN.CapitalLetter:
		floatValue := reflect.ValueOf(args[0]).Float()
		value.SetString(CapitalLetter(int(floatValue)))
	case dict().FN.FirstName:
		value.SetString(FirstName())
	case dict().FN.LastName:
		value.SetString(LastName())
	case dict().FN.Name:
		value.SetString(Name())
	case dict().FN.FullName:
		value.SetString(FullName())
	case dict().FN.CFirstName:
		value.SetString(CFirstName())
	case dict().FN.CLastName:
		value.SetString(CLastName())
	case dict().FN.CName:
		value.SetString(CName())
	case dict().FN.Int:
		min := int(reflect.ValueOf(args[0]).Float())
		max := int(reflect.ValueOf(args[1]).Float())
		value.SetInt(int64(Int(min, max)))
	case dict().FN.Bool:
		value.SetBool(Bool())
	case dict().FN.TrueLikelihood:
		likelihood := int(reflect.ValueOf(args[0]).Float())
		value.SetBool(TrueLikelihood(likelihood))
	default:
		tag := mockTag + `:"` + valueType.Tag.Get(mockTag) + `"`
		panic(errors.New("chatty: "+ tag + "\r\n" + funcName + " can not found"))
	}
}
const mockTag = "cha"
func coreUnsafeMock(valuePtr reflect.Value) {
	value := valuePtr.Elem()
	valueType := value.Type()
	Run(valueType.NumField(), func(typeIndex int) (_break bool) {
		fieldValue := value.Field(typeIndex)
		fieldType := valueType.Field(typeIndex)
		tag, has := fieldType.Tag.Lookup(mockTag)
		if fieldValue.Kind() == reflect.Struct {
			coreUnsafeMock(fieldValue.Addr())
			return
		}
		if !has { return }
		fnCall := strings.Split(tag, "()")
		if tag == "" {
			fnCall = []string{}
		}
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
type Dataer interface {
	Chatty()
}
const mockDataChattyName = "Chatty"
func coreSafeMock(valuePtr reflect.Value) {
	value := valuePtr.Elem()
	valuePtrType := valuePtr.Type()
	valueType := value.Type()
	Run(valuePtrType.NumMethod(), func(i int) (_break bool) {
			method := valuePtrType.Method(i)
			if method.Name == mockDataChattyName {
				method.Func.Call([]reflect.Value{valuePtr})
			}
			return
	})
	Run(valueType.NumField(), func(typeIndex int) (_break bool) {
		fieldValue := value.Field(typeIndex)
		if fieldValue.Kind() == reflect.Struct {
			coreSafeMock(fieldValue.Addr())
			return
		}
		return
	})
}
func Mock(data Dataer) {
	value := reflect.ValueOf(data)
	coreSafeMock(value)
}
