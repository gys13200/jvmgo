package rtda

import (
	"math"
)

type OperandStack struct {
	size uint
	slots []Slot
}

func newOperandStack(maxSize uint) *OperandStack  {
	if maxSize <= 0{
		return nil
	}
	return &OperandStack{
		slots:make([]Slot, maxSize),
	}
}

func (c *OperandStack) PushInt(val int32) {
	c.slots[c.size].num = val
	c.size++
}
func (c *OperandStack) PopInt() int32 {
	c.size--
	result := c.slots[c.size]
	return result.num
}

func (c *OperandStack) PushFloat(val float32) {
	c.slots[c.size].num = int32(math.Float32bits(val))
	c.size++
}
func (c *OperandStack) PopFloat() float32 {
	c.size--
	bits := uint32(c.slots[c.size].num)
	return math.Float32frombits(bits)
}

func (c *OperandStack) PushLong(val int64)  {
	c.PushInt(int32(val))
	c.PushInt(int32(val >> 32))
}
func (c *OperandStack) PopLong() uint64 {
	lower := uint32(c.PopInt())
	higher := uint32(c.PopInt())
	return uint64(higher) << 32 | uint64(lower)
}

func (c *OperandStack) PushDouble(val float64) {
	c.PushLong(int64(math.Float64bits(val)))
}
func (c *OperandStack) PopDouble() float64 {
	return math.Float64frombits(uint64(c.PopLong()))
}

func (c *OperandStack) PushRef(ref *Object) {
	c.slots[c.size].ref = ref;
}
func (c *OperandStack) PopRef() *Object {
	c.size--
	ref := c.slots[c.size].ref
	c.slots[c.size].ref = nil // help gc
	return ref
}
