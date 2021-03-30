package repo

import "github.com/saidamir98/goauth_service/modules/rest"

// ClientStorageI ...
type ClientStorageI interface {
	Get(clientPlatformID, clientTypeID string) (res rest.ClientModel, err error)
}
