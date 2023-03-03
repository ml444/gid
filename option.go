package gid

import "github.com/ml444/gid/core"

type IdType uint64
type FillerType uint8

const EPOCH = int64(1610351662000) //起始时间，预计可用34+34年
const (
	MinGranularity IdType = 0 // 最小颗粒度
	MaxPeak        IdType = 1 // 最大峰值

	FillerTypeSync   FillerType = 1
	FillerTypeAtomic FillerType = 2
	FillerTypeLock   FillerType = 3
)

type OptionFunc func(*IdGenerator)

func SetVersionOption(version uint64) OptionFunc {
	return func(ig *IdGenerator) {
		ig.Version = version
	}
}

func SetTypeOption(typ IdType, epoch int64) OptionFunc {
	return func(ig *IdGenerator) {
		if epoch <= 0 {
			panic("epoch must be gather than 0")
		}
		ig.Type = uint64(typ)
		switch typ {
		case MaxPeak:
			//fmt.Println("选择最大峰值模式")
			ig.meta = core.NewMetaWithMaxPeak()
			ig.timeOp = core.NewTimeOpWithMaxPeak(epoch)
		case MinGranularity:
			//fmt.Println("选择最小粒度模式")
			ig.meta = core.NewMetaWithMinGranularity()
			ig.timeOp = core.NewTimeOpWithMinGranularity(epoch)
		default:
			panic("other types are not supported yet")
		}
	}
}

func SetMethodOption(method uint64) OptionFunc {
	return func(ig *IdGenerator) {
		ig.Method = method
	}
}

func SetDeviceIdOption(di uint64) OptionFunc {
	return func(ig *IdGenerator) {
		ig.DeviceId = di
	}
}

func SetFillerTypeOption(fTyp FillerType) OptionFunc {
	return func(ig *IdGenerator) {
		ig.fillerType = fTyp
	}
}
func SetFillerOption(f core.IFiller) OptionFunc {
	return func(ig *IdGenerator) {
		ig.filler = f
	}
}

func SetConverterOption(c core.IConverter) OptionFunc {
	return func(ig *IdGenerator) {
		ig.convertor = c
	}
}
