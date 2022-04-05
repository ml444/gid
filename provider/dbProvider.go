package provider

type DbMachineIdProvider struct {
	machineId uint64
	// job
}

func (p *DbMachineIdProvider) GetMachineId() uint64 {
	return p.machineId
}

func (p *DbMachineIdProvider) SetMachineId(machineId uint64) {
	p.machineId = machineId
}
