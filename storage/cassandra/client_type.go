package cassandra

import (
	"github.com/saidamir98/goauth_service/storage/repo"

	"github.com/gocql/gocql"
	"github.com/saidamir98/goauth/modules/rest"
)

type clientTypeRepo struct {
	db *gocql.Session
}

// NewClientTypeRepo ...
func NewClientTypeRepo(db *gocql.Session) repo.ClientTypeStorageI {
	return &clientTypeRepo{db: db}
}

func (r *clientTypeRepo) GetByID(id string) (res rest.ClientTypeModel, err error) {
	query := `SELECT
	id,
	name,
	register_strategy,
	login_strategy,
	recovery_strategy
	FROM client_type
	WHERE id = ?`

	if err = r.db.Query(query, id).Scan(
		&res.ID,
		&res.Name,
		&res.RegisterStrategy,
		&res.LoginStrategy,
		&res.RecoveryStrategy,
	); err != nil {
		return res, err
	}

	return res, nil
}
