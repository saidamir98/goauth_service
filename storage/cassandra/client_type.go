package cassandra

import (
	"github.com/saidamir98/goauth_service/storage/repo"

	"github.com/gocql/gocql"
	"github.com/saidamir98/goauth_service/modules/rest"
)

type clientTypeRepo struct {
	db *gocql.Session
}

// NewClientTypeRepo ...
func NewClientTypeRepo(db *gocql.Session) repo.ClientTypeStorageI {
	return &clientTypeRepo{db: db}
}

func (r *clientTypeRepo) GetByID(id string) (res rest.ClientTypeModel, err error) {
	stmt := `SELECT
	id,
	name,
	login_strategy,
	confirm_by,
	self_register,
	self_recover
	FROM client_type
	WHERE id = ?`

	if err = r.db.Query(stmt, id).Scan(
		&res.ID,
		&res.Name,
		&res.LoginStrategy,
		&res.ConfirmBy,
		&res.SelfRegister,
		&res.SelfRecover,
	); err != nil {
		return res, err
	}

	return res, nil
}
