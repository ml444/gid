package populator

import (
	"github.com/ml444/gid/core"
	"sync"
)

type LockPopulator struct {
	basePopulator
	mx sync.RWMutex
}

func (p *LockPopulator) PopulateId(id core.Ider, idMeta core.Metaer) {
	p.mx.RLock()
	defer p.mx.RUnlock()
	p.populateId(id, idMeta)
}

func (p *LockPopulator) Reset() {
	p.mx.Lock()
	defer p.mx.Unlock()
	p.sequence = 0
	p.lastTimestamp = 0
}

func NewLockPopulator() *LockPopulator {
	return nil
}