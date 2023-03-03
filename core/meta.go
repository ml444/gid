package core

type IMeta interface {
	SetVersionBits(versionBits byte)
	GetVersionBits() byte
	GetVersionBitsMask() uint64
	GetVersionBitsStartPos() byte
	SetTypeBits(typeBits byte)
	GetTypeBits() byte
	GetTypeBitsMask() uint64
	GetTypeBitsStartPos() byte
	SetMethodBits(genMethodBits byte)
	GetMethodBits() byte
	GetMethodBitsMask() uint64
	GetMethodBitsStartPos() byte
	SetTimeBits(timeBits byte)
	GetTimeBits() byte
	GetTimeBitsMask() uint64
	GetTimeBitsStartPos() byte
	SetSequenceBits(seqBits byte)
	GetSequenceBits() byte
	GetSequenceBitsMask() uint64
	GetSequenceBitsStartPos() byte
	SetDeviceBits(deviceBits byte)
	GetDeviceBits() byte
	GetDeviceBitsMask() uint64
}

var _ IMeta = &Meta{}

type Meta struct {
	machineBits  byte
	sequenceBits byte
	timeBits     byte
	methodBits   byte
	typeBits     byte
	versionBits  byte
}

func (m *Meta) GetDeviceBits() byte {
	return m.machineBits
}

func (m *Meta) SetDeviceBits(deviceBits byte) {
	m.machineBits = deviceBits
}

func (m *Meta) GetDeviceBitsMask() uint64 {
	// return -1L ^ -1L << machineBits
	res := -1 ^ (-1 << m.machineBits)
	return uint64(res)
}

func (m *Meta) GetSequenceBitsStartPos() byte {
	return m.machineBits
}

func (m *Meta) GetSequenceBits() byte {
	return m.sequenceBits
}

func (m *Meta) SetSequenceBits(seqBits byte) {
	m.sequenceBits = seqBits
}

func (m *Meta) GetSequenceBitsMask() uint64 {
	// return -1L ^ -1L << sequenceBits
	res := -1 ^ (-1 << m.sequenceBits)
	return uint64(res)
}

func (m *Meta) GetTimeBitsStartPos() byte {
	return m.machineBits + m.sequenceBits
}

func (m *Meta) GetTimeBits() byte {
	return m.timeBits
}

func (m *Meta) SetTimeBits(timeBits byte) {
	m.timeBits = timeBits
}

func (m *Meta) GetTimeBitsMask() uint64 {
	// return -1L ^ -1L << timeBits
	res := -1 ^ (-1 << m.timeBits)
	return uint64(res)
}

func (m *Meta) GetMethodBitsStartPos() byte {
	return m.machineBits + m.sequenceBits + m.timeBits
}

func (m *Meta) GetMethodBits() byte {
	return m.methodBits
}

func (m *Meta) SetMethodBits(genMethodBits byte) {
	m.methodBits = genMethodBits
}

func (m *Meta) GetMethodBitsMask() uint64 {
	// return -1L ^ -1L << methodBits
	res := -1 ^ (-1 << m.methodBits)
	return uint64(res)
}

func (m *Meta) GetTypeBitsStartPos() byte {
	// 10+10+30+2
	// 10+20+20+2
	return m.machineBits + m.sequenceBits + m.timeBits + m.methodBits
}

func (m *Meta) GetTypeBits() byte {
	return m.typeBits
}

func (m *Meta) SetTypeBits(typeBits byte) {
	m.typeBits = typeBits
}

func (m *Meta) GetTypeBitsMask() uint64 {
	// return -1L ^ -1L << mtypeBits
	res := -1 ^ (-1 << m.typeBits)
	return uint64(res)
}

func (m *Meta) GetVersionBitsStartPos() byte {
	return m.machineBits + m.sequenceBits + m.timeBits + m.methodBits + m.typeBits
}

func (m *Meta) GetVersionBits() byte {
	return m.versionBits
}

func (m *Meta) SetVersionBits(versionBits byte) {
	m.versionBits = versionBits
}

func (m *Meta) GetVersionBitsMask() uint64 {
	// return -1L ^ -1L << versionBits
	res := -1 ^ (-1 << m.versionBits)
	return uint64(res)
}

// // 工厂类，决定返回的idMeta是最大峰值还是最小粒度
// type IdMetaFactory struct {
// 	maxPeak        Meta
// 	minGranularity Meta
// }

func NewMetaWithMaxPeak() *Meta {
	return &Meta{
		machineBits:  10,
		sequenceBits: 20,
		timeBits:     30,
		methodBits:   2,
		typeBits:     1,
		versionBits:  1,
	}
}

func NewMetaWithMinGranularity() *Meta {
	return &Meta{
		machineBits:  10,
		sequenceBits: 10,
		timeBits:     40,
		methodBits:   2,
		typeBits:     1,
		versionBits:  1,
	}
}
