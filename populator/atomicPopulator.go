package populator

import (
	"github.com/ml444/gid/server"
	"github.com/ml444/gid/utils"
	"sync/atomic"
)

type Variant struct {
	sequence     uint64
	lastTimestamp uint64
}

type AtomicPopulator struct {
	variant atomic.Value
}

func (p *AtomicPopulator) PopulateId(id *server.Id, idMeta *server.Meta) {
	var varOld Variant
	var varNew Variant
	var timestamp uint64
	var sequence uint64

	for {
		// 取得并保存原来的变量，这个变量包含原来的时间和序号字段
		varOld = p.variant.Load().(Variant)

		// 基于原来的变量计算新的时间和序列号
		timestamp = utils.GenTime(id.GetType())

		utils.ValidateTimestamp(varOld.lastTimestamp, timestamp) // 校验时间是否被回调变慢了

		sequence = varOld.sequence

		if timestamp == varOld.lastTimestamp {
			sequence++
			sequence &= idMeta.GetSeqBitsMask()
			if sequence == 0 {
				// 使用自旋锁
				timestamp = utils.TillNextTimeUnit(varOld.lastTimestamp, id.GetType())
			}
		} else {
			sequence = 0
		}
		// 使用CAS操作更新原来的变量，在更新的过程中需要传递保存原来的变量
		// Assign the current variant by the atomic tools
		varNew.sequence = sequence
		varNew.lastTimestamp = timestamp

		// 如果保存的原来的变量被其他线程改变了，就需要拿到最新的变量，并再次计算和尝试
		// 如果未被修改，则更新变量，跳出循环
		if p.variant.CompareAndSwap(varOld, varNew) {
			id.SetSeq(sequence)
			id.SetTime(timestamp)
			break
		}
	}
}

func (p *AtomicPopulator) Reset() {
	v := Variant{
		sequence:      0,
		lastTimestamp: 0,
	}
	p.variant.Store(v)
}
