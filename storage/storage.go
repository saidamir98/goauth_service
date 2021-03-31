package storage

import (
	"github.com/saidamir98/goauth_service/config"
	"github.com/saidamir98/goauth_service/storage/cassandra"
	"github.com/saidamir98/goauth_service/storage/repo"

	"github.com/gocql/gocql"
)

// NewStorageCassandra ...
func NewStorageCassandra(cfg config.Config) CassandraStorageI {
	cassandraConfig := gocql.NewCluster(cfg.CassandraHost)
	cassandraConfig.Authenticator = gocql.PasswordAuthenticator{
		Username: cfg.CassandraUser,
		Password: cfg.CassandraPassword,
	}
	cassandraConfig.Port = cfg.CassandraPort
	cassandraConfig.Keyspace = cfg.CassandraKeyspace
	cassandraConfig.Consistency = gocql.One
	db, err := cassandraConfig.CreateSession()
	if err != nil {
		panic(err)
	}

	return &storageCassandra{
		db:             db,
		cachedDataRepo: cassandra.NewCachedDataRepo(db),
		authRepo:       cassandra.NewAuthRepo(db),
		sessionRepo:    cassandra.NewSessionRepo(db),
		passcodeRepo:   cassandra.NewPasscodeRepo(db),
		clientRepo:     cassandra.NewClientRepo(db),
		clientTypeRepo: cassandra.NewClientTypeRepo(db),
		userRepo:       cassandra.NewUserRepo(db),
	}
}

type storageCassandra struct {
	db             *gocql.Session
	cachedDataRepo repo.CachedDataStorageI
	authRepo       repo.AuthStorageI
	sessionRepo    repo.SessionStorageI
	passcodeRepo   repo.PasscodeStorageI
	clientRepo     repo.ClientStorageI
	clientTypeRepo repo.ClientTypeStorageI
	userRepo       repo.UserStorageI
}

// CassandraStorageI ...
type CassandraStorageI interface {
	CachedData() repo.CachedDataStorageI
	Auth() repo.AuthStorageI
	Session() repo.SessionStorageI
	Passcode() repo.PasscodeStorageI
	Client() repo.ClientStorageI
	ClientType() repo.ClientTypeStorageI
	User() repo.UserStorageI
}

// CachedData ...
func (s storageCassandra) CachedData() repo.CachedDataStorageI {
	return s.cachedDataRepo
}

// Auth ...
func (s storageCassandra) Auth() repo.AuthStorageI {
	return s.authRepo
}

// Session ...
func (s storageCassandra) Session() repo.SessionStorageI {
	return s.sessionRepo
}

// Passcode ...
func (s storageCassandra) Passcode() repo.PasscodeStorageI {
	return s.passcodeRepo
}

// Client ...
func (s storageCassandra) Client() repo.ClientStorageI {
	return s.clientRepo
}

// ClientType ...
func (s storageCassandra) ClientType() repo.ClientTypeStorageI {
	return s.clientTypeRepo
}

// User ...
func (s storageCassandra) User() repo.UserStorageI {
	return s.userRepo
}
