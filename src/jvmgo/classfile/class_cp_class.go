package classfile

type ConstantClassInfo struct {
	cp    ConstantPool
	index uint16
}

func (self *ConstantClassInfo) readInfo(reader *ClassReader) {
	self.index = reader.readUint16()
}

func (self *ConstantClassInfo) String() string {
	return self.cp.getUtf8(self.index)
}
