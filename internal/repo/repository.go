package repo

import (
	"fmt"
	"sort"
)

type Repository struct {
	packages map[string]int
}

func NewRepository(packages []int) *Repository {
	//sort on init-step for performance purporse
	sort.Ints(packages)

	packageMap := make(map[string]int)
	for i, pkg := range packages {
		packageMap[fmt.Sprint(i+1)] = pkg
	}
	return &Repository{packages: packageMap}
}
func (r *Repository) AddPackage(packageSize int) {
	key := fmt.Sprint(len(r.packages) + 1)
	r.packages[key] = packageSize
}

func (r *Repository) GetPackages() []int {
	values := make([]int, 0, len(r.packages))
	for _, v := range r.packages {
		values = append(values, v)
	}
	sort.Ints(values)
	return values
}

func (r *Repository) GetPackagesMap() map[string]int {

	return r.packages
}

func (r *Repository) DeletePackageById(id string) {
	delete(r.packages, id)
}
