package gid

import "github.com/ml444/gid/core"

type Config struct {
	Epoch       int64
	WorkID      uint64
	SegmentBits *SegmentBits
}

type SegmentBits = core.SegmentBits
type IDComponents = core.IDComponents
