package classfile

type ConstantStringInfo struct {
	cp    ConstantPool
	index uint16
}

func (self *ConstantStringInfo) readInfo(reader *ClassReader) {
	self.index = reader.readUint16()
}

func (self *ConstantStringInfo) String() string {
	return self.cp.getUtf8(self.index)
}
