package asnl

import (
	"reflect"
	"testing"
)

func TestBasic(t *testing.T) {
	var x = NewAml(32)

	x.Struct()
	x.Int(1, 10)
	x.Struct()
	x.String("p")
	x.String("1234")
	x.EndStruct()
	x.EndStruct()

	if !reflect.DeepEqual(x.buffer[0:x.pos], []byte{
		STRUCT, 0x0e, INT, 0x01, 0x0a, STRUCT, 0x09, STRING,
		0x01, 0x70, STRING, 0x04, 0x31, 0x32, 0x33, 0x34,
	}) {
		t.Error("Basic Test failed")
	}
}
