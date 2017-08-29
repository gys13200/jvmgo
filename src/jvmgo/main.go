package main

import (
	"fmt"
	"jvmgo/classfile"
	"jvmgo/classpath"
	"jvmgo/rtda"
	"strings"
)

func main() {
	cmd := parseCmd()

	if cmd.versionFlag {
		fmt.Printf("version 0.0.1")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		startJVM(cmd)
	}
}

func startJVM(cmd *Cmd) {

	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)

	fmt.Printf("classpath: %s, class:%s, args: %v\n",
		cmd.cpOption, cmd.class, cmd.args)

	className := strings.Replace(cmd.class, ".", "/", -1)
	cf := loadClass(className, cp)
	printClassInfo(cf)

	frame := rtda.NewFrame(10, 10)
	testLocalVars(frame.LocalVars())
	testOperandStack(frame.OperandStack())
}

func testLocalVars(vars rtda.LocalVars)  {
	vars. SetInt( 0, 100)
	vars. SetInt( 1, -100)
	vars. SetLong( 2, 2997924580)
	vars. SetLong( 4, -2997924580)
	vars. SetFloat( 6, 3.1415926)
	vars. SetDouble( 7, 2.71828182845)
	vars. SetRef( 9, nil)
	println( vars. GetInt( 0))
	println( vars. GetInt( 1))
	println( vars. GetLong( 2))
	println( vars. GetLong( 4))
	println( vars. GetFloat( 6))
	println(vars.GetDouble(7))
	println(vars.GetRef(9))

}

func testOperandStack(ops *rtda.OperandStack) {
	ops. PushInt( 100)
	ops. PushInt(-100)
	ops. PushLong(2997924580)
	ops. PushLong(-2997924580)
	ops. PushFloat(3.1415926)
	ops. PushDouble(2.71828182845)
	ops. PushRef(nil)
	println(ops. PopRef())
	println( ops. PopDouble())
	println( ops. PopFloat())
	println( ops. PopLong())
	println( ops. PopLong())
	println( ops. PopInt())
	println( ops. PopInt())

}

func loadClass(className string, cp *classpath.Classpath) *classfile.ClassFile {
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		panic(err)
	}

	cf, err := classfile.Parse(classData)
	if err != nil {
		panic(err)
	}

	return cf
}

func printClassInfo(cf *classfile.ClassFile) {
	fmt.Printf("version:%v.%v \n", cf.MajorVersion(), cf.MinorVersion())
	fmt.Printf("constant count: %v \n", len(cf.ConstantPool()))
	fmt.Printf("accessFlags: 0x%x \n", cf.AccessFlags())
	fmt.Printf("thisClass: %s \n", cf.ClassName())
	fmt.Printf("superClass: %s \n", cf.SuperClassName())
	fmt.Printf("interfaces: %v \n", cf.InterfaceNames())
	fmt.Printf("fields count: %v \n", len(cf.Fields()))
	for _, item := range cf.Fields() {
		fmt.Printf("fieldName : %s ,fieldDescriptor: %s \n", item.Name(), item.Descriptor())
	}

	fmt.Printf("methods counts:%v \n", len(cf.Methods()))

	for _, item := range cf.Methods() {
		fmt.Printf("methodName: %s , descriptor: %s \n", item.Name(), item.Descriptor())
	}
}
