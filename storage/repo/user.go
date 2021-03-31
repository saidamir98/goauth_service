package repo

import "github.com/saidamir98/goauth_service/modules/rest"

// UserStorageI ...
type UserStorageI interface {
	GetByID(id string) (res rest.UserModel, err error)
	Create(entity rest.CreateUserModel) (userID string, err error)
	Update(entity rest.UpdateUserModel) (err error)
}
