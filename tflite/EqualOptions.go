// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package tflite

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type EqualOptions struct {
	_tab flatbuffers.Table
}

func GetRootAsEqualOptions(buf []byte, offset flatbuffers.UOffsetT) *EqualOptions {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &EqualOptions{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *EqualOptions) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *EqualOptions) Table() flatbuffers.Table {
	return rcv._tab
}

func EqualOptionsStart(builder *flatbuffers.Builder) {
	builder.StartObject(0)
}
func EqualOptionsEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
