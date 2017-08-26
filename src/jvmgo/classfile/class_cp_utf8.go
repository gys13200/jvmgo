package classfile

import ()

type ConstantUtf8Info struct {
	str string
}

func (self *ConstantUtf8Info) readInfo(reader *ClassReader) {
	length := uint32(reader.readUint16())
	bytes := reader.readBytes(length)
	self.str = decodeUtf8(bytes)
}

func decodeUtf8(bytes []byte) string {
	return string(bytes)
}
