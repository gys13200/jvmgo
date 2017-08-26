package classfile

import ()

type AttributeInfo interface {
	readInfo(reader *ClassReader)
}

func readAttributes(reader *ClassReader, cp ConstantPool) []AttributeInfo {
	attrCount := reader.readUint16()
	attrs := make([]AttributeInfo, attrCount)

	for i := range attrs {
		attrs[i] = readAttribute(reader, cp)
	}

	return attrs
}

func readAttribute(reader *ClassReader, cp ConstantPool) AttributeInfo {
	nameIndex := reader.readUint16()
	nameString := cp.getUtf8(nameIndex)
	attrlen := reader.readUint32()
	attrInfo := newAttributeInfo(nameString, attrlen, cp)
	attrInfo.readInfo(reader)
	return attrInfo
}

func newAttributeInfo(attrName string, attrLen uint32, cp ConstantPool) AttributeInfo {
	switch attrName {
	case "Code":
		return &CodeAttributeInfo{
			cp: cp,
		}

	case "ConstantValue":
		return &ConstantValueAttributeInfo{}
	case "Deprecated":
		return &DeprecatedAttributeInfo{}
	case "Exceptions":
		return &ExceptionsAttributeInfo{}
	case "LineNumberTable":
		return &LineNumberTableAttributeInfo{}
	case "LocalVariableTable":
		return nil
	case "SourceFile":
		return &SourceFileAttributeInfo{cp: cp}
	case "Synthetic":
		return &SyntheticAttributeInfo{}
	default:
		return &UnparsedAttributeInfo{attrName, attrLen, nil}
	}
}
