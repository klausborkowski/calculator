package repo

// RepositoryInterface defines the interface for Repository to enable mocking in tests
type RepositoryInterface interface {
	AddPackage(packageSize int) error
	GetPackages() ([]int, error)
	GetPackagesMap() (map[string]int, error)
	DeletePackageById(id string) error
	Close() error
}

