// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package tflite

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type ConcatenationOptions struct {
	_tab flatbuffers.Table
}

func GetRootAsConcatenationOptions(buf []byte, offset flatbuffers.UOffsetT) *ConcatenationOptions {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &ConcatenationOptions{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *ConcatenationOptions) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *ConcatenationOptions) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *ConcatenationOptions) Axis() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *ConcatenationOptions) MutateAxis(n int32) bool {
	return rcv._tab.MutateInt32Slot(4, n)
}

func (rcv *ConcatenationOptions) FusedActivationFunction() int8 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetInt8(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *ConcatenationOptions) MutateFusedActivationFunction(n int8) bool {
	return rcv._tab.MutateInt8Slot(6, n)
}

func ConcatenationOptionsStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func ConcatenationOptionsAddAxis(builder *flatbuffers.Builder, axis int32) {
	builder.PrependInt32Slot(0, axis, 0)
}
func ConcatenationOptionsAddFusedActivationFunction(builder *flatbuffers.Builder, fusedActivationFunction int8) {
	builder.PrependInt8Slot(1, fusedActivationFunction, 0)
}
func ConcatenationOptionsEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
