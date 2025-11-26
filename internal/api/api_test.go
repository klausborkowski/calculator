package api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi"
	"github.com/mikolabarkouski/calculator/internal/app"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

// MockApp is a mock implementation of AppInterface
type MockApp struct {
	mock.Mock
	app.AppInterface
}

func (m *MockApp) GetPackages() ([]int, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]int), args.Error(1)
}

func (m *MockApp) GetPackagesMap() (map[string]int, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(map[string]int), args.Error(1)
}

func (m *MockApp) AddPackage(packageSize int) error {
	args := m.Called(packageSize)
	return args.Error(0)
}

func (m *MockApp) DeletePackage(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockApp) CalculatePacksNeeded(orderQuantity int, packSizes []int) (map[int]int, error) {
	args := m.Called(orderQuantity, packSizes)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(map[int]int), args.Error(1)
}

func TestHealthCheck(t *testing.T) {
	handler := &Handler{}
	req := httptest.NewRequest("GET", "/health", nil)
	rec := httptest.NewRecorder()

	handler.HealthCheck(rec, req)

	require.Equal(t, http.StatusOK, rec.Code)
	require.Equal(t, "OK", rec.Body.String())
}

func TestAddPackageHandler(t *testing.T) {
	tests := []struct {
		name           string
		requestBody    interface{}
		setupMock      func(*MockApp)
		expectedStatus int
		expectedBody   string
	}{
		{
			name:        "successful add",
			requestBody: map[string]int{"packageSize": 10},
			setupMock: func(m *MockApp) {
				m.On("AddPackage", 10).Return(nil)
			},
			expectedStatus: http.StatusOK,
		},
		{
			name:           "invalid JSON",
			requestBody:    "invalid",
			setupMock:      func(m *MockApp) {},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "Invalid request format\n",
		},
		{
			name:        "app error",
			requestBody: map[string]int{"packageSize": 5},
			setupMock: func(m *MockApp) {
				m.On("AddPackage", 5).Return(errors.New("database error"))
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   "Failed to add package: database error\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockApp := new(MockApp)
			tt.setupMock(mockApp)

			handler := &Handler{app: mockApp}

			var bodyBytes []byte
			if str, ok := tt.requestBody.(string); ok {
				bodyBytes = []byte(str)
			} else {
				bodyBytes, _ = json.Marshal(tt.requestBody)
			}

			req := httptest.NewRequest("POST", "/package", bytes.NewBuffer(bodyBytes))
			req.Header.Set("Content-Type", "application/json")
			rec := httptest.NewRecorder()

			handler.addPackage(rec, req)

			require.Equal(t, tt.expectedStatus, rec.Code)
			if tt.expectedBody != "" {
				require.Equal(t, tt.expectedBody, rec.Body.String())
			}

			mockApp.AssertExpectations(t)
		})
	}
}

func TestDeletePackageHandler(t *testing.T) {
	tests := []struct {
		name           string
		id             string
		setupMock      func(*MockApp)
		expectedStatus int
		expectedBody   string
	}{
		{
			name: "successful delete",
			id:   "1",
			setupMock: func(m *MockApp) {
				m.On("DeletePackage", "1").Return(nil)
			},
			expectedStatus: http.StatusOK,
		},
		{
			name:           "missing id",
			id:             "",
			setupMock:      func(m *MockApp) {},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "ID is required\n",
		},
		{
			name: "app error",
			id:   "999",
			setupMock: func(m *MockApp) {
				m.On("DeletePackage", "999").Return(errors.New("package not found"))
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   "Failed to delete package: package not found\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockApp := new(MockApp)
			tt.setupMock(mockApp)

			handler := &Handler{app: mockApp}

			req := httptest.NewRequest("DELETE", "/package/"+tt.id, nil)
			rctx := chi.NewRouteContext()
			if tt.id != "" {
				rctx.URLParams.Add("id", tt.id)
			}
			req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

			rec := httptest.NewRecorder()

			handler.deletePackage(rec, req)

			require.Equal(t, tt.expectedStatus, rec.Code)
			if tt.expectedBody != "" {
				require.Equal(t, tt.expectedBody, rec.Body.String())
			}

			mockApp.AssertExpectations(t)
		})
	}
}

func TestGetPackagesHandler(t *testing.T) {
	tests := []struct {
		name           string
		setupMock      func(*MockApp)
		expectedStatus int
		expectedBody   map[string]int
	}{
		{
			name: "successful get",
			setupMock: func(m *MockApp) {
				m.On("GetPackagesMap").Return(map[string]int{"1": 10, "2": 20}, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody:   map[string]int{"1": 10, "2": 20},
		},
		{
			name: "app error",
			setupMock: func(m *MockApp) {
				m.On("GetPackagesMap").Return(nil, errors.New("database error"))
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockApp := new(MockApp)
			tt.setupMock(mockApp)

			handler := &Handler{app: mockApp}

			req := httptest.NewRequest("GET", "/packages", nil)
			rec := httptest.NewRecorder()

			handler.getPackages(rec, req)

			require.Equal(t, tt.expectedStatus, rec.Code)
			if tt.expectedBody != nil {
				var response map[string]int
				err := json.Unmarshal(rec.Body.Bytes(), &response)
				require.NoError(t, err)
				require.Equal(t, tt.expectedBody, response)
			}

			mockApp.AssertExpectations(t)
		})
	}
}

func TestCalculateHandler(t *testing.T) {
	tests := []struct {
		name           string
		orderSize      int
		setupMock      func(*MockApp)
		expectedStatus int
		expectedBody   map[int]int
	}{
		{
			name:      "successful calculate",
			orderSize: 10,
			setupMock: func(m *MockApp) {
				m.On("GetPackages").Return([]int{5, 10}, nil)
				m.On("CalculatePacksNeeded", 10, []int{5, 10}).Return(map[int]int{10: 1}, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody:   map[int]int{10: 1},
		},
		{
			name:           "invalid JSON",
			orderSize:      0,
			setupMock:      func(m *MockApp) {},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   nil,
		},
		{
			name:      "app get packages error",
			orderSize: 10,
			setupMock: func(m *MockApp) {
				m.On("GetPackages").Return(nil, errors.New("database error"))
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   nil,
		},
		{
			name:      "calculate error",
			orderSize: 10,
			setupMock: func(m *MockApp) {
				m.On("GetPackages").Return([]int{5, 10}, nil)
				m.On("CalculatePacksNeeded", 10, []int{5, 10}).Return(nil, errors.New("calculation error"))
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockApp := new(MockApp)
			tt.setupMock(mockApp)

			handler := &Handler{app: mockApp}

			var bodyBytes []byte
			if tt.orderSize > 0 {
				bodyBytes, _ = json.Marshal(tt.orderSize)
			} else {
				bodyBytes = []byte("invalid")
			}

			req := httptest.NewRequest("POST", "/calculate", bytes.NewBuffer(bodyBytes))
			req.Header.Set("Content-Type", "application/json")
			rec := httptest.NewRecorder()

			handler.calculate(rec, req)

			require.Equal(t, tt.expectedStatus, rec.Code)
			if tt.expectedBody != nil {
				var response map[int]int
				err := json.Unmarshal(rec.Body.Bytes(), &response)
				require.NoError(t, err)
				require.Equal(t, tt.expectedBody, response)
			}

			mockApp.AssertExpectations(t)
		})
	}
}
