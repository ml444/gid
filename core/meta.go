package core

type IConverter interface {
	ConvertToGen(data IData) uint64  // Synthesize a long integer ID
	ConvertToExp(id uint64) []uint64 // Disassembling the ID of a long integer
}
type IData interface {
	GetValue(idx int) uint64
	SetTimeDuration(v uint64)
	SetSequence(v uint64)
}
type IMeta interface {
	//GetSequence
}

var _ IConverter = &Meta{}
var _ IData = &Data{}

type Data struct {
	kv     map[int]uint64
	tsIdx  int
	seqIdx int
}

func NewData(tsIdx, seqIdx int, kv map[int]uint64) *Data {
	return &Data{
		kv:     kv,
		tsIdx:  tsIdx,
		seqIdx: seqIdx,
	}
}
func (m *Data) GetValue(idx int) uint64 {
	return m.kv[idx]
}
func (m *Data) SetTimeDuration(v uint64) {
	m.kv[m.tsIdx] = v
}
func (m *Data) SetSequence(v uint64) {
	m.kv[m.seqIdx] = v
}

type Meta struct {
	kv              Data
	bitList         []uint8
	bitStartPosList []uint8
	bitsLen         int
}

func NewMeta(bits ...uint8) *Meta {
	var spList []uint8
	for i, _ := range bits {
		var v uint8
		if i == 0 {
			v = 0
		} else {
			for j := i - 1; j >= 0; j-- {
				v = v + bits[j]
			}
		}
		spList = append(spList, v)
	}
	return &Meta{
		bitList:         bits,
		bitsLen:         len(bits),
		bitStartPosList: spList,
	}
}

func (m *Meta) GetBit(idx int) uint8 {
	return m.bitList[idx]
}
func (m *Meta) GetBitMask(idx int) uint64 {
	res := -1 ^ (-1 << m.bitList[idx])
	return uint64(res)
}
func (m *Meta) GetBitStartPos(idx int) uint8 {
	return m.bitStartPosList[idx]
}

func (m *Meta) ConvertToGen(data IData) uint64 {
	ret := uint64(0)
	for i := 0; i < m.bitsLen; i++ {
		ret |= data.GetValue(i) << m.GetBitStartPos(i)
	}
	return ret
}

func (m *Meta) ConvertToExp(id uint64) []uint64 {
	var out []uint64
	for i := 0; i < m.bitsLen; i++ {
		var v uint64
		if i == 0 {
			v = id & m.GetBitMask(i)
		} else {
			v = (id >> m.GetBitStartPos(i)) & m.GetBitMask(i)
		}
		out = append(out, v)
	}
	return out
}
