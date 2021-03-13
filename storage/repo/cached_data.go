package repo

// CachedDataStorageI ...
type CachedDataStorageI interface {
	RoleScopePermission(roleID, method, url string) (res []string, err error)
}
