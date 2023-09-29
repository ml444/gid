package core

type IMachineIdProvider interface {
	SetMachineId(machineId uint64)
	GetMachineId() uint64
}

//type IMachineIdsProvider interface {
//	SetMachineIds(machineIds []uint64)
//	GetMachineIds() []uint64
//	GetNextMachineId() uint64
//}
