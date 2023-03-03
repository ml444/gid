package core

type IFiller interface {
	PopulateId(id IId)
}

type IResetFiller interface {
	Reset()
}
