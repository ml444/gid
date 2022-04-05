package convert

import (
	"github.com/ml444/gid/server"
)

type Converter interface {
	ConvertToGen(id server.Id) // 合成一个长整型的ID
	ConvertToExp(id int64)     // 拆解长整型的ID
}

type Convertor struct {
	meta *server.Meta
}

func NewConvertor(meta *server.Meta) *Convertor {
	return &Convertor{meta: meta}
}

func (c *Convertor) ConvertToGen(id *server.Id) uint64 {
	ret := uint64(0)
	ret |= id.GetMachine()
	ret |= id.GetSeq() << c.meta.GetSeqBitsStartPos()
	ret |= id.GetTime() << c.meta.GetTimeBitsStartPos()
	ret |= id.GetGenMethod() << c.meta.GetGenMethodBitsStartPos()
	ret |= id.GetType() << c.meta.GetMtypeBitsStartPos()
	ret |= id.GetVersion() << c.meta.GetVersionBitsStartPos()
	return ret
}

func (c *Convertor) ConvertToExp(id uint64) *server.Id {
	ret := new(server.Id)
	ret.SetMachine(id & c.meta.GetMachineBitsMask())
	ret.SetSeq((id >> c.meta.GetSeqBitsStartPos()) & c.meta.GetSeqBitsMask()) //无符号右移操作
	ret.SetTime((id >> c.meta.GetTimeBitsStartPos()) & c.meta.GetTimeBitsMask())
	ret.SetGenMethod((id >> c.meta.GetGenMethodBitsStartPos()) & c.meta.GetGenMethodBitsMask())
	ret.SetType((id >> c.meta.GetMtypeBitsStartPos()) & c.meta.GetMtypeBitsMask())
	ret.SetVersion((id >> c.meta.GetVersionBitsStartPos()) & c.meta.GetVersionBitsMask())
	return ret
}
