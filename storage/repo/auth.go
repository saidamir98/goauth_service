package repo

import "github.com/saidamir98/goauth_service/modules/rest"

// AuthStorageI ...
type AuthStorageI interface {
	GetUserIDByUsername(username string) (userID string, err error)
	RegisterUser(entity rest.RegisterUserModel) (userID string, err error)
}
