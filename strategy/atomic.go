package strategy

import (
	"sync/atomic"

	"github.com/ml444/gid/core"
)

var _ core.IStrategy = &AtomicStrategy{}

type Variant struct {
	sequence      uint64
	lastTimestamp uint64
}

type AtomicStrategy struct {
	variant         atomic.Value
	sequenceBitMask uint64
	timeOp          core.ITimeOp
}

func (p *AtomicStrategy) Caught() (timeDuration, sequence uint64) {
	var varOld Variant
	var varNew Variant

	for {
		// 取得并保存原来的变量，这个变量包含原来的时间和序号字段
		varOld = p.variant.Load().(Variant)

		// 基于原来的变量计算新的时间和序列号
		timeDuration = p.timeOp.TimeNow()

		p.timeOp.ValidateTimestamp(varOld.lastTimestamp, timeDuration) // 校验时间是否被回调变慢了

		sequence = varOld.sequence

		if timeDuration == varOld.lastTimestamp {
			sequence++
			sequence &= p.sequenceBitMask
			if sequence == 0 {
				timeDuration = p.timeOp.WaitNextTime(varOld.lastTimestamp)
			}
		} else {
			sequence = 0
		}
		// 使用CAS操作更新原来的变量，在更新的过程中需要传递保存原来的变量
		// Assign the current variant by the atomic tools
		varNew.sequence = sequence
		varNew.lastTimestamp = timeDuration

		// 如果保存的原来的变量被其他线程改变了，就需要拿到最新的变量，并再次计算和尝试
		// 如果未被修改，则更新变量，跳出循环
		if p.variant.CompareAndSwap(varOld, varNew) {
			return timeDuration, sequence
		}
	}
}
func (p *AtomicStrategy) Reset() {
	v := Variant{
		sequence:      0,
		lastTimestamp: 0,
	}
	p.variant.Store(v)
}

func NewAtomicFiller(sequenceBitMask uint64, timeOp core.ITimeOp) *AtomicStrategy {
	f := AtomicStrategy{variant: atomic.Value{}, sequenceBitMask: sequenceBitMask, timeOp: timeOp}
	f.Reset()
	return &f
}
