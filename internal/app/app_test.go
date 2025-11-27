package app

import (
	"testing"

	"github.com/mikolabarkouski/calculator/internal/repo"
	"github.com/stretchr/testify/require"
)

func TestAppMethods(t *testing.T) {
	repository := repo.NewRepository([]int{1, 2})
	app := NewApp(repository)

	// Test AddPackage
	app.AddPackage(10)
	packages := app.GetPackages()
	require.Contains(t, packages, 10, "Package 10 should be added")

	// Test GetPackagesMap
	packagesMap := app.GetPackagesMap()
	require.Equal(t, 10, packagesMap["3"], "Package 10 should exist in map with count 1")

	// Test DeletePackage
	app.DeletePackage("10")
	packagesMap = app.GetPackagesMap()
	require.NotContains(t, packagesMap, "10", "Package 10 should be deleted")
}
