package asnl

import (
	"fmt"
)

const (
	INT    = 1
	UINT   = 2
	STRING = 3
	STRUCT = 4
)

type aml struct {
	buffer []byte
	stack  int
	pos    int
}

func NewAml(bufferSize int) *aml {
	a := new(aml)
	a.buffer = make([]byte, bufferSize)
	a.stack = 0
	a.pos = 0
	return a
}

func (a *aml) Int(size int, value int) {
	a.buffer[a.pos] = INT
	a.pos++
	a.buffer[a.pos] = byte(size)
	a.pos++
	for i := 0; i < size; i++ {
		a.buffer[a.pos] = byte(value % 256)
		value = value / 256
		a.pos++
	}
}

func (a *aml) Uint(size int, value int) {
	a.buffer[a.pos] = UINT
	a.pos++
	a.buffer[a.pos] = byte(size)
	a.pos++
	for i := 0; i < size; i++ {
		a.buffer[a.pos] = byte(value % 256)
		value = value / 256
		a.pos++
	}
}

func (a *aml) String(value string) {
	a.buffer[a.pos] = STRING
	a.pos++
	a.buffer[a.pos] = byte(len(value))
	a.pos++
	for i := 0; i < len(value); i++ {
		a.buffer[a.pos] = byte(value[i])
		a.pos++
	}
}

func (a *aml) Struct() {
	a.buffer[a.pos] = STRUCT
	a.pos++
	a.buffer[a.pos] = byte(a.stack)
	a.stack = a.pos
	a.pos++
}

func (a *aml) EndStruct() {
	if a.stack > 0 {
		i := int(a.stack)
		a.stack = int(a.buffer[i])
		a.buffer[i] = byte(a.pos - i - 1)
	}
}

func (a aml) Dump() {
	for i := 0; i < a.pos; i++ {
		fmt.Printf("%02x ", a.buffer[i])
	}
	fmt.Println()
}
