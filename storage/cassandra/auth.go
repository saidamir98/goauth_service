package cassandra

import (
	"errors"
	"time"

	"github.com/google/uuid"
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

func (r *authRepo) RegisterUser(entity rest.RegisterUserModel) (userID string, err error) {
	uuid, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	entity.ID = uuid.String()

	batch := gocql.NewBatch(gocql.LoggedBatch)

	stmt := `SELECT
	client_type_id,
	name,
	id
	FROM role WHERE id = ? LIMIT 1`
	var (
		clientTypeID string
		roleName     string
		roleID       string
		c            int
	)
	if err = r.db.Query(stmt, entity.RoleID).Scan(
		&clientTypeID,
		&roleName,
		&roleID,
	); err != nil {
		return "", err
	}

	if clientTypeID != entity.ClientTypeID {
		return "", errors.New("mismatch between role_id and client_type_id")
	}

	stmtInsertUserPhone := `INSERT INTO user_phone
		(phone,
		user_id,
		created_at,
		updated_at)
	VALUES(?, ?, ?, ?)`
	for i := 0; i < len(entity.Phones); i++ {
		if err = r.db.Query(`SELECT count(user_id) FROM user_phone WHERE phone = ? LIMIT 1`, entity.Phones[i]).Scan(
			&c,
		); err != nil {
			return "", err
		}
		if c > 0 {
			return "", errors.New("phone is taken")
		}
		batch.Query(stmtInsertUserPhone, entity.Phones[i], entity.ID, time.Now(), time.Now())
	}

	if err = r.db.Query(`SELECT count(user_id) FROM user_email WHERE email = ? LIMIT 1`, entity.Email).Scan(
		&c,
	); err != nil {
		return "", err
	}
	if c > 0 {
		return "", errors.New("email is taken")
	}

	stmtInsertUserEmail := `INSERT INTO user_email
		(email,
		user_id,
		created_at,
		updated_at)
	VALUES(?, ?, ?, ?)`
	batch.Query(stmtInsertUserEmail, entity.Email, entity.ID, time.Now(), time.Now())

	if err = r.db.Query(`SELECT count(user_id) FROM user_login WHERE login = ? LIMIT 1`, entity.Login).Scan(
		&c,
	); err != nil {
		return "", err
	}
	if c > 0 {
		return "", errors.New("login is taken")
	}

	stmtInsertUserLogin := `INSERT INTO user_login
		(login,
		user_id,
		created_at,
		updated_at)
	VALUES(?, ?, ?, ?)`
	batch.Query(stmtInsertUserLogin, entity.Login, entity.ID, time.Now(), time.Now())

	stmtInsertUser := `INSERT INTO user
		(id,
		client_type_id,
		role_id,
		password,
		active,
		expires_at,
		created_at,
		updated_at)
	VALUES(?, ?, ?, ?, ?, ?, ?, ?)`

	batch.Query(stmtInsertUser, entity.ID, entity.ClientTypeID, entity.RoleID, entity.Password, entity.Active, entity.ExpiresAt, time.Now(), time.Now())

	err = r.db.ExecuteBatch(batch)
	if err != nil {
		return "", err
	}

	return entity.ID, nil
}
