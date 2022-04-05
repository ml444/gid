package server

import (
	"encoding/base64"
	"fmt"
	"strconv"
)

type Id struct {
	serialVersionUID uint64 //6870931236218221183L

	Machine uint64 `json:"machine_id"`
	Seq     uint64 `json:"sequence"`
	Time    uint64 `json:"time_duration"`
	Method  uint64 `json:"method"`
	Type    uint64 `json:"type"`
	Version uint64 `json:"Version"`
}

func (id *Id) GetSeq() uint64 {
	return id.Seq
}

func (id *Id) SetSeq(seq uint64) {
	id.Seq = seq
}

func (id *Id) GetTime() uint64 {
	return id.Time
}

func (id *Id) SetTime(time uint64) {
	id.Time = time
}

func (id *Id) GetMachine() uint64 {
	return id.Machine
}

func (id *Id) SetMachine(machine uint64) {
	id.Machine = machine
}

func (id *Id) GetGenMethod() uint64 {
	return id.Method
}

func (id *Id) SetGenMethod(genMethod uint64) {
	id.Method = genMethod
}
func (id *Id) GetType() uint64 {
	return id.Type
}

func (id *Id) SetType(mtype uint64) {
	id.Type = mtype
}
func (id *Id) GetVersion() uint64 {
	return id.Version
}

func (id *Id) SetVersion(version uint64) {
	id.Version = version
}

func NewId(machine uint64, genMethod uint64, mtype uint64, version uint64) *Id {
	return &Id{
		// serialVersionUID: 6870931236218221183L,
		// Seq:		uint64
		// Time     uint64
		Machine: machine,
		Method:  genMethod,
		Type:    mtype,
		Version: version,
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
