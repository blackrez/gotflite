// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package tflite

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type BidirectionalSequenceRNNOptions struct {
	_tab flatbuffers.Table
}

func GetRootAsBidirectionalSequenceRNNOptions(buf []byte, offset flatbuffers.UOffsetT) *BidirectionalSequenceRNNOptions {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &BidirectionalSequenceRNNOptions{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *BidirectionalSequenceRNNOptions) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *BidirectionalSequenceRNNOptions) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *BidirectionalSequenceRNNOptions) TimeMajor() byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetByte(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *BidirectionalSequenceRNNOptions) MutateTimeMajor(n byte) bool {
	return rcv._tab.MutateByteSlot(4, n)
}

func (rcv *BidirectionalSequenceRNNOptions) FusedActivationFunction() int8 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetInt8(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *BidirectionalSequenceRNNOptions) MutateFusedActivationFunction(n int8) bool {
	return rcv._tab.MutateInt8Slot(6, n)
}

func (rcv *BidirectionalSequenceRNNOptions) MergeOutputs() byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetByte(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *BidirectionalSequenceRNNOptions) MutateMergeOutputs(n byte) bool {
	return rcv._tab.MutateByteSlot(8, n)
}

func BidirectionalSequenceRNNOptionsStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func BidirectionalSequenceRNNOptionsAddTimeMajor(builder *flatbuffers.Builder, timeMajor byte) {
	builder.PrependByteSlot(0, timeMajor, 0)
}
func BidirectionalSequenceRNNOptionsAddFusedActivationFunction(builder *flatbuffers.Builder, fusedActivationFunction int8) {
	builder.PrependInt8Slot(1, fusedActivationFunction, 0)
}
func BidirectionalSequenceRNNOptionsAddMergeOutputs(builder *flatbuffers.Builder, mergeOutputs byte) {
	builder.PrependByteSlot(2, mergeOutputs, 0)
}
func BidirectionalSequenceRNNOptionsEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}