// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package tflite

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type LogicalOrOptions struct {
	_tab flatbuffers.Table
}

func GetRootAsLogicalOrOptions(buf []byte, offset flatbuffers.UOffsetT) *LogicalOrOptions {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &LogicalOrOptions{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *LogicalOrOptions) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *LogicalOrOptions) Table() flatbuffers.Table {
	return rcv._tab
}

func LogicalOrOptionsStart(builder *flatbuffers.Builder) {
	builder.StartObject(0)
}
func LogicalOrOptionsEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
