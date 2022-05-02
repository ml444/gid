package core

type Populater interface {
	PopulateId(id Ider, idMeta Metaer)
}

type ResetPopulater interface {
	Reset()
}
