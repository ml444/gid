package filler

import (
	"github.com/ml444/gid/core"
	"sync"
)

//var _ core.IFiller = &LockFiller{}

type LockFiller struct {
	sequence      uint64 //= 0
	lastTimestamp uint64 //= -1
	mx            sync.RWMutex
	meta          core.IMeta
	timeOp        core.ITimeOp
}

func (p *LockFiller) populateId(id core.IId) {
	timestamp := p.timeOp.TimeNow()
	isValid := p.timeOp.ValidateTimestamp(p.lastTimestamp, timestamp)
	if isValid != true {
		return
	}

	// 查看当前时间是否已经到了下一个时间单位了，
	// 如果到了，则将序号清零。
	// 如果还在上一个时间单位，就对序列号进行累加。
	// 如果累加后越界了，就需要等待下一单位时间再产生唯一ID
	if timestamp == p.lastTimestamp {
		p.sequence++
		p.sequence = p.sequence & p.meta.GetSequenceBitsMask()
		if p.sequence == 0 {
			timestamp = p.timeOp.WaitNextTime(p.lastTimestamp)
		}
	} else {
		p.lastTimestamp = timestamp
		p.sequence = 0
	}
	id.SetSequence(p.sequence)
	id.SetTime(p.lastTimestamp)
}
func (p *LockFiller) PopulateId(id core.IId) {
	p.mx.RLock()
	defer p.mx.RUnlock()
	p.populateId(id)
}

func (p *LockFiller) Reset() {
	p.mx.Lock()
	defer p.mx.Unlock()
	p.sequence = 0
	p.lastTimestamp = 0
}

func NewLockFiller(meta core.IMeta, timeOp core.ITimeOp) *LockFiller {
	return &LockFiller{
		sequence:      0,
		lastTimestamp: 0,
		mx:            sync.RWMutex{},
		meta:          meta,
		timeOp:        timeOp,
	}
}
