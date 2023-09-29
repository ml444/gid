package gid

import "testing"

const EPOCH = int64(1610351662000) //起始时间，预计可用34+34年

func BenchmarkStrategyAtomic(b *testing.B) {
	ig, err := NewSnowflakeIdGenerator(EPOCH, 1, 0, SetStrategyByTypeOption(StrategyTypeAtomic))
	if err != nil {
		b.Error(err.Error())
	}
	for i := 0; i < b.N; i++ {
		ig.GenerateId()
	}
}

func BenchmarkStrategySync(b *testing.B) {
	ig, err := NewSnowflakeIdGenerator(EPOCH, 1, 0, SetStrategyByTypeOption(StrategyTypeSync))
	if err != nil {
		b.Error(err.Error())
	}
	for i := 0; i < b.N; i++ {
		ig.GenerateId()
	}
}

func BenchmarkStrategyLock(b *testing.B) {
	ig, err := NewSnowflakeIdGenerator(EPOCH, 1, 0, SetStrategyByTypeOption(StrategyTypeLock))
	if err != nil {
		b.Error(err.Error())
	}
	for i := 0; i < b.N; i++ {
		ig.GenerateId()
	}
}
