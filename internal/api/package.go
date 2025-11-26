package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

// @Summary Add a new package size
// @Description Adds a new package size to the system
// @Tags Packages
// @Accept json
// @Produce json
// @Param request body object true "Package size request" SchemaExample({"packageSize": 10})
// @Success 200 {string} string "Package added successfully"
// @Failure 400 {string} string "Invalid request format"
// @Router /package [post]
func (h *Handler) addPackage(w http.ResponseWriter, r *http.Request) {
	var request struct {
		PackageSize int `json:"packageSize"`
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading request body: %v", err)
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if err := json.Unmarshal(body, &request); err != nil {
		log.Printf("Error unmarshaling request body: %v", err)
		http.Error(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	if err := h.app.AddPackage(request.PackageSize); err != nil {
		log.Printf("Error adding package (size: %d): %v", request.PackageSize, err)
		http.Error(w, "Failed to add package: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// @Summary Delete a package
// @Description Deletes a package by its ID
// @Tags Packages
// @Accept json
// @Produce json
// @Param id path string true "ID of the package to delete"
// @Success 200 {string} string "Package deleted successfully"
// @Failure 400 {string} string "Invalid request format"
// @Router /package/{id} [delete]
func (h *Handler) deletePackage(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		log.Printf("Error: missing package ID in delete request")
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	if err := h.app.DeletePackage(id); err != nil {
		log.Printf("Error deleting package (id: %s): %v", id, err)
		http.Error(w, "Failed to delete package: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// @Summary Get all package sizes
// @Description Retrieves a list of all available package sizes
// @Tags Packages
// @Accept json
// @Produce json
// @Success 200 {object} map[string]int "List of package sizes"
// @Failure 500 {string} string "Failed to encode response"
// @Router /packages [get]
func (h *Handler) getPackages(w http.ResponseWriter, r *http.Request) {
	packages, err := h.app.GetPackagesMap()
	if err != nil {
		log.Printf("Error getting packages: %v", err)
		http.Error(w, "Failed to get packages: "+err.Error(), http.StatusInternalServerError)
		return
	}

	responseBody, err := json.Marshal(packages)
	if err != nil {
		log.Printf("Error marshaling packages response: %v", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseBody)
}
