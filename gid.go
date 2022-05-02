package gid

import (
	"encoding/base64"
	"fmt"
	"github.com/ml444/gid/core"
	"github.com/ml444/gid/populator"
	"github.com/ml444/gid/utils"
	"strconv"
	"sync"
)

type IdServer interface {
	GenerateId() uint64
	ExplainId(id uint64) *Id
	MakeId(timestamp int64, seq int64) uint64
	TransTime(time uint64) (int64, string)
}

const (
	SYNC   = "sync"
	ATOMIC = "atomic"
)

type IdService struct {
	DeviceId uint64
	Method   uint64
	Type     uint64
	Version  uint64

	meta      *Meta
	populator core.Populater
	convertor core.Converter

	idPool sync.Pool
}

func NewIdService(cfg Config) *IdService {
	s := IdService{
		DeviceId: cfg.DeviceId,
		Method:   cfg.Method,
		Type:     cfg.Type,
		Version:  cfg.Version,
	}
	s.meta = NewIdMeta(cfg.Type)

	switch cfg.PopulatorType {
	case PopulatorTypeSync:
		s.populator = populator.NewSyncIdPopulator()
	case PopulatorTypeAtomic:
		s.populator = populator.NewAtomicPopulator()
	case PopulatorTypeLock:
		s.populator = populator.NewLockPopulator()
	default:
		panic(fmt.Sprintf("invalid populator type: %d", cfg.PopulatorType))
	}

	s.convertor = NewConvertor(s.meta)
	s.idPool.New = func() interface{} {
		return NewId(s.Version, s.Type, s.Method, s.DeviceId)
	}
	return &s
}

// GenerateId: generate globally unique id.
func (s *IdService) GenerateId() uint64 {
	id := s.idPool.Get().(*Id)
	defer s.idPool.Put(id)
	s.populateId(id) // 这是一个抽象方法，调用子类的
	return s.convertor.ConvertToGen(id)
}

// ExplainId: parse the components of the unique id.
func (s *IdService) ExplainId(id uint64, out core.Ider) {
	s.convertor.ConvertToExp(id, out)
}

// MakeId: According to the incoming timestamp and sequence parameters,
// manually synthesize the unique id.
func (s *IdService) MakeId(timestamp int64, sequence uint64) (uint64, error) {
	timeDuration, err := utils.TransDuration(timestamp)
	if err != nil {
		return 0, err
	}

	id := s.idPool.Get().(*Id)
	defer s.idPool.Put(id)
	id.SetSequence(sequence)
	id.SetTime(timeDuration)
	return s.convertor.ConvertToGen(id), nil
}

// TransTime 转换时间
func (s *IdService) TransTime(timeDuration uint64) (int64, string) {
	return utils.TransTime(timeDuration, s.Type)
}

func (s *IdService) populateId(id *Id) {
	// 填充ID
	s.populator.PopulateId(id, s.meta)
}

func (s *IdService) SetIdPopulator(idPopulator core.Populater) {
	// 设置填充器
	s.populator = idPopulator
}




type Id struct {
	serialVersionUID uint64 //6870931236218221183L

	DeviceId uint64 `json:"device_id"`
	Sequence uint64 `json:"sequence"`
	Time     uint64 `json:"time_duration"`
	Method   uint64 `json:"method"`
	Type     uint64 `json:"type"`
	Version  uint64 `json:"Version"`
}

func (id *Id) GetSequence() uint64 {
	return id.Sequence
}

func (id *Id) SetSequence(seq uint64) {
	id.Sequence = seq
}

func (id *Id) GetTime() uint64 {
	return id.Time
}

func (id *Id) SetTime(time uint64) {
	id.Time = time
}

func (id *Id) GetDevice() uint64 {
	return id.DeviceId
}

func (id *Id) SetDevice(deviceId uint64) {
	id.DeviceId = deviceId
}

func (id *Id) GetMethod() uint64 {
	return id.Method
}

func (id *Id) SetMethod(genMethod uint64) {
	id.Method = genMethod
}
func (id *Id) GetType() uint64 {
	return id.Type
}

func (id *Id) SetType(typ uint64) {
	id.Type = typ
}
func (id *Id) GetVersion() uint64 {
	return id.Version
}

func (id *Id) SetVersion(version uint64) {
	id.Version = version
}

func NewId(version, typ, method, deviceId uint64) *Id {
	return &Id{
		// serialVersionUID: 6870931236218221183L,
		// Sequence:		uint64
		// Time     uint64
		DeviceId: deviceId,
		Method:   method,
		Type:     typ,
		Version:  version,
	}
}

type LongId uint64

func (f LongId) String() string {
	return strconv.FormatUint(uint64(f), 10)
}

func (f LongId) ToBinary() string {
	return strconv.FormatUint(uint64(f), 10)
}

func (f LongId) ToBase64() string {
	return base64.StdEncoding.EncodeToString([]byte(f.String()))
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

// A JSONSyntaxError is returned from UnmarshalJSON if an invalid ID is provided.
type JSONSyntaxError struct{ original []byte }

func (j JSONSyntaxError) Error() string {
	return fmt.Sprintf("invalid snowflake ID %q", string(j.original))
}

// MarshalJSON returns a json byte array string of the snowflake ID.
func (f LongId) MarshalJSON() ([]byte, error) {
	buff := make([]byte, 0, 22)
	buff = append(buff, '"')
	buff = strconv.AppendUint(buff, uint64(f), 10)
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
