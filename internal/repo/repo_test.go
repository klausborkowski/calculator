package repo

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRepository(t *testing.T) {
	initialPackages := []int{10, 5, 15}
	repo := NewRepository(initialPackages)

	expectedSorted := []int{5, 10, 15}
	require.Equal(t, expectedSorted, repo.GetPackages(), "GetPackages should return sorted list")

	packagesMap := repo.GetPackagesMap()
	require.Equal(t, 3, len(packagesMap), "Packages map should have 3 elements")

	repo.AddPackage(20)
	require.Contains(t, repo.GetPackagesMap(), "4", "Package 4 should be added")
	require.ElementsMatch(t, []int{5, 10, 15, 20}, repo.GetPackages(), "New package should be included")

	repo.DeletePackageById("2")
	require.NotContains(t, repo.GetPackagesMap(), "2", "Package 2 should be deleted")
}
