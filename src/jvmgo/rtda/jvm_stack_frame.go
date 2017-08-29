package rtda

type Frame struct {

	lower *Frame
	localVars LocalVars
	operandStack *OperandStack
}

func NewFrame(maxLocals uint, maxStack uint) *Frame {
	return &Frame{
		localVars: newLocalVars(maxLocals),
		operandStack:newOperandStack(maxStack),
	}
}

func (c *Frame) LocalVars() LocalVars {
	return c.localVars
}

func (c *Frame) OperandStack() *OperandStack {
	return c.operandStack
}
