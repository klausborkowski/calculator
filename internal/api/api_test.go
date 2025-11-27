package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"

	"github.com/mikolabarkouski/calculator/internal/app"
	"github.com/mikolabarkouski/calculator/internal/repo"
	"github.com/mikolabarkouski/calculator/templates"
	"github.com/stretchr/testify/require"
)

func TestHealthCheckHandler(t *testing.T) {
	handler := &Handler{}
	req := httptest.NewRequest("GET", "/health", nil)
	rec := httptest.NewRecorder()

	handler.HealthCheck(rec, req)

	require.Equal(t, http.StatusOK, rec.Code)
	require.Equal(t, "OK", rec.Body.String())
}

func TestAddPackageHandler(t *testing.T) {
	repo := repo.NewRepository([]int{1})
	app := app.NewApp(repo)
	handler := &Handler{app: app}

	req := httptest.NewRequest("POST", "/package", bytes.NewBufferString(`5`))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	handler.addPackage(rec, req)

	require.Equal(t, http.StatusOK, rec.Code)
}

func TestAddPackageHandler_InvalidJSON(t *testing.T) {
	handler := &Handler{}

	req := httptest.NewRequest("POST", "/package", bytes.NewBufferString(`invalid`))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	handler.addPackage(rec, req)

	require.Equal(t, http.StatusBadRequest, rec.Code)
	require.Equal(t, "Invalid request format\n", rec.Body.String())
}

func TestDeletePackageHandler(t *testing.T) {
	repo := repo.NewRepository([]int{1})
	app := app.NewApp(repo)
	handler := &Handler{app: app}

	req := httptest.NewRequest("DELETE", "/package", bytes.NewBufferString(`5`))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	handler.deletePackage(rec, req)

	require.Equal(t, http.StatusOK, rec.Code)
}

func TestDeletePackageHandler_InvalidJSON(t *testing.T) {
	handler := &Handler{}

	req := httptest.NewRequest("DELETE", "/package", bytes.NewBufferString(`invalid`))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	handler.deletePackage(rec, req)

	require.Equal(t, http.StatusBadRequest, rec.Code)
	require.Equal(t, "Invalid request format\n", rec.Body.String())
}

func TestGetPackagesHandler(t *testing.T) {
	repo := repo.NewRepository([]int{1})
	app := app.NewApp(repo)
	handler := &Handler{app: app}

	expectedPackages := map[string]int{"1": 1}

	req := httptest.NewRequest("GET", "/packages", nil)
	rec := httptest.NewRecorder()

	handler.getPackages(rec, req)

	require.Equal(t, http.StatusOK, rec.Code)

	var response map[string]int
	err := json.Unmarshal(rec.Body.Bytes(), &response)
	require.NoError(t, err)
	require.Equal(t, expectedPackages, response)
}

func TestRenderAppHandler(t *testing.T) {
	repo := repo.NewRepository([]int{1})
	app := app.NewApp(repo)
	handler := &Handler{app: app}

	expectedPackages := map[string]int{"1": 1}

	req := httptest.NewRequest("GET", "/app", nil)
	rec := httptest.NewRecorder()

	templates.IndexHTML = `{{ range $size, $count := . }}Package Size: {{$size}}, Count: {{$count}}{{end}}`

	handler.RenderApp(rec, req)

	require.Equal(t, http.StatusOK, rec.Code)

	tmpl := template.Must(template.New("index").Parse(templates.IndexHTML))
	expectedResponse := bytes.NewBufferString("")
	_ = tmpl.Execute(expectedResponse, expectedPackages)

	require.Equal(t, expectedResponse.String(), rec.Body.String())
}
