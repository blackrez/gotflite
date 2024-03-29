// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package tflite

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type RNNOptions struct {
	_tab flatbuffers.Table
}

func GetRootAsRNNOptions(buf []byte, offset flatbuffers.UOffsetT) *RNNOptions {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &RNNOptions{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *RNNOptions) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *RNNOptions) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *RNNOptions) FusedActivationFunction() int8 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetInt8(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *RNNOptions) MutateFusedActivationFunction(n int8) bool {
	return rcv._tab.MutateInt8Slot(4, n)
}

func RNNOptionsStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func RNNOptionsAddFusedActivationFunction(builder *flatbuffers.Builder, fusedActivationFunction int8) {
	builder.PrependInt8Slot(0, fusedActivationFunction, 0)
}
func RNNOptionsEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
