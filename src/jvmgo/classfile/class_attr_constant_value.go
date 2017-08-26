package classfile

type ConstantValueAttributeInfo struct {
	constantValueIndex uint16
}

func (self *ConstantValueAttributeInfo) readInfo(reader *ClassReader) {
	self.constantValueIndex = reader.readUint16()
}

func (self *ConstantValueAttributeInfo) ConstantValueIndex() uint16 {
	return self.constantValueIndex
}
