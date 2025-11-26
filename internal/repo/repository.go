package repo

import (
	"database/sql"
	"fmt"
	"log"
	"sort"

	_ "github.com/lib/pq"
)

type Repository struct {
	db *sql.DB
}

// Ensure Repository implements RepositoryInterface
var _ RepositoryInterface = (*Repository)(nil)

func NewRepository(dbHost, dbPort, dbUser, dbPassword, dbName string) (*Repository, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Printf("Error opening database connection: %v", err)
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	if err := db.Ping(); err != nil {
		log.Printf("Error pinging database: %v", err)
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return &Repository{db: db}, nil
}

func (r *Repository) AddPackage(packageSize int) error {
	query := `INSERT INTO package (size) VALUES ($1) RETURNING id`
	var id int
	err := r.db.QueryRow(query, packageSize).Scan(&id)
	if err != nil {
		log.Printf("Error adding package (size: %d): %v", packageSize, err)
		return fmt.Errorf("failed to add package: %w", err)
	}
	return nil
}

func (r *Repository) GetPackages() ([]int, error) {
	query := `SELECT size FROM package ORDER BY size`
	rows, err := r.db.Query(query)
	if err != nil {
		log.Printf("Error querying packages: %v", err)
		return nil, fmt.Errorf("failed to get packages: %w", err)
	}
	defer rows.Close()

	packages := make([]int, 0)
	for rows.Next() {
		var size int
		if err := rows.Scan(&size); err != nil {
			log.Printf("Error scanning package row: %v", err)
			return nil, fmt.Errorf("failed to scan package: %w", err)
		}
		packages = append(packages, size)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Error iterating package rows: %v", err)
		return nil, fmt.Errorf("error iterating packages: %w", err)
	}

	sort.Ints(packages)
	return packages, nil
}

func (r *Repository) GetPackagesMap() (map[string]int, error) {
	query := `SELECT id, size FROM package`
	rows, err := r.db.Query(query)
	if err != nil {
		log.Printf("Error querying packages map: %v", err)
		return nil, fmt.Errorf("failed to get packages map: %w", err)
	}
	defer rows.Close()

	packagesMap := make(map[string]int)
	for rows.Next() {
		var id int
		var size int
		if err := rows.Scan(&id, &size); err != nil {
			log.Printf("Error scanning package map row: %v", err)
			return nil, fmt.Errorf("failed to scan package: %w", err)
		}
		packagesMap[fmt.Sprint(id)] = size
	}

	if err := rows.Err(); err != nil {
		log.Printf("Error iterating package map rows: %v", err)
		return nil, fmt.Errorf("error iterating packages: %w", err)
	}

	return packagesMap, nil
}

func (r *Repository) DeletePackageById(id string) error {
	query := `DELETE FROM package WHERE id = $1`
	result, err := r.db.Exec(query, id)
	if err != nil {
		log.Printf("Error executing delete package query (id: %s): %v", id, err)
		return fmt.Errorf("failed to delete package: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Error getting rows affected for delete (id: %s): %v", id, err)
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		log.Printf("Package with id %s not found for deletion", id)
		return fmt.Errorf("package with id %s not found", id)
	}

	return nil
}

func (r *Repository) Close() error {
	if r.db != nil {
		if err := r.db.Close(); err != nil {
			log.Printf("Error closing database connection: %v", err)
			return err
		}
	}
	return nil
}
