package classfile

import (
	"encoding/binary"
)

type ClassReader struct {
	data []byte
}

func (self *ClassReader) readUint8() uint8 {
	val := self.data[0]
	self.data = self.data[1:]
	return val
}

func (self *ClassReader) readUint16() uint16 {
	val := binary.BigEndian.Uint16(self.data)
	self.data = self.data[2:]
	return val
}

func (self *ClassReader) readUint32() uint32 {
	val := binary.BigEndian.Uint32(self.data)
	self.data = self.data[4:]
	return val
}

func (self *ClassReader) readUint64() uint64 {
	val := binary.BigEndian.Uint64(self.data)
	self.data = self.data[8:]
	return val
}

func (self *ClassReader) readUinit16s() []uint16 {
	length := self.readUint16()
	result := make([]uint16, length)
	for i := range result {
		result[i] = self.readUint16()
	}
	return result
}

func (self *ClassReader) readBytes(length uint32) []byte {
	bytes := self.data[0:length]
	self.data = self.data[length:]
	return bytes
}
