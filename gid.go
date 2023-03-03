package gid

import (
	"encoding/base64"
	"fmt"
	"github.com/ml444/gid/core"
	"github.com/ml444/gid/filler"
	"strconv"
	"sync"
)

type IdServer interface {
	GenerateId() uint64
	ExplainId(id uint64) *core.Id
	MakeId(timestamp int64, seq int64) uint64
	TransTime(time uint64) (int64, string)
}

type IdGenerator struct {
	Version  uint64
	Type     uint64
	Method   uint64
	DeviceId uint64

	fillerType FillerType
	meta       core.IMeta
	filler     core.IFiller
	convertor  core.IConverter
	timeOp     core.ITimeOp

	idPool sync.Pool
}

func NewIdGenerator(opts ...OptionFunc) *IdGenerator {
	idGen := IdGenerator{}
	for _, optFunc := range opts {
		optFunc(&idGen)
	}
	if idGen.Version == 0 {
		panic("IdGenerator.Version need to be set")
	}
	if idGen.Type == 0 {
		panic("IdGenerator.Type need to be set")
	}
	if idGen.meta == nil || idGen.timeOp == nil {
		panic("IdGenerator.meta or IdGenerator.timeOp need to be set")
	}
	if idGen.Method == 0 {
		panic("IdGenerator.Method need to be set")
	}
	if idGen.DeviceId == 0 {
		panic("IdGenerator.DeviceId need to be set")
	}
	if idGen.filler == nil {
		if idGen.fillerType == 0 {
			panic("IdGenerator.filler need to be set")
		}
		switch idGen.fillerType {
		case FillerTypeSync:
			idGen.filler = filler.NewSyncFiller(idGen.meta, idGen.timeOp)
		case FillerTypeAtomic:
			idGen.filler = filler.NewAtomicFiller(idGen.meta, idGen.timeOp)
		case FillerTypeLock:
			idGen.filler = filler.NewLockFiller(idGen.meta, idGen.timeOp)
		default:
			panic("other types of filler are not supported yet")
		}
	}

	if idGen.filler == nil {

	}
	if idGen.convertor == nil {
		idGen.convertor = core.NewConvertor(idGen.meta)
	}
	idGen.idPool.New = func() interface{} {
		return core.NewId(idGen.Version, idGen.Type, idGen.Method, idGen.DeviceId)
	}
	return &idGen
}

// GenerateId generate globally unique id.
func (s *IdGenerator) GenerateId() uint64 {
	id := s.idPool.Get().(*core.Id)
	defer s.idPool.Put(id)
	s.populateId(id) // 这是一个抽象方法，调用子类的
	return s.convertor.ConvertToGen(id)
}
func (s *IdGenerator) GenerateId2() uint64 {
	id := s.idPool.Get().(*core.Id)
	defer s.idPool.Put(id)
	return s.convertor.ConvertToGen(id)
}

// ExplainId parse the components of the unique id.
func (s *IdGenerator) ExplainId(id uint64, out core.IId) {
	s.convertor.ConvertToExp(id, out)
}

// MakeId According to the incoming timestamp and sequence parameters,
// manually synthesize the unique id.
func (s *IdGenerator) MakeId(timestamp int64, sequence uint64) (uint64, error) {
	timeDuration, err := s.timeOp.ParseTimestampToDuration(timestamp)
	if err != nil {
		return 0, err
	}

	id := s.idPool.Get().(*core.Id)
	defer s.idPool.Put(id)
	id.SetSequence(sequence)
	id.SetTime(timeDuration)
	return s.convertor.ConvertToGen(id), nil
}

// TransTime 转换时间
func (s *IdGenerator) TransTime(timeDuration uint64) (int64, string) {
	return s.timeOp.ParseDuration(timeDuration)
}

func (s *IdGenerator) populateId(id *core.Id) {
	// 填充ID
	s.filler.PopulateId(id)
}

func (s *IdGenerator) SetIdPopulator(idPopulator core.IFiller) {
	// 设置填充器
	s.filler = idPopulator
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
