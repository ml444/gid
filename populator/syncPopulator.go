package populator

type SyncIdPopulator struct {
	basePopulator
}


func NewSyncIdPopulator() *SyncIdPopulator {
	idPopulator := &SyncIdPopulator{basePopulator{0, 0}}
	return idPopulator
}
