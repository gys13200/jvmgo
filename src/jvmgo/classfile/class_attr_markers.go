package classfile

type MarkerAttributeInfo struct{}

func (self *MarkerAttributeInfo) readInfo(reader *ClassReader) {
	// do nothing
}

type DeprecatedAttributeInfo struct{ MarkerAttributeInfo }

type SyntheticAttributeInfo struct{ MarkerAttributeInfo }
