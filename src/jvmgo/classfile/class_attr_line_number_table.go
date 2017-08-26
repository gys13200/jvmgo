package classfile

type LineNumberTableAttributeInfo struct {
	lineNumberTable []*LineNumberTableEntry
}

type LineNumberTableEntry struct {
	startPc    uint16
	lineNumber uint16
}

func (self *LineNumberTableAttributeInfo) readInfo(reader *ClassReader) {
	tableLen := reader.readUint16()
	lineTable := make([]*LineNumberTableEntry, tableLen)
	for i := range lineTable {
		lineTable[i] = &LineNumberTableEntry{
			startPc:    reader.readUint16(),
			lineNumber: reader.readUint16(),
		}
	}

	self.lineNumberTable = lineTable
}
