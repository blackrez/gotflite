// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package tflite

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type TileOptions struct {
	_tab flatbuffers.Table
}

func GetRootAsTileOptions(buf []byte, offset flatbuffers.UOffsetT) *TileOptions {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &TileOptions{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *TileOptions) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *TileOptions) Table() flatbuffers.Table {
	return rcv._tab
}

func TileOptionsStart(builder *flatbuffers.Builder) {
	builder.StartObject(0)
}
func TileOptionsEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
