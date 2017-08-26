package classfile

type ExceptionsAttributeInfo struct {
	exceptionIndexTable []uint16
}

func (self *ExceptionsAttributeInfo) readInfo(reader *ClassReader) {
	self.exceptionIndexTable = reader.readUinit16s()
}

func (self *ExceptionsAttributeInfo) ExceptionIndexTable() []uint16 {
	return self.exceptionIndexTable
}
