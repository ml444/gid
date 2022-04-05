package gid

import (
	"fmt"
	"github.com/ml444/gid/convert"
	"github.com/ml444/gid/populator"
	"github.com/ml444/gid/server"
	"github.com/ml444/gid/utils"
)

type IdServer interface {
	GenId() uint64              // 主要API，用来产生唯一ID
	ExpId(id uint64) *server.Id // 解读ID
	// MakeID(time int64, seq int64)
	// MakeID(time int64, seq int64, machine string)
	// MakeID(time int64, seq int64, machine string, genMethod string)
	MakeId(machine, seq, time, method, mtype, version uint64) uint64 // 伪造某一时间的ID
	TransTime(time uint64) (int64, string)                           // 将整型时间翻译成格式化时间
}

const (
	SYNC   = "sync"
	ATOMIC = "atomic"
)

type IdService struct {
	MachineId   uint64
	GenMethod   uint64
	Type        uint64
	Version     uint64
	idMeta      *server.Meta
	idPopulator populator.Populater
}

func NewIdService(machineId, method, idType, version uint64, populatorType string) *IdService {
	self := IdService{
		MachineId: machineId,
		GenMethod: method,
		Type:      idType,
		Version:   version,
	}
	// 初始化IdMeta
	self.idMeta = server.NewIdMeta(idType)

	if self.idPopulator != nil {
		fmt.Println("The idPopulator is used.")
	} else if populatorType == SYNC {
		fmt.Println("The SyncIdPopulator is used.")
		self.idPopulator = populator.NewSyncIdPopulator()
		// ResetPopulater 接口查询
		//if Populator, ok := self.idPopulator.(populator.ResetPopulater); ok {
		//	fmt.Printf("===>%T 实现了reset接口！！\n", Populator)
		//}
	} else if populatorType == ATOMIC {
		fmt.Println("The AtomicPopulator is used.")
		// self.idPopulator = new(populator.AtomicPopulator)
	} else {
		fmt.Println("The default LockPopulator is used.")
		// self.idPopulator = new(populator.LockPopulator)
	}
	// populator.PrintType(self.idPopulator)
	return &self
}
// 产生唯一ID
func (s *IdService) GenId() uint64 {
	id := server.NewId(s.MachineId, s.GenMethod, s.Type, s.Version)
	s.populateId(id) // 这是一个抽象方法，调用子类的
	return convert.NewConvertor(s.idMeta).ConvertToGen(id)
}

// 解读ID
func (s *IdService) ExpId(id uint64) *server.Id {
	convertor := convert.NewConvertor(s.idMeta)
	return convertor.ConvertToExp(id)
}

// MakeId 手动合成ID
func (s *IdService) MakeId(machine, sequence, timestamp, method, mtype, version uint64) uint64 {
	// 把时间戳转为项目的时间量
	timeDuration, err := utils.TransDuration(timestamp)
	fmt.Println(timeDuration, timestamp)
	if err != nil {
		fmt.Println(err)
	}
	// 实例化Id对象
	id := server.Id{}
	id.SetMachine(machine)
	id.SetSeq(sequence)
	id.SetTime(timeDuration)
	id.SetGenMethod(method)
	id.SetType(mtype)
	id.SetVersion(version)

	convertor := convert.NewConvertor(s.idMeta)
	return convertor.ConvertToGen(&id)
}

// TransTime 转换时间
func (s *IdService) TransTime(timeDuration uint64) (int64, string) {
	return utils.TransTime(timeDuration, s.Type)
}

func (s *IdService) populateId(id *server.Id) {
	// 填充ID
	s.idPopulator.PopulateId(id, s.idMeta)
}

func (s *IdService) SetIdPopulator(idPopulator populator.Populater) {
	// 设置填充器
	s.idPopulator = idPopulator
}

