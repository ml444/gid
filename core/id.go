package core

type Ider interface {
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
