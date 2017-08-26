package classfile

type SourceFileAttributeInfo struct {
	cp            ConstantPool
	fileNameIndex uint16
}

func (self *SourceFileAttributeInfo) readInfo(reader *ClassReader) {
	self.fileNameIndex = reader.readUint16()
}

func (self *SourceFileAttributeInfo) FileName() string {
	return self.cp.getUtf8(self.fileNameIndex)
}
