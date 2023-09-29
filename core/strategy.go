package core

type IStrategy interface {
	Caught() (timeDuration, sequence uint64)
}

type IResetFiller interface {
	Reset()
}
