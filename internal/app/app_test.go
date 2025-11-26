package app

import (
	"errors"
	"testing"

	"github.com/klausborkowski/calculator/internal/repo"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

// MockRepository is a mock implementation of RepositoryInterface
type MockRepository struct {
	mock.Mock
	repo.RepositoryInterface
}

func (m *MockRepository) AddPackage(packageSize int) error {
	args := m.Called(packageSize)
	return args.Error(0)
}

func (m *MockRepository) GetPackages() ([]int, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]int), args.Error(1)
}

func (m *MockRepository) GetPackagesMap() (map[string]int, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(map[string]int), args.Error(1)
}

func (m *MockRepository) DeletePackageById(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockRepository) Close() error {
	args := m.Called()
	return args.Error(0)
}

func TestNewApp(t *testing.T) {
	mockRepo := new(MockRepository)
	app := NewApp(mockRepo)

	require.NotNil(t, app)
	require.Equal(t, mockRepo, app.repo)
}

func TestApp_GetPackages(t *testing.T) {
	tests := []struct {
		name      string
		setupMock func(*MockRepository)
		want      []int
		wantErr   bool
	}{
		{
			name: "successful get",
			setupMock: func(m *MockRepository) {
				m.On("GetPackages").Return([]int{1, 2, 3}, nil)
			},
			want:    []int{1, 2, 3},
			wantErr: false,
		},
		{
			name: "repository error",
			setupMock: func(m *MockRepository) {
				m.On("GetPackages").Return(nil, errors.New("database error"))
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(MockRepository)
			tt.setupMock(mockRepo)

			app := NewApp(mockRepo)
			got, err := app.GetPackages()

			if tt.wantErr {
				require.Error(t, err)
				require.Nil(t, got)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.want, got)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestApp_GetPackagesMap(t *testing.T) {
	tests := []struct {
		name      string
		setupMock func(*MockRepository)
		want      map[string]int
		wantErr   bool
	}{
		{
			name: "successful get",
			setupMock: func(m *MockRepository) {
				m.On("GetPackagesMap").Return(map[string]int{"1": 10, "2": 20}, nil)
			},
			want:    map[string]int{"1": 10, "2": 20},
			wantErr: false,
		},
		{
			name: "repository error",
			setupMock: func(m *MockRepository) {
				m.On("GetPackagesMap").Return(nil, errors.New("database error"))
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(MockRepository)
			tt.setupMock(mockRepo)

			app := NewApp(mockRepo)
			got, err := app.GetPackagesMap()

			if tt.wantErr {
				require.Error(t, err)
				require.Nil(t, got)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.want, got)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestApp_AddPackage(t *testing.T) {
	tests := []struct {
		name        string
		packageSize int
		setupMock   func(*MockRepository)
		wantErr     bool
	}{
		{
			name:        "successful add",
			packageSize: 10,
			setupMock: func(m *MockRepository) {
				m.On("AddPackage", 10).Return(nil)
			},
			wantErr: false,
		},
		{
			name:        "repository error",
			packageSize: 5,
			setupMock: func(m *MockRepository) {
				m.On("AddPackage", 5).Return(errors.New("database error"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(MockRepository)
			tt.setupMock(mockRepo)

			app := NewApp(mockRepo)
			err := app.AddPackage(tt.packageSize)

			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestApp_DeletePackage(t *testing.T) {
	tests := []struct {
		name      string
		id        string
		setupMock func(*MockRepository)
		wantErr   bool
	}{
		{
			name: "successful delete",
			id:   "1",
			setupMock: func(m *MockRepository) {
				m.On("DeletePackageById", "1").Return(nil)
			},
			wantErr: false,
		},
		{
			name: "repository error",
			id:   "999",
			setupMock: func(m *MockRepository) {
				m.On("DeletePackageById", "999").Return(errors.New("package not found"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(MockRepository)
			tt.setupMock(mockRepo)

			app := NewApp(mockRepo)
			err := app.DeletePackage(tt.id)

			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}
