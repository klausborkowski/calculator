package app

import (
	"github.com/klausborkowski/calculator/internal/repo"
)

type App struct {
	repo repo.RepositoryInterface
}

// Ensure App implements AppInterface
var _ AppInterface = (*App)(nil)

func NewApp(r repo.RepositoryInterface) *App {
	return &App{repo: r}
}

// GetPackages returns a slice of all stored package sizes
func (a *App) GetPackages() ([]int, error) {
	return a.repo.GetPackages()
}

// GetPackagesMap returns package sizes as a key-value collection
func (a *App) GetPackagesMap() (map[string]int, error) {
	return a.repo.GetPackagesMap()
}

// AddPackage adds a new package size
func (a *App) AddPackage(packageSize int) error {
	return a.repo.AddPackage(packageSize)
}

// DeletePackage deletes a package by its ID
func (a *App) DeletePackage(id string) error {
	return a.repo.DeletePackageById(id)
}
