package strategy

import (
	"github.com/ml444/gid/core"
)

var _ core.IStrategy = &SyncStrategy{}

type SyncStrategy struct {
	sequence        uint64 //= 0
	lastTimestamp   uint64 //= -1
	sequenceBitMask uint64
	timeOp          core.ITimeOp
}

func (p *SyncStrategy) get() (timeDuration, sequence uint64) {
	timeDuration = p.timeOp.TimeNow()
	isValid := p.timeOp.ValidateTimestamp(p.lastTimestamp, timeDuration)
	if isValid != true {
		return
	}

	// 查看当前时间是否已经到了下一个时间单位了，
	// 如果到了，则将序号清零。
	// 如果还在上一个时间单位，就对序列号进行累加。
	// 如果累加后越界了，就需要等待下一时间单位再产生唯一ID
	if timeDuration == p.lastTimestamp {
		p.sequence++
		p.sequence = p.sequence & p.sequenceBitMask
		if p.sequence == 0 {
			timeDuration = p.timeOp.WaitNextTime(p.lastTimestamp)
		}
	} else {
		p.lastTimestamp = timeDuration
		p.sequence = 0
	}
	return p.lastTimestamp, p.sequence
}

func (p *SyncStrategy) Caught() (timeDuration, sequence uint64) {
	return p.get()
}
func (p *SyncStrategy) Reset() {
	p.sequence = 0
	p.lastTimestamp = 0
}
func NewSyncFiller(sequenceBitMask uint64, timeOp core.ITimeOp) *SyncStrategy {
	return &SyncStrategy{
		sequence:        0,
		lastTimestamp:   0,
		sequenceBitMask: sequenceBitMask,
		timeOp:          timeOp,
	}
}
