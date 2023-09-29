package strategy

import (
	"sync"

	"github.com/ml444/gid/core"
)

var _ core.IStrategy = &LockStrategy{}

type LockStrategy struct {
	sequence        uint64 //= 0
	lastTimestamp   uint64 //= -1
	mx              sync.RWMutex
	sequenceBitMask uint64
	timeOp          core.ITimeOp
}

func (p *LockStrategy) get() (timeDuration, sequence uint64) {
	timeDuration = p.timeOp.TimeNow()
	isValid := p.timeOp.ValidateTimestamp(p.lastTimestamp, timeDuration)
	if isValid != true {
		return 0, 0
	}

	// 查看当前时间是否已经到了下一个时间单位了，
	// 如果到了，则将序号清零。
	// 如果还在上一个时间单位，就对序列号进行累加。
	// 如果累加后越界了，就需要等待下一单位时间再产生唯一ID
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
func (p *LockStrategy) Caught() (uint64, uint64) {
	p.mx.RLock()
	defer p.mx.RUnlock()
	return p.get()
}

func (p *LockStrategy) Reset() {
	p.mx.Lock()
	defer p.mx.Unlock()
	p.sequence = 0
	p.lastTimestamp = 0
}

func NewLockFiller(sequenceBitMask uint64, timeOp core.ITimeOp) *LockStrategy {
	return &LockStrategy{
		sequence:        0,
		lastTimestamp:   0,
		mx:              sync.RWMutex{},
		sequenceBitMask: sequenceBitMask,
		timeOp:          timeOp,
	}
}
