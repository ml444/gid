package core


type Metaer interface {
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
