package provider

type MachineIdProvider interface {
	SetMachineId(machineId uint64)
	GetMachineId() uint64
}

type MachineIdsProvider interface {
	SetMachineIds(machineIds []uint64)
	GetMachineIds() []uint64
	GetNextMachineId() uint64
}
