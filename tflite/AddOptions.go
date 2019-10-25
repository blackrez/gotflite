// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package tflite

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type AddOptions struct {
	_tab flatbuffers.Table
}

func GetRootAsAddOptions(buf []byte, offset flatbuffers.UOffsetT) *AddOptions {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &AddOptions{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *AddOptions) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *AddOptions) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *AddOptions) FusedActivationFunction() int8 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetInt8(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *AddOptions) MutateFusedActivationFunction(n int8) bool {
	return rcv._tab.MutateInt8Slot(4, n)
}

func AddOptionsStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func AddOptionsAddFusedActivationFunction(builder *flatbuffers.Builder, fusedActivationFunction int8) {
	builder.PrependInt8Slot(0, fusedActivationFunction, 0)
}
func AddOptionsEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}