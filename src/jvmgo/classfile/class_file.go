package classfile

import (
	"fmt"
)

type ClassFile struct {

	// 魔数
	magicNumber uint32

	// 次要版本号
	minorVersion uint16

	// 主要版本号
	majorVersion uint16

	// 常量池
	canstantPool ConstantPool

	// 访问标志
	accessFlags uint16

	// 本类全名
	thisClass uint16

	// 父类全名
	superClass uint16

	// 实现的接口
	interfaces []uint16

	// 所有属性
	fields []*MemberInfo

	// 所有方法
	methods []*MemberInfo

	attributes []AttributeInfo
}

func Parse(data []byte) (cf *ClassFile, err error) {

	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)

			if !ok {
				err = fmt.Errorf("%v", err)
			}
		}
	}()

	cr := &ClassReader{data}
	cf = &ClassFile{}
	cf.read(cr)
	return

}

func (self *ClassFile) read(reader *ClassReader) {
	self.readAndCheckMagicNumber(reader)
	self.readAndCheckVersion(reader)
}

func (self *ClassFile) readAndCheckMagicNumber(reader *ClassReader) {
	magic := reader.readUint32()
	if magic != 0xCAFEBABE {
		panic("java.lang.ClassFormt:magic")
	}
	self.magicNumber = magic
}

func (self *ClassFile) readAndCheckVersion(reader *ClassReader) {
	self.minorVersion = reader.readUint16()
	self.majorVersion = reader.readUint16()

	switch self.majorVersion {
	case 45:
		return
	case 45, 46, 47, 48, 49, 50, 51, 52:
		if self.minorVersion == 0 {
			return
		}
	}
	panic("java.long.UnsuportedClassVersion")

}

// getters
func (self *ClassFile) MinorVersion() uint16 {

}

func (self *ClassFile) MajorVersion() uint16 {

}

func (self *ClassFile) ConstantPool() ConstantPool {

}

func (self *ClassFile) Fields() []*MemberInfo {

}

func (self *ClassFile) Methods() []*MemberInfo {

}

func (self *ClassFile) AttributeInfos() []AttributeInfo {

}

func (self *ClassFile) SuperClassName() uint16 {

}

func (self *ClassFile) InterfaceNames() []uint16 {

}

func (self *ClassFile) ClassName() uint16 {

}
