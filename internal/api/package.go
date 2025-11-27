package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// @Summary Add a new package size
// @Description Adds a new package size to the system
// @Tags Packages
// @Accept json
// @Produce json
// @Param packageSize body int true "Package size to add"
// @Success 200 {string} string "Package added successfully"
// @Failure 400 {string} string "Invalid request format"
// @Router /package [post]
func (h *Handler) addPackage(w http.ResponseWriter, r *http.Request) {
	var packageSize int
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return

	}
	defer r.Body.Close()

	if err := json.Unmarshal(body, &packageSize); err != nil {
		http.Error(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	h.app.AddPackage(packageSize)
	w.WriteHeader(http.StatusOK)
}

// @Summary Delete a package
// @Description Deletes a package by its ID
// @Tags Packages
// @Accept json
// @Produce json
// @Param id body int true "ID of the package to delete"
// @Success 200 {string} string "Package deleted successfully"
// @Failure 400 {string} string "Invalid request format"
// @Router /package [delete]
func (h *Handler) deletePackage(w http.ResponseWriter, r *http.Request) {
	var id int
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return

	}
	defer r.Body.Close()

	if err := json.Unmarshal(body, &id); err != nil {
		http.Error(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	h.app.DeletePackage(fmt.Sprint(id))
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
	packages := h.app.GetPackagesMap()
	responseBody, err := json.Marshal(packages)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseBody)

}
