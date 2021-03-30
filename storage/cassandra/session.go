package cassandra

import (
	"github.com/saidamir98/goauth_service/storage/repo"

	"github.com/gocql/gocql"
	"github.com/saidamir98/goauth_service/modules/rest"
)

type sessionRepo struct {
	db *gocql.Session
}

// NewSessionRepo ...
func NewSessionRepo(db *gocql.Session) repo.SessionStorageI {
	return &sessionRepo{db: db}
}

func (r *sessionRepo) Create(entity rest.SessionModel) (err error) {
	qry := r.db.Query(`INSERT INTO session(
		client_platform_id,
		client_type_id,
		user_id,
		id,
		role_id,
		ip,
		data, 
		created_at,
		updated_at,
		expires_at) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?) USING TTL ?`,
		entity.ClientPlatformID,
		entity.ClientTypeID,
		entity.UserID,
		entity.ID,
		entity.RoleID,
		entity.IP,
		entity.Data,
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

func (r *sessionRepo) Get(clientPlatformID, clientTypeID, userID, id string) (res rest.SessionModel, err error) {
	stmt := `SELECT
	client_platform_id,
	client_type_id,
	user_id,
	id,
	role_id,
	ip,
	data, 
	created_at,
	updated_at,
	expires_at
	FROM session
	WHERE client_platform_id = ? AND client_type_id = ? AND user_id = ? AND id = ?`

	if err = r.db.Query(stmt, clientPlatformID, clientTypeID, userID, id).Scan(
		&res.ClientPlatformID,
		&res.ClientTypeID,
		&res.UserID,
		&res.ID,
		&res.RoleID,
		&res.IP,
		&res.Data,
		&res.CreatedAt,
		&res.UpdatedAt,
		&res.ExpiresAt,
	); err != nil {
		return res, err
	}

	return res, nil
}

func (r *sessionRepo) Delete(clientPlatformID, clientTypeID, userID, id string) (err error) {
	stmt := `DELETE FROM session 
	WHERE client_platform_id = ? AND client_type_id = ? AND user_id = ? AND id = ?`

	if err = r.db.Query(stmt, clientPlatformID, clientTypeID, userID, id).Exec(); err != nil {
		return err
	}

	return nil
}

func (r *sessionRepo) GetByUserID(userID string) (res []rest.SessionModel, err error) {
	stmt := `SELECT
	client_platform_id,
	client_type_id,
	user_id,
	id,
	role_id,
	ip,
	data, 
	created_at,
	updated_at,
	expires_at
	FROM session
	WHERE user_id = ?`

	scanner := r.db.Query(stmt, userID).Iter().Scanner()

	for scanner.Next() {
		var session rest.SessionModel

		if err = scanner.Scan(
			&session.ClientPlatformID,
			&session.ClientTypeID,
			&session.UserID,
			&session.ID,
			&session.RoleID,
			&session.IP,
			&session.Data,
			&session.CreatedAt,
			&session.UpdatedAt,
			&session.ExpiresAt,
		); err != nil {
			return res, err
		}

		res = append(res, session)
	}

	return res, nil
}
