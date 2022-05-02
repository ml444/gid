package gid

const (
	MinGranularity = 0 // 最小颗粒度
	MaxPeak        = 1 // 最大峰值


	PopulatorTypeSync = 1
	PopulatorTypeAtomic  = 2
	PopulatorTypeLock  = 3

)

type Config struct {
	Version  uint64 `json:"version"`
	Type     uint64 `json:"type"`
	Method   uint64 `json:"method"`
	DeviceId uint64 `json:"device_id"`

	PopulatorType uint8 `json:"populator_type"`

}
