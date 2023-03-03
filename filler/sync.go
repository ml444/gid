package filler

import (
	"github.com/ml444/gid/core"
)

//var _ core.IFiller = &SyncFiller{}

type SyncFiller struct {
	sequence      uint64 //= 0
	lastTimestamp uint64 //= -1
	meta          core.IMeta
	timeOp        core.ITimeOp
}

func (p *SyncFiller) populateId(id core.IId) {
	timestamp := p.timeOp.TimeNow()
	isValid := p.timeOp.ValidateTimestamp(p.lastTimestamp, timestamp)
	if isValid != true {
		return
	}

	// 查看当前时间是否已经到了下一个时间单位了，
	// 如果到了，则将序号清零。
	// 如果还在上一个时间单位，就对序列号进行累加。
	// 如果累加后越界了，就需要等待下一时间单位再产生唯一ID
	if timestamp == p.lastTimestamp {
		// fmt.Println("没到下一秒，叠加！！！", p)
		p.sequence++
		p.sequence = p.sequence & p.meta.GetSequenceBitsMask()
		// fmt.Println("没到下一秒，叠加！！！", p.sequence)
		if p.sequence == 0 {
			timestamp = p.timeOp.WaitNextTime(p.lastTimestamp)
		}
	} else {
		//fmt.Println("下一秒到了！！！")
		p.lastTimestamp = timestamp
		p.sequence = 0
	}
	id.SetSequence(p.sequence)
	id.SetTime(p.lastTimestamp)
}

func (p *SyncFiller) PopulateId(id core.IId) {
	p.populateId(id)
}
func (p *SyncFiller) Reset() {
	p.sequence = 0
	p.lastTimestamp = 0
}
func NewSyncFiller(meta core.IMeta, timeOp core.ITimeOp) *SyncFiller {
	return &SyncFiller{
		sequence:      0,
		lastTimestamp: 0,
		meta:          meta,
		timeOp:        timeOp,
	}
}
