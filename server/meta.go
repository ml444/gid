/*
ID类型分为最大峰值和最小粒度
【最大峰值】：
| 版本 | 类型 | 生成方式 | 秒级时间 | 序列号 | 机器ID |
| 63   | 62  | 60-61  | 59-30   | 29-10 | 0-9   |

【最小粒度】：
| 版本 | 类型 | 生成方式 | 秒级时间 | 序列号 | 机器ID |
| 63   | 62  | 60-61  | 59-20   | 19-10 | 0-9   |
*/
package server

import (
	"github.com/ml444/gid/config"
	"fmt"
)

type Meta struct {
	machineBits   byte
	seqBits       byte
	timeBits      byte
	genMethodBits byte
	typeBits      byte
	versionBits   byte
}

func (m *Meta) GetMachineBits() byte {
	return m.machineBits
}

func (m *Meta) SetMachineBits(machineBits byte) {
	m.machineBits = machineBits
}

func (m *Meta) GetMachineBitsMask() uint64 {
	// return -1L ^ -1L << machineBits
	res := -1 ^ (-1 << m.machineBits)
	return uint64(res)
}

func (m *Meta) GetSeqBitsStartPos() byte {
	return m.machineBits
}

func (m *Meta) GetSeqBits() byte {
	return m.seqBits
}

func (m *Meta) SetSeqBits(seqBits byte) {
	m.seqBits = seqBits
}

func (m *Meta) GetSeqBitsMask() uint64 {
	// return -1L ^ -1L << seqBits
	res := -1 ^ (-1 << m.seqBits)
	return uint64(res)
}

func (m *Meta) GetTimeBitsStartPos() byte {
	return m.machineBits + m.seqBits
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

func (m *Meta) GetGenMethodBitsStartPos() byte {
	return m.machineBits + m.seqBits + m.timeBits
}

func (m *Meta) GetGenMethodBits() byte {
	return m.genMethodBits
}

func (m *Meta) SetGenMethodBits(genMethodBits byte) {
	m.genMethodBits = genMethodBits
}

func (m *Meta) GetGenMethodBitsMask() uint64 {
	// return -1L ^ -1L << genMethodBits
	res := -1 ^ (-1 << m.genMethodBits)
	return uint64(res)
}

func (m *Meta) GetMtypeBitsStartPos() byte {
	// 10+10+30+2
	// 10+20+20+2
	return m.machineBits + m.seqBits + m.timeBits + m.genMethodBits
}

func (m *Meta) GetMtypeBits() byte {
	return m.typeBits
}

func (m *Meta) SetMtypeBits(typeBits byte) {
	m.typeBits = typeBits
}

func (m *Meta) GetMtypeBitsMask() uint64 {
	// return -1L ^ -1L << mtypeBits
	res := -1 ^ (-1 << m.typeBits)
	return uint64(res)
}

func (m *Meta) GetVersionBitsStartPos() byte {
	return (m.machineBits + m.seqBits + m.timeBits + m.genMethodBits + m.typeBits)
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
	if idType == config.MaxPeak {
		fmt.Println("选择最大峰值模式")
		return &Meta{
			machineBits:   10,
			seqBits:       20,
			timeBits:      30,
			genMethodBits: 2,
			typeBits:      1,
			versionBits:   1,
			// 10, 20, 30, 2, 1, 1
		}
	} else if idType == config.MinGranularity {
		fmt.Println("选择最小粒度模式")
		return &Meta{10, 10, 40, 2, 1, 1}
	} else {
		return &Meta{}
	}
}

// func init()  {

// }
