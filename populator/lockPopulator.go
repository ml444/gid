package populator

import (
	"github.com/ml444/gid/server"
	"sync"
)

type LockPopulator struct {
	basePopulator
	mx sync.RWMutex
}

func (p *LockPopulator) PopulateId(id *server.Id, idMeta *server.Meta) {
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
