package net

import (
	"bufio"
	"strings"
)

type Packet struct {
	payload     []byte
	readerIndex int
	writerIndex int
}

func NewPacket(capacity int) *Packet {
	return &Packet{payload: make([]byte, capacity)}
}

func (pkt *Packet) ReadInt8() int8 {
	v := int8(pkt.payload[pkt.readerIndex])
	pkt.readerIndex += 1
	return v
}

func (pkt *Packet) ReadBool() bool {
	return pkt.ReadInt8() == 1
}

func (pkt *Packet) ReadInt16() int16 {
	v1 := pkt.ReadInt8()
	v2 := pkt.ReadInt8()

	return int16(v1)<<8 | int16(v2)
}

func (pkt *Packet) ReadInt24() int32 {
	v1 := uint8(pkt.ReadInt8())
	v2 := uint8(pkt.ReadInt8())
	v3 := uint8(pkt.ReadInt8())

	return int32(v1)<<16 | int32(v2)<<8 | int32(v3)
}

func (pkt *Packet) ReadInt32() int32 {
	v1 := uint8(pkt.ReadInt8())
	v2 := uint8(pkt.ReadInt8())
	v3 := uint8(pkt.ReadInt8())
	v4 := uint8(pkt.ReadInt8())

	return int32(v1)<<24 | int32(v2)<<16 | int32(v3)<<8 | int32(v4)
}

func (pkt *Packet) ReadInt48() int {
	v1 := uint8(pkt.ReadInt8())
	v2 := uint8(pkt.ReadInt8())
	v3 := uint8(pkt.ReadInt8())
	v4 := uint8(pkt.ReadInt8())
	v5 := uint8(pkt.ReadInt8())
	v6 := uint8(pkt.ReadInt8())

	return int(v1)<<40 | int(v2)<<32 | int(v3)<<24 | int(v4)<<16 | int(v5)<<8 | int(v6)
}

func (pkt *Packet) ReadInt64() int64 {
	v1 := uint8(pkt.ReadInt8())
	v2 := uint8(pkt.ReadInt8())
	v3 := uint8(pkt.ReadInt8())
	v4 := uint8(pkt.ReadInt8())
	v5 := uint8(pkt.ReadInt8())
	v6 := uint8(pkt.ReadInt8())
	v7 := uint8(pkt.ReadInt8())
	v8 := uint8(pkt.ReadInt8())

	return int64(v1)<<56 | int64(v2)<<48 | int64(v3)<<40 | int64(v4)<<32 | int64(v5)<<24 | int64(v6)<<16 | int64(v7)<<8 | int64(v8)
}

func (pkt *Packet) ReadCString() string {
	var bldr strings.Builder
	var character uint8 = 0

	for pkt.IsReadable() {
		character = uint8(pkt.ReadInt8())
		if character == 10 {
			break
		}

		bldr.WriteByte(character)
	}

	return bldr.String()
}

func (pkt *Packet) Fill(reader *bufio.Reader, amount int) {
	reader.Read(pkt.payload[:amount])

	if (pkt.writerIndex + amount) >= pkt.writerIndex {
		pkt.writerIndex += amount
	}

	pkt.readerIndex = 0
}

func (pkt *Packet) WriteInt8(value int) {
	pkt.payload[pkt.writerIndex] = byte(value)
	pkt.writerIndex += 1
}

func (pkt *Packet) WriteBool(value bool) {
	if value {
		pkt.WriteInt8(1)
	} else {
		pkt.WriteInt8(0)
	}
}

func (pkt *Packet) WriteInt16(value int) {
	pkt.WriteInt8(value >> 8)
	pkt.WriteInt8(value)
}

func (pkt *Packet) WriteInt24(value int) {
	pkt.WriteInt8(value >> 16)
	pkt.WriteInt8(value >> 8)
	pkt.WriteInt8(value)
}

func (pkt *Packet) WriteInt32(value int) {
	pkt.WriteInt8(value >> 24)
	pkt.WriteInt8(value >> 16)
	pkt.WriteInt8(value >> 8)
	pkt.WriteInt8(value)
}

func (pkt *Packet) WriteInt48(value int) {
	pkt.WriteInt8(value >> 40)
	pkt.WriteInt8(value >> 32)
	pkt.WriteInt8(value >> 24)
	pkt.WriteInt8(value >> 16)
	pkt.WriteInt8(value >> 8)
	pkt.WriteInt8(value)
}

func (pkt *Packet) WriteInt64(value int64) {
	pkt.WriteInt8(int(value >> 56))
	pkt.WriteInt8(int(value >> 48))
	pkt.WriteInt8(int(value >> 40))
	pkt.WriteInt8(int(value >> 32))
	pkt.WriteInt8(int(value >> 24))
	pkt.WriteInt8(int(value >> 16))
	pkt.WriteInt8(int(value >> 8))
	pkt.WriteInt8(int(value))
}

func (pkt *Packet) WriteCString(value string) {
	for i := 0; i < len(value); i++ {
		pkt.WriteInt8(int(value[i]))
	}

	pkt.WriteInt8(10)
}

func (pkt *Packet) ByteBlock(block func()) {
	offset := pkt.writerIndex
	pkt.WriteInt8(0)

	block()

	blockSize := (pkt.writerIndex - offset) - 1

	pkt.payload[offset] = byte(blockSize)
}

func (pkt *Packet) ShortBlock(block func()) {
	offset := pkt.writerIndex
	pkt.WriteInt16(0)

	block()

	blockSize := (pkt.writerIndex - offset) - 2

	pkt.payload[offset] = byte(blockSize >> 8)
	pkt.payload[offset + 1] = byte(blockSize)
}

func (pkt *Packet) WriteAndFlush(writer *bufio.Writer) {
	writer.Write(pkt.payload[:pkt.writerIndex])
	writer.Flush()

	pkt.writerIndex = 0
	pkt.readerIndex = 0
}

func (pkt *Packet) IsReadable() bool {
	return pkt.ReadableBytes() > 0
}

func (pkt *Packet) ReadableBytes() int {
	return pkt.writerIndex - pkt.readerIndex
}

func (pkt *Packet) IsWritable() bool {
	return pkt.WriteableBytes() > 0
}

func (pkt *Packet) WriteableBytes() int {
	return len(pkt.payload) - pkt.writerIndex
}