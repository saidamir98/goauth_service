package repo

import "github.com/saidamir98/goauth/modules/rest"

// AuthStorageI ...
type AuthStorageI interface {
	GetUserIDByUsername(username string) (userID string, err error)
	GetUserByID(id string) (res rest.UserModel, err error)
	GetClient(clientPlatformID, clientTypeID string) (res rest.ClientModel, err error)
	CreateSession(entity rest.SessionModel) (err error)
	GetSession(clientPlatformID, clientTypeID, userID, id string) (res rest.SessionModel, err error)
	DeleteSession(clientPlatformID, clientTypeID, userID, id string) (err error)
}
