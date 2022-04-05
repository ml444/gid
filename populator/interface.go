package populator

import "github.com/ml444/gid/server"

type Populater interface {
	PopulateId(id *server.Id, idMeta *server.Meta)
}

type ResetPopulater interface {
	Reset()
}
