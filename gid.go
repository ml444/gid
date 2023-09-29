package gid

import (
	"encoding/base64"
	"errors"
	"fmt"
	"strconv"
	"sync"

	"github.com/ml444/gid/core"
	"github.com/ml444/gid/strategy"
)

var _ IdServer = &IdGenerator{}

type IdServer interface {
	GenerateId() uint64
	ExplainId(id uint64) []uint64
	MakeId(timestamp int64, sequence uint64) (uint64, error)
	TransTime(time uint64) (int64, string)
}

type IdGenerator struct {
	seqIdx   int
	meta     *core.Meta
	strategy core.IStrategy
	timeOp   core.ITimeOp
	dataPool sync.Pool
}

func NewIdGenerator(tsIdx, seqIdx int, kv map[int]uint64, bitset []uint8, epoch int64, opts ...OptionFunc) (*IdGenerator, error) {
	lenBitset := len(bitset)
	if lenBitset == 0 || tsIdx >= lenBitset || seqIdx >= lenBitset {
		return nil, fmt.Errorf("the length of the bitset incoming parameter is error: %d", lenBitset)
	}
	if len(kv) != lenBitset {
		panic("the length and order of the incoming parameters kv and bitset should be consistent")
	}
	if len(strconv.FormatInt(epoch, 10)) != 13 {
		return nil, errors.New("incorrect length of incoming epoch")
	}
	timeUnit := core.TimeUnitMilliSecond
	if timeBit := bitset[tsIdx]; timeBit <= 33 {
		timeUnit = core.TimeUnitSecond
	}
	timeOp, err := core.NewTimeOp(epoch, timeUnit)
	if err != nil {
		return nil, err
	}
	idGen := IdGenerator{
		seqIdx: seqIdx,
		meta:   core.NewMeta(bitset...),
		timeOp: timeOp,
	}
	for _, optFunc := range opts {
		optFunc(&idGen)
	}
	if idGen.strategy == nil {
		// default strategy
		idGen.strategy = strategy.NewAtomicFiller(idGen.meta.GetBitMask(seqIdx), idGen.timeOp)
	}

	idGen.dataPool.New = func() interface{} {
		return core.NewData(tsIdx, seqIdx, kv)
	}
	return &idGen, nil
}

// GenerateId generate globally unique id.
func (s *IdGenerator) GenerateId() uint64 {
	d := s.dataPool.Get().(*core.Data)
	defer s.dataPool.Put(d)
	duration, sequence := s.strategy.Caught() // 这是一个抽象方法，调用子类的
	d.SetTimeDuration(duration)
	d.SetSequence(sequence)
	return s.meta.ConvertToGen(d)
}

// ExplainId parse the components of the unique id.
func (s *IdGenerator) ExplainId(id uint64) []uint64 {
	return s.meta.ConvertToExp(id)
}

// MakeId According to the incoming timestamp and sequence parameters,
// manually synthesize the unique id.
func (s *IdGenerator) MakeId(timestamp int64, sequence uint64) (uint64, error) {
	timeDuration, err := s.timeOp.ParseTimestampToDuration(timestamp)
	if err != nil {
		return 0, err
	}

	d := s.dataPool.Get().(*core.Data)
	defer s.dataPool.Put(d)
	d.SetSequence(sequence)
	d.SetTimeDuration(timeDuration)
	return s.meta.ConvertToGen(d), nil
}

// TransTime 转换时间
func (s *IdGenerator) TransTime(timeDuration uint64) (int64, string) {
	return s.timeOp.ParseDuration(timeDuration)
}

type LongId uint64

func (f *LongId) String() string {
	return strconv.FormatUint(uint64(*f), 10)
}

func (f *LongId) ToBinary() string {
	return strconv.FormatUint(uint64(*f), 10)
}

func (f *LongId) ToBase64() string {
	return base64.StdEncoding.EncodeToString([]byte(f.String()))
}

// MarshalJSON returns a json byte array string of the snowflake ID.
func (f *LongId) MarshalJSON() ([]byte, error) {
	buff := make([]byte, 0, 22)
	buff = append(buff, '"')
	buff = strconv.AppendUint(buff, uint64(*f), 10)
	buff = append(buff, '"')
	return buff, nil
}

// UnmarshalJSON converts a json byte array of a snowflake ID into an ID type.
func (f *LongId) UnmarshalJSON(b []byte) error {
	if len(b) < 3 || b[0] != '"' || b[len(b)-1] != '"' {
		return JSONSyntaxError{b}
	}

	i, err := strconv.ParseUint(string(b[1:len(b)-1]), 10, 64)
	if err != nil {
		return err
	}

	*f = LongId(i)
	return nil
}

// A JSONSyntaxError is returned from UnmarshalJSON if an invalid ID is provided.
type JSONSyntaxError struct{ original []byte }

func (j JSONSyntaxError) Error() string {
	return fmt.Sprintf("invalid snowflake ID %q", string(j.original))
}

// ParseBase64 converts a base64 string into a snowflake ID
func ParseBase64(id string) (LongId, error) {
	b, err := base64.StdEncoding.DecodeString(id)
	if err != nil {
		return 0, err
	}
	i, err := strconv.ParseInt(string(b), 10, 64)
	return LongId(i), err
	//return ParseBytes(b)
}
