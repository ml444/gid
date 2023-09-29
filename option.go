package gid

import (
	"github.com/ml444/gid/core"
	"github.com/ml444/gid/strategy"
)

type StrategyType uint8

const (
	StrategyTypeSync   StrategyType = 1
	StrategyTypeAtomic StrategyType = 2
	StrategyTypeLock   StrategyType = 3
)

type OptionFunc func(*IdGenerator)

func SetStrategyByTypeOption(strategyType StrategyType) OptionFunc {
	return func(ig *IdGenerator) {
		switch strategyType {
		case StrategyTypeSync:
			ig.strategy = strategy.NewSyncFiller(ig.meta.GetBitMask(ig.seqIdx), ig.timeOp)
		case StrategyTypeAtomic:
			ig.strategy = strategy.NewAtomicFiller(ig.meta.GetBitMask(ig.seqIdx), ig.timeOp)
		case StrategyTypeLock:
			ig.strategy = strategy.NewLockFiller(ig.meta.GetBitMask(ig.seqIdx), ig.timeOp)
		}
	}
}
func SetStrategyOption(strategy core.IStrategy) OptionFunc {
	return func(ig *IdGenerator) {
		ig.strategy = strategy
	}
}
