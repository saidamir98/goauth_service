package repo

import "github.com/saidamir98/goauth/modules/rest"

// ClientTypeStorageI ...
type ClientTypeStorageI interface {
	GetByID(id string) (res rest.ClientTypeModel, err error)
}
