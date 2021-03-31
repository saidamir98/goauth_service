package repo

import "github.com/saidamir98/goauth_service/modules/rest"

// PasscodeStorageI ...
type PasscodeStorageI interface {
	Create(entity rest.PasscodeModel) (err error)
	Get(userID, id string) (res rest.PasscodeModel, err error)
	// Update(entity rest.PasscodeModel) (err error)
	// Delete(userID, id string) (err error)
	// GetByUserID(userID string) (res []rest.PasscodeModel, err error)
}
