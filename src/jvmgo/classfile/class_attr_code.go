package classfile

type CodeAttributeInfo struct {
	cp                  ConstantPool
	maxStack            uint16
	maxLocals           uint16
	code                []byte
	exceptionTableEntry []*ExceptionTableEntry
	attributes          []AttributeInfo
}

type ExceptionTableEntry struct {
	startPc   uint16
	endPc     uint16
	handlerPc uint16
	catchType uint16
}

func (self *CodeAttributeInfo) readInfo(reader *ClassReader) {
	self.maxStack = reader.readUint16()
	self.maxLocals = reader.readUint16()
	codeLen := reader.readUint32()
	self.code = reader.readBytes(codeLen)
	self.exceptionTableEntry = readExceptionTable(reader)
	self.attributes = readAttributes(reader, self.cp)
}

func readExceptionTable(reader *ClassReader) []*ExceptionTableEntry {
	exceptionTableLen := reader.readUint16()
	results := make([]*ExceptionTableEntry, exceptionTableLen)
	for i := range results {
		results[i] = &ExceptionTableEntry{
			startPc:   reader.readUint16(),
			endPc:     reader.readUint16(),
			handlerPc: reader.readUint16(),
			catchType: reader.readUint16(),
		}
	}

	return results
}
