package cassandra

import (
	"github.com/saidamir98/goauth_service/modules/rest"
	"github.com/saidamir98/goauth_service/storage/repo"

	"github.com/gocql/gocql"
)

type userRepo struct {
	db *gocql.Session
}

// NewUserRepo ...
func NewUserRepo(db *gocql.Session) repo.UserStorageI {
	return &userRepo{db: db}
}

func (r *userRepo) GetByID(id string) (res rest.UserModel, err error) {
	query := `SELECT
	id,
	client_type_id,
	role_id,
	password,
	active,
	expires_at,
	created_at,
	updated_at
	FROM user WHERE id = ?`

	if err = r.db.Query(query, id).Scan(
		&res.ID,
		&res.ClientTypeID,
		&res.RoleID,
		&res.Password,
		&res.Active,
		&res.ExpiresAt,
		&res.CreatedAt,
		&res.UpdatedAt,
	); err != nil {
		return res, err
	}

	return res, nil
}
