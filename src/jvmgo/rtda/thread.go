package rtda

type Thread struct {
	pc    int
	stack *Stack
}

func NewTread() *Thread {
	return &Thread{
		stack:newStack(1024),
	};
}

func (c *Thread) PC() int{
	return c.pc;
}

func (c *Thread) SetPC(pc int) {
	c.pc = pc
}

func (c *Thread) PushFrame(frame *Frame)  {
	c.stack.push(frame)
}

func (c *Thread) PopFrame() *Frame {
	return c.stack.pop()
}

func (c *Thread) currentFrame() *Frame {
	return c.stack.top()
}