package classfile

type UnparsedAttributeInfo struct {
	attrName string
	attrLen  uint32
	info     []byte
}

func (self *UnparsedAttributeInfo) readInfo(reader *ClassReader) {
	self.info = reader.readBytes(self.attrLen)
}
