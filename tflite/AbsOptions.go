// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package tflite

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type AbsOptions struct {
	_tab flatbuffers.Table
}

func GetRootAsAbsOptions(buf []byte, offset flatbuffers.UOffsetT) *AbsOptions {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &AbsOptions{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *AbsOptions) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *AbsOptions) Table() flatbuffers.Table {
	return rcv._tab
}

func AbsOptionsStart(builder *flatbuffers.Builder) {
	builder.StartObject(0)
}
func AbsOptionsEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
