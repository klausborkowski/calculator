package app

// AppInterface defines the interface for App to enable mocking in tests
type AppInterface interface {
	GetPackages() ([]int, error)
	GetPackagesMap() (map[string]int, error)
	AddPackage(packageSize int) error
	DeletePackage(id string) error
	CalculatePacksNeeded(orderQuantity int, packSizes []int) (map[int]int, error)
}

