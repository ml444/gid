package core

import (
	"testing"
)

var seg = &SegmentBits{
	DurationBits: 32,
	WorkerIDBits: 10,
	SequenceBits: 21,
}

func TestMeta_ConvertToExp(t *testing.T) {
	m1 := NewMeta(seg)
	d := &IDComponents{
		Duration: 123,
		WorkerID: 111,
		Sequence: 123,
	}

	id1 := m1.Generate(d)
	t.Log(id1)
}

func BenchmarkMeta_ConvertToGen(b *testing.B) {
	m := NewMeta(seg)
	d := &IDComponents{
		Duration: 123,
		WorkerID: 111,
		Sequence: 123,
	}
	for i := 0; i < b.N; i++ {
		d.Sequence = uint64(i)
		d.Duration = uint64(i)
		m.Generate(d)
	}
}

func BenchmarkMeta_ConvertToGen2(b *testing.B) {
	for i := 0; i < b.N; i++ {
	}
}
