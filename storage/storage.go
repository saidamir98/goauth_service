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
	session, err := cassandraConfig.CreateSession()
	if err != nil {
		panic(err)
	}

	return &storageCassandra{
		db:             session,
		cachedDataRepo: cassandra.NewCachedDataRepo(session),
		authRepo:       cassandra.NewAuthRepo(session),
		clientTypeRepo: cassandra.NewClientTypeRepo(session),
	}
}

type storageCassandra struct {
	db             *gocql.Session
	cachedDataRepo repo.CachedDataStorageI
	authRepo       repo.AuthStorageI
	clientTypeRepo repo.ClientTypeStorageI
}

// CassandraStorageI ...
type CassandraStorageI interface {
	CachedData() repo.CachedDataStorageI
	Auth() repo.AuthStorageI
	ClientType() repo.ClientTypeStorageI
}

// CachedData ...
func (s storageCassandra) CachedData() repo.CachedDataStorageI {
	return s.cachedDataRepo
}

// Auth ...
func (s storageCassandra) Auth() repo.AuthStorageI {
	return s.authRepo
}

// ClientType ...
func (s storageCassandra) ClientType() repo.ClientTypeStorageI {
	return s.clientTypeRepo
}
