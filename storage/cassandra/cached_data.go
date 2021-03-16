package cassandra

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/saidamir98/goauth_service/storage/repo"

	"github.com/gocql/gocql"
)

type cachedDataRepo struct {
	db                   *gocql.Session
	roleScopePermissions map[string][]string
}

// NewCachedDataRepo ...
func NewCachedDataRepo(db *gocql.Session) repo.CachedDataStorageI {
	return &cachedDataRepo{
		db:                   db,
		roleScopePermissions: getRoleScopePermissions(db),
	}
}

func roleScopeKey(roleID, method, url string) string {
	return roleID + "|" + method + "|" + url
}

func getRoleScopePermissions(db *gocql.Session) map[string][]string {
	t1 := time.Now()
	roleScopePermissions := make(map[string][]string)
	var (
		err error
	)

	stmt1 := `SELECT
	role_id,
	permission_id
	FROM role_permission`

	scanner1 := db.Query(stmt1).Iter().Scanner()
	for scanner1.Next() {
		var (
			roleID       string
			permissionID string
		)

		err = scanner1.Scan(&roleID, &permissionID)
		if err != nil {
			log.Panic(err)
		}

		stmt2 := `SELECT
		method,
		url
		FROM permission_scope
		WHERE permission_id = ?`

		scanner2 := db.Query(stmt2, permissionID).Iter().Scanner()
		for scanner2.Next() {
			var (
				method string
				url    string
			)

			err = scanner2.Scan(&method, &url)
			if err != nil {
				log.Panic(err)
			}

			key := roleScopeKey(roleID, method, url)
			roleScopePermissions[key] = append(roleScopePermissions[key], permissionID)
		}
	}

	count := 0
	for key, element := range roleScopePermissions {
		count++
		fmt.Println("Key:", key, "=>", "Element:", element)
	}

	fmt.Printf("%d in %d ms \n", count, time.Now().Sub(t1).Milliseconds())

	return roleScopePermissions
}

func (r *cachedDataRepo) RoleScopePermission(roleID, method, url string) (res []string, err error) {
	if res, ok := r.roleScopePermissions[roleScopeKey(roleID, method, url)]; ok {
		return res, nil
	}

	return res, errors.New("role scope doesn't exists")
}
