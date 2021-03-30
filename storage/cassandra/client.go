package cassandra

import (
	"github.com/saidamir98/goauth_service/storage/repo"

	"github.com/gocql/gocql"
	"github.com/saidamir98/goauth_service/modules/rest"
)

type clientRepo struct {
	db *gocql.Session
}

// NewClientRepo ...
func NewClientRepo(db *gocql.Session) repo.ClientStorageI {
	return &clientRepo{db: db}
}

func (r *clientRepo) Get(clientPlatformID, clientTypeID string) (res rest.ClientModel, err error) {
	stmt := `SELECT
	client_platform_id,
	client_type_id
	FROM client
	WHERE client_platform_id = ? AND client_type_id = ?`

	if err = r.db.Query(stmt, clientPlatformID, clientTypeID).Scan(
		&res.ClientPlatformID,
		&res.ClientTypeID,
	); err != nil {
		return res, err
	}

	return res, nil
}
