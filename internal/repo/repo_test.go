package repo

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

// Ensure Repository implements RepositoryInterface
var _ RepositoryInterface = (*Repository)(nil)

func TestRepository_AddPackage(t *testing.T) {
	tests := []struct {
		name        string
		packageSize int
		setupMock   func(sqlmock.Sqlmock)
		wantErr     bool
	}{
		{
			name:        "successful add",
			packageSize: 10,
			setupMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(`INSERT INTO package \(size\) VALUES \(\$1\) RETURNING id`).
					WithArgs(10).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
			},
			wantErr: false,
		},
		{
			name:        "database error",
			packageSize: 5,
			setupMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(`INSERT INTO package \(size\) VALUES \(\$1\) RETURNING id`).
					WithArgs(5).
					WillReturnError(sql.ErrConnDone)
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			require.NoError(t, err)
			defer db.Close()

			repo := &Repository{db: db}
			tt.setupMock(mock)

			err = repo.AddPackage(tt.packageSize)

			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}

			require.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestRepository_GetPackages(t *testing.T) {
	tests := []struct {
		name      string
		setupMock func(sqlmock.Sqlmock)
		want      []int
		wantErr   bool
	}{
		{
			name: "successful get",
			setupMock: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"size"}).
					AddRow(1).
					AddRow(2).
					AddRow(3)
				mock.ExpectQuery(`SELECT size FROM package ORDER BY size`).
					WillReturnRows(rows)
			},
			want:    []int{1, 2, 3},
			wantErr: false,
		},
		{
			name: "empty result",
			setupMock: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"size"})
				mock.ExpectQuery(`SELECT size FROM package ORDER BY size`).
					WillReturnRows(rows)
			},
			want:    []int{},
			wantErr: false,
		},
		{
			name: "database error",
			setupMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(`SELECT size FROM package ORDER BY size`).
					WillReturnError(sql.ErrConnDone)
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			require.NoError(t, err)
			defer db.Close()

			repo := &Repository{db: db}
			tt.setupMock(mock)

			got, err := repo.GetPackages()

			if tt.wantErr {
				require.Error(t, err)
				require.Nil(t, got)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.want, got)
			}

			require.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestRepository_GetPackagesMap(t *testing.T) {
	tests := []struct {
		name      string
		setupMock func(sqlmock.Sqlmock)
		want      map[string]int
		wantErr   bool
	}{
		{
			name: "successful get",
			setupMock: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "size"}).
					AddRow(1, 10).
					AddRow(2, 20).
					AddRow(3, 30)
				mock.ExpectQuery(`SELECT id, size FROM package`).
					WillReturnRows(rows)
			},
			want:    map[string]int{"1": 10, "2": 20, "3": 30},
			wantErr: false,
		},
		{
			name: "database error",
			setupMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(`SELECT id, size FROM package`).
					WillReturnError(sql.ErrConnDone)
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			require.NoError(t, err)
			defer db.Close()

			repo := &Repository{db: db}
			tt.setupMock(mock)

			got, err := repo.GetPackagesMap()

			if tt.wantErr {
				require.Error(t, err)
				require.Nil(t, got)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.want, got)
			}

			require.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestRepository_DeletePackageById(t *testing.T) {
	tests := []struct {
		name      string
		id        string
		setupMock func(sqlmock.Sqlmock)
		wantErr   bool
	}{
		{
			name: "successful delete",
			id:   "1",
			setupMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(`DELETE FROM package WHERE id = \$1`).
					WithArgs("1").
					WillReturnResult(sqlmock.NewResult(0, 1))
			},
			wantErr: false,
		},
		{
			name: "package not found",
			id:   "999",
			setupMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(`DELETE FROM package WHERE id = \$1`).
					WithArgs("999").
					WillReturnResult(sqlmock.NewResult(0, 0))
			},
			wantErr: true,
		},
		{
			name: "database error",
			id:   "1",
			setupMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(`DELETE FROM package WHERE id = \$1`).
					WithArgs("1").
					WillReturnError(sql.ErrConnDone)
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			require.NoError(t, err)
			defer db.Close()

			repo := &Repository{db: db}
			tt.setupMock(mock)

			err = repo.DeletePackageById(tt.id)

			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}

			require.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestRepository_Close(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)

	repo := &Repository{db: db}
	mock.ExpectClose()

	err = repo.Close()
	require.NoError(t, err)
	require.NoError(t, mock.ExpectationsWereMet())
}
