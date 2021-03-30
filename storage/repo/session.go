package repo

import "github.com/saidamir98/goauth_service/modules/rest"

// SessionStorageI ...
type SessionStorageI interface {
	Create(entity rest.SessionModel) (err error)
	Get(clientPlatformID, clientTypeID, userID, id string) (res rest.SessionModel, err error)
	Delete(clientPlatformID, clientTypeID, userID, id string) (err error)
	GetByUserID(userID string) (res []rest.SessionModel, err error)
}
