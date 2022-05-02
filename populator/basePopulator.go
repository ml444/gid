package populator

import (
	"fmt"
	"github.com/ml444/gid/core"
	"github.com/ml444/gid/utils"
	"os"
)

type basePopulator struct {
	sequence      uint64 //= 0
	lastTimestamp uint64 //= -1
}

func (p *basePopulator) populateId(id core.Ider, idMeta core.Metaer) {
	timestamp := utils.GenTime(id.GetType())
	isValid := utils.ValidateTimestamp(p.lastTimestamp, timestamp)
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
		p.sequence = p.sequence & idMeta.GetSequenceBitsMask()
		// fmt.Println("没到下一秒，叠加！！！", p.sequence)
		if p.sequence == 0 {
			timestamp = utils.TillNextTimeUnit(p.lastTimestamp, id.GetType())
		}
	} else {
		//fmt.Println("下一秒到了！！！")
		p.lastTimestamp = timestamp
		p.sequence = 0
	}
	id.SetSequence(p.sequence)
	id.SetTime(p.lastTimestamp)
}

func (p *basePopulator) PopulateId(id core.Ider, idMeta core.Metaer) {
	p.populateId(id, idMeta)
}
func (p *basePopulator) Reset() {
	p.sequence = 0
	p.lastTimestamp = 0
}

// func ValidatePopulator(p_intf Populater) {
// 	// 查询接口类型是否是populator类型
// 	// if p, ok := p_intf.(*SyncIdPopulator); ok {
// 	p := reflect.ValueOf(p_intf)
// 	if p == basePopulator {
// 		fmt.Println("===>该接口类型是basepopulator的子类型")
// 		fmt.Printf("%T \n", p)
// 	} else {
// 		fmt.Println("===>该接口类型不是basepopulator的子类型")
// 		fmt.Printf("%T \n", p)
// 	}
// }
func PrintType(args ...interface{}) {
	for i, arg := range args {
		if i > 0 {
			os.Stdout.WriteString("")
		}
		switch t := arg.(type) {
		case core.Populater:
			os.Stdout.WriteString("Populater")
		case SyncIdPopulator:
			os.Stdout.WriteString("SyncIdPopulator")
			// case LockPopulator:
			// 	os.Stdout.WriteString(t.name)
			// case AtomicPopulator:
			// 	os.Stdout.WriteString(t.name)
		default:
			os.Stdout.WriteString("???")
			fmt.Println(t)

		}
	}
}
