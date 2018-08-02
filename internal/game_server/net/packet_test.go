package net

import (
	"testing"
)

func TestPacket_WriteInt8(t *testing.T) {
	writtenValue := 16

	pkt := NewPacket(32)
	pkt.WriteInt8(writtenValue)

	if int(pkt.ReadInt8()) != writtenValue {
		t.Errorf("read value did not equal written value of %v", writtenValue)
	}
}

func TestPacket_WriteBool(t *testing.T) {
	writtenValue := true

	pkt := NewPacket(32)
	pkt.WriteBool(writtenValue)

	if pkt.ReadBool() != writtenValue {
		t.Errorf("read value did not equal written value of %v", writtenValue)
	}
}

func TestPacket_WriteInt16(t *testing.T) {
	writtenValue := 524

	pkt := NewPacket(32)
	pkt.WriteInt16(writtenValue)

	readValue := int(pkt.ReadInt16())

	if readValue != writtenValue {
		t.Errorf("read value of %v did not equal written value of %v", readValue, writtenValue)
	}
}

func TestPacket_WriteInt24(t *testing.T) {
	writtenValue := 3448484

	pkt := NewPacket(32)
	pkt.WriteInt24(writtenValue)

	readValue := int(pkt.ReadInt24())

	if readValue != writtenValue {
		t.Errorf("read value of %v did not equal written value of %v", readValue, writtenValue)
	}
}

func TestPacket_WriteInt32(t *testing.T) {
	writtenValue := 1838383

	pkt := NewPacket(32)
	pkt.WriteInt32(writtenValue)

	readValue := int(pkt.ReadInt32())

	if readValue != writtenValue {
		t.Errorf("read value of %v did not equal written value of %v", readValue, writtenValue)
	}
}

func TestPacket_WriteInt48(t *testing.T) {
	writtenValue := 1838383344

	pkt := NewPacket(32)
	pkt.WriteInt48(int(writtenValue))

	readValue := int(pkt.ReadInt48())

	if readValue != writtenValue {
		t.Errorf("read value of %v did not equal written value of %v", readValue, writtenValue)
	}
}

func TestPacket_WriteInt64(t *testing.T) {
	writtenValue := 183838334534534

	pkt := NewPacket(32)
	pkt.WriteInt64(int64(writtenValue))

	readValue := int(pkt.ReadInt64())

	if readValue != writtenValue {
		t.Errorf("read value of %v did not equal written value of %v", readValue, writtenValue)
	}
}

func TestPacket_ByteBlock(t *testing.T) {
	pkt := NewPacket(32)
	pkt.ByteBlock(func() {
		pkt.WriteInt8(0)
		pkt.WriteInt16(1)
		pkt.WriteInt32(2)
	})

	blockSize := pkt.ReadInt8()
	if blockSize != 7 {
		t.Errorf("block size of %v did not equal expected size of %v", blockSize, 7)
	}
}

func TestPacket_ShortBlock(t *testing.T) {
	pkt := NewPacket(32)
	pkt.ShortBlock(func() {
		pkt.WriteInt8(0)
		pkt.WriteInt16(1)
		pkt.WriteInt32(2)
	})

	blockSize := pkt.ReadInt16()
	if blockSize != 7 {
		t.Errorf("block size of %v did not equal expected size of %v", blockSize, 7)
	}
}
