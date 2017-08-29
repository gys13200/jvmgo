package rtda

import (
	"math"
)
type LocalVars []Slot

type Slot struct {
	num int32
	ref *Object
}

func newLocalVars(maxLocals uint) LocalVars {
	if maxLocals > 0 {
		return make([]Slot, maxLocals)
	}
	return nil
}

func (c LocalVars) SetInt(index uint, val int32)  {
	c[index].num = val
}
func (c LocalVars) GetInt(index uint) int32 {
	return c[index].num;
}


func (c LocalVars) SetFloat(index uint, val float32)  {
	c[index].num = int32(math.Float32bits(val))
}
func (c LocalVars) GetFloat(index uint) float32 {
	bits := uint32(c[index].num)
	return math.Float32frombits(bits);
}

func (c LocalVars) SetLong(index uint, val int64)  {
	c[index+1].num = int32(val)
	c[index].num = int32(val >> 32)
}
func (c LocalVars) GetLong(index uint) int64 {
	higher := uint32(c[index].num)
	lower := uint32(c[index+1].num)
	return int64(higher) << 32 | int64(lower);
}

func (c LocalVars) SetDouble(index uint, val float64) {
	l := int64(math.Float64bits(val))
	c.SetLong(index, l)
}
func (c LocalVars) GetDouble(index uint) float64  {
	l := uint64(c.GetLong(index))
	return math.Float64frombits(l)
}

func (c LocalVars) SetRef(index uint, ref *Object)  {
	c[index].ref = ref
}
func (c LocalVars)GetRef(index uint) *Object {
	return c[index].ref
}



