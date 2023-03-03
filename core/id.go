package core

type IId interface {
	SetVersion(version uint64)
	GetVersion() uint64
	SetType(typ uint64)
	GetType() uint64
	SetMethod(method uint64)
	GetMethod() uint64
	SetDevice(deviceId uint64)
	GetDevice() uint64
	SetTime(time uint64)
	GetTime() uint64
	SetSequence(sequence uint64)
	GetSequence() uint64
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
