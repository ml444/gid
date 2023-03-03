package core

type IConverter interface {
	ConvertToGen(id IId) uint64      // 合成一个长整型的ID
	ConvertToExp(id uint64, out IId) // 拆解长整型的ID
}

type Convertor struct {
	meta IMeta
}

func NewConvertor(meta IMeta) *Convertor {
	return &Convertor{meta: meta}
}

func (c *Convertor) ConvertToGen(id IId) uint64 {
	ret := uint64(0)
	ret |= id.GetDevice()
	ret |= id.GetSequence() << c.meta.GetSequenceBitsStartPos()
	ret |= id.GetTime() << c.meta.GetTimeBitsStartPos()
	ret |= id.GetMethod() << c.meta.GetMethodBitsStartPos()
	ret |= id.GetType() << c.meta.GetTypeBitsStartPos()
	ret |= id.GetVersion() << c.meta.GetVersionBitsStartPos()
	return ret
}

func (c *Convertor) ConvertToExp(id uint64, out IId) {
	out.SetDevice(id & c.meta.GetDeviceBitsMask())
	out.SetSequence((id >> c.meta.GetSequenceBitsStartPos()) & c.meta.GetSequenceBitsMask())
	out.SetTime((id >> c.meta.GetTimeBitsStartPos()) & c.meta.GetTimeBitsMask())
	out.SetMethod((id >> c.meta.GetMethodBitsStartPos()) & c.meta.GetMethodBitsMask())
	out.SetType((id >> c.meta.GetTypeBitsStartPos()) & c.meta.GetTypeBitsMask())
	out.SetVersion((id >> c.meta.GetVersionBitsStartPos()) & c.meta.GetVersionBitsMask())
	return
}
