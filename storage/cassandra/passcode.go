package cassandra

import (
	"github.com/saidamir98/goauth_service/storage/repo"

	"github.com/gocql/gocql"
	"github.com/saidamir98/goauth_service/modules/rest"
)

type passcodeRepo struct {
	db *gocql.Session
}

// NewPasscodeRepo ...
func NewPasscodeRepo(db *gocql.Session) repo.PasscodeStorageI {
	return &passcodeRepo{db: db}
}

func (r *passcodeRepo) Create(entity rest.PasscodeModel) (err error) {
	qry := r.db.Query(`INSERT INTO passcode(
		user_id,
		id,
		hashed_code,
		state, 
		created_at,
		updated_at,
		expires_at) VALUES(?, ?, ?, ?, ?, ?, ?) USING TTL ?`,
		entity.UserID,
		entity.ID,
		entity.HashedCode,
		entity.State,
		entity.CreatedAt,
		entity.UpdatedAt,
		entity.ExpiresAt,
		int(entity.ExpiresAt.Sub(entity.CreatedAt).Seconds()),
	)

	if err := qry.Exec(); err != nil {
		return err
	}

	return nil
}

func (r *passcodeRepo) Get(userID, id string) (res rest.PasscodeModel, err error) {
	stmt := `SELECT
	user_id,
	id,
	hashed_code,
	state,
	created_at,
	updated_at,
	expires_at
	FROM passcode
	WHERE user_id = ? AND id = ?`

	if err = r.db.Query(stmt, userID, id).Scan(
		&res.UserID,
		&res.ID,
		&res.HashedCode,
		&res.State,
		&res.CreatedAt,
		&res.UpdatedAt,
		&res.ExpiresAt,
	); err != nil {
		return res, err
	}

	return res, nil
}
