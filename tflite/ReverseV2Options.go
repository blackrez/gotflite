// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package tflite

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type ReverseV2Options struct {
	_tab flatbuffers.Table
}

func GetRootAsReverseV2Options(buf []byte, offset flatbuffers.UOffsetT) *ReverseV2Options {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &ReverseV2Options{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *ReverseV2Options) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *ReverseV2Options) Table() flatbuffers.Table {
	return rcv._tab
}

func ReverseV2OptionsStart(builder *flatbuffers.Builder) {
	builder.StartObject(0)
}
func ReverseV2OptionsEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
