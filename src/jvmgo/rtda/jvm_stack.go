package rtda

type Stack struct {
	maxSize uint
	size uint
	_top *Frame
}

func newStack(maxSize uint) *Stack {
	return &Stack{
		maxSize:maxSize,
	}
}

/*
   链表push操作
 */
func (c *Stack) push(frame *Frame) {
	if c.size > c.maxSize {
		panic("java.lang.StackOverflowError")
	}
	if c._top == nil {
		c._top = frame
		return
	}
	frame.lower = c._top
	c._top = frame
	c.size++

}

func (c *Stack) pop() *Frame {

	if c._top == nil {
		panic("empty stack")
	}
	result := c.top()
	c._top = result.lower
	result.lower = nil
	c.size--
	return result
}

func (c *Stack) top() *Frame {
	if c._top == nil {
		panic("empty stack")
	}
	return c._top
}