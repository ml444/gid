/*
ID类型分为最大峰值和最小粒度
【最大峰值】：
| 版本 | 类型 | 生成方式 | 秒级时间 | 序列号 | 机器ID |
| 63   | 62  | 60-61  | 59-30   | 29-10 | 0-9   |

【最小粒度】：
| 版本 | 类型 | 生成方式 | 秒级时间 | 序列号 | 机器ID |
| 63   | 62  | 60-61  | 59-20   | 19-10 | 0-9   |
*/
package gid

import (
	"fmt"
)

type Meta struct {
	machineBits  byte
	sequenceBits byte
	timeBits     byte
	methodBits   byte
	typeBits     byte
	versionBits  byte
}

func (m *Meta) GetDeviceBits() byte {
	return m.machineBits
}

func (m *Meta) SetDeviceBits(deviceBits byte) {
	m.machineBits = deviceBits
}

func (m *Meta) GetDeviceBitsMask() uint64 {
	// return -1L ^ -1L << machineBits
	res := -1 ^ (-1 << m.machineBits)
	return uint64(res)
}

func (m *Meta) GetSequenceBitsStartPos() byte {
	return m.machineBits
}

func (m *Meta) GetSequenceBits() byte {
	return m.sequenceBits
}

func (m *Meta) SetSequenceBits(seqBits byte) {
	m.sequenceBits = seqBits
}

func (m *Meta) GetSequenceBitsMask() uint64 {
	// return -1L ^ -1L << sequenceBits
	res := -1 ^ (-1 << m.sequenceBits)
	return uint64(res)
}

func (m *Meta) GetTimeBitsStartPos() byte {
	return m.machineBits + m.sequenceBits
}

func (m *Meta) GetTimeBits() byte {
	return m.timeBits
}

func (m *Meta) SetTimeBits(timeBits byte) {
	m.timeBits = timeBits
}

func (m *Meta) GetTimeBitsMask() uint64 {
	// return -1L ^ -1L << timeBits
	res := -1 ^ (-1 << m.timeBits)
	return uint64(res)
}

func (m *Meta) GetMethodBitsStartPos() byte {
	return m.machineBits + m.sequenceBits + m.timeBits
}

func (m *Meta) GetMethodBits() byte {
	return m.methodBits
}

func (m *Meta) SetMethodBits(genMethodBits byte) {
	m.methodBits = genMethodBits
}

func (m *Meta) GetMethodBitsMask() uint64 {
	// return -1L ^ -1L << methodBits
	res := -1 ^ (-1 << m.methodBits)
	return uint64(res)
}

func (m *Meta) GetTypeBitsStartPos() byte {
	// 10+10+30+2
	// 10+20+20+2
	return m.machineBits + m.sequenceBits + m.timeBits + m.methodBits
}

func (m *Meta) GetTypeBits() byte {
	return m.typeBits
}

func (m *Meta) SetTypeBits(typeBits byte) {
	m.typeBits = typeBits
}

func (m *Meta) GetTypeBitsMask() uint64 {
	// return -1L ^ -1L << mtypeBits
	res := -1 ^ (-1 << m.typeBits)
	return uint64(res)
}

func (m *Meta) GetVersionBitsStartPos() byte {
	return (m.machineBits + m.sequenceBits + m.timeBits + m.methodBits + m.typeBits)
}

func (m *Meta) GetVersionBits() byte {
	return m.versionBits
}

func (m *Meta) SetVersionBits(versionBits byte) {
	m.versionBits = versionBits
}

func (m *Meta) GetVersionBitsMask() uint64 {
	// return -1L ^ -1L << versionBits
	res := -1 ^ (-1 << m.versionBits)
	return uint64(res)
}

// // 工厂类，决定返回的idMeta是最大峰值还是最小粒度
// type IdMetaFactory struct {
// 	maxPeak        Meta
// 	minGranularity Meta
// }

func NewIdMeta(idType uint64) *Meta {
	if idType == MaxPeak {
		fmt.Println("选择最大峰值模式")
		return &Meta{
			machineBits:  10,
			sequenceBits: 20,
			timeBits:     30,
			methodBits:   2,
			typeBits:     1,
			versionBits:  1,
			// 10, 20, 30, 2, 1, 1
		}
	} else if idType == MinGranularity {
		fmt.Println("选择最小粒度模式")
		return &Meta{10, 10, 40, 2, 1, 1}
	} else {
		return &Meta{}
	}
}

// func init()  {

// }
