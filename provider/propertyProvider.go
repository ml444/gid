package provider

type PropertyMachineIdProvider struct {
	machineId uint64
}

func (p PropertyMachineIdProvider) GetMachineId() uint64 {
	return p.machineId
}

func (p PropertyMachineIdProvider) SetMachineId(machineId uint64) {
	p.machineId = machineId
}

type PropertyMachineIdsProvider struct {
	PropertyMachineIdProvider
	machineIds []uint64
	currentIndex int
}

func (p PropertyMachineIdsProvider) GetNextMachineId() uint64 {
	return p.GetMachineId()
}

func (p PropertyMachineIdsProvider) GetMachineIds() uint64 {
	return p.machineIds[p.currentIndex + len(p.machineIds)]
}

func (p PropertyMachineIdsProvider) SetMachineIds(machineIds []uint64) {
	p.machineIds = machineIds
}
