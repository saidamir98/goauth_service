package cassandra

import (
	"errors"

	"github.com/saidamir98/goauth_service/storage/repo"

	"github.com/saidamir98/goauth_service/pkg/util"

	"github.com/gocql/gocql"
	"github.com/saidamir98/goauth_service/modules/rest"
)

type authRepo struct {
	db *gocql.Session
}

// NewAuthRepo ...
func NewAuthRepo(db *gocql.Session) repo.AuthStorageI {
	return &authRepo{db: db}
}

func (r *authRepo) GetUserIDByUsername(username string) (userID string, err error) {
	var query string
	if util.IsValidEmail(username) {
		query = "SELECT user_id FROM user_email WHERE email = ?"
	} else if util.IsValidPhone(username) {
		query = "SELECT user_id FROM user_phone WHERE phone = ?"
	} else if util.IsValidLogin(username) {
		query = "SELECT user_id FROM user_login WHERE login = ?"
	} else {
		return "", errors.New("username is not valid")
	}

	if err = r.db.Query(query, username).Scan(
		&userID,
	); err != nil {
		return "", err
	}

	return userID, nil
}

func (r *authRepo) GetUserByID(id string) (res rest.UserModel, err error) {
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

func (r *authRepo) GetClient(clientPlatformID, clientTypeID string) (res rest.ClientModel, err error) {
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

func (r *authRepo) CreateSession(entity rest.SessionModel) (err error) {
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

func (r *authRepo) GetSession(clientPlatformID, clientTypeID, userID, id string) (res rest.SessionModel, err error) {
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

func (r *authRepo) DeleteSession(clientPlatformID, clientTypeID, userID, id string) (err error) {
	stmt := `DELETE FROM session 
	WHERE client_platform_id = ? AND client_type_id = ? AND user_id = ? AND id = ?`

	if err = r.db.Query(stmt, clientPlatformID, clientTypeID, userID, id).Exec(); err != nil {
		return err
	}

	return nil
}

func (r *authRepo) GetSessionsByUserID(userID string) (res []rest.SessionModel, err error) {
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
