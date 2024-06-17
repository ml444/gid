package core

type IConverter interface {
	Generate(data *IDComponents) uint64 // Synthesize a long integer ID
	Parse(id uint64) *IDComponents      // Disassembling the ID of a long integer
}

type SegmentBits struct {
	DurationBits uint8
	WorkerIDBits uint8
	SequenceBits uint8
}

var _ IConverter = &Meta{}

type Meta struct {
	durationBits    uint8
	workerIDBits    uint8
	sequenceBits    uint8
	durationBitPos  uint8
	workerIDBitPos  uint8
	sequenceBitPos  uint8
	DurationBitMask uint64
	WorkerIDBitMask uint64
	SequenceBitMask uint64
}

func NewMeta(seg *SegmentBits) *Meta {
	m := &Meta{
		durationBits:    seg.DurationBits,
		workerIDBits:    seg.WorkerIDBits,
		sequenceBits:    seg.SequenceBits,
		durationBitPos:  seg.SequenceBits + seg.WorkerIDBits,
		workerIDBitPos:  seg.SequenceBits,
		sequenceBitPos:  0,
		DurationBitMask: 1 ^ (1 << seg.DurationBits) - 2,
		WorkerIDBitMask: 1 ^ (1 << seg.WorkerIDBits) - 2,
		SequenceBitMask: 1 ^ (1 << seg.SequenceBits) - 2,
	}

	return m
}

func (m *Meta) Generate(d *IDComponents) uint64 {
	ret := uint64(0)
	ret |= d.Sequence
	ret |= d.WorkerID << m.workerIDBitPos
	ret |= d.Duration << m.durationBitPos
	return ret
}

func (m *Meta) Parse(id uint64) *IDComponents {
	d := IDComponents{}
	d.Sequence = id & m.SequenceBitMask
	d.WorkerID = id >> m.workerIDBitPos & m.WorkerIDBitMask
	d.Duration = id >> m.durationBitPos & m.DurationBitMask
	return &d
}

type IDComponents struct {
	Duration uint64
	WorkerID uint64
	Sequence uint64
}
