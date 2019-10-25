// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package tflite

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type FullyConnectedOptions struct {
	_tab flatbuffers.Table
}

func GetRootAsFullyConnectedOptions(buf []byte, offset flatbuffers.UOffsetT) *FullyConnectedOptions {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &FullyConnectedOptions{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *FullyConnectedOptions) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *FullyConnectedOptions) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *FullyConnectedOptions) FusedActivationFunction() int8 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetInt8(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *FullyConnectedOptions) MutateFusedActivationFunction(n int8) bool {
	return rcv._tab.MutateInt8Slot(4, n)
}

func (rcv *FullyConnectedOptions) WeightsFormat() int8 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetInt8(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *FullyConnectedOptions) MutateWeightsFormat(n int8) bool {
	return rcv._tab.MutateInt8Slot(6, n)
}

func (rcv *FullyConnectedOptions) KeepNumDims() byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetByte(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *FullyConnectedOptions) MutateKeepNumDims(n byte) bool {
	return rcv._tab.MutateByteSlot(8, n)
}

func FullyConnectedOptionsStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func FullyConnectedOptionsAddFusedActivationFunction(builder *flatbuffers.Builder, fusedActivationFunction int8) {
	builder.PrependInt8Slot(0, fusedActivationFunction, 0)
}
func FullyConnectedOptionsAddWeightsFormat(builder *flatbuffers.Builder, weightsFormat int8) {
	builder.PrependInt8Slot(1, weightsFormat, 0)
}
func FullyConnectedOptionsAddKeepNumDims(builder *flatbuffers.Builder, keepNumDims byte) {
	builder.PrependByteSlot(2, keepNumDims, 0)
}
func FullyConnectedOptionsEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
