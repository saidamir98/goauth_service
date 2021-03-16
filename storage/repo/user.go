package repo

import "github.com/saidamir98/goauth/modules/rest"

// UserStorageI ...
type UserStorageI interface {
	RegisterUser(entity rest.RegisterUserModel) (userID string, err error)
}
