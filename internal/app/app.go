package app

import (
	"github.com/mikolabarkouski/calculator/internal/repo"
)

type App struct {
	repo *repo.Repository
}

func NewApp(r *repo.Repository) *App {
	return &App{repo: r}
}

// get slice(array) of all stored packages (package-sizes)
func (a *App) GetPackages() []int {
	packages := a.repo.GetPackages()
	return packages
}

// get package-sizes as key-value collection
func (a *App) GetPackagesMap() map[string]int {
	packagesMap := a.repo.GetPackagesMap()
	return packagesMap
}

// add package size
func (a *App) AddPackage(packageSize int) {
	a.repo.AddPackage(packageSize)
}

// delete package size
func (a *App) DeletePackage(id string) {
	a.repo.DeletePackageById(id)
}
