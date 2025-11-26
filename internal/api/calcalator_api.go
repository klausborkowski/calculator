package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// @Summary Calculate package sizes needed
// @Description Calculates the number of packages required for an order size
// @Tags Orders
// @Accept json
// @Produce json
// @Param orderSize body int true "Order size"
// @Success 200 {object} map[string]interface{} "Calculated package details"
// @Failure 400 {string} string "Invalid request format"
// @Failure 500 {string} string "Internal server error"
// @Router /calculate [post]
func (h *Handler) calculate(w http.ResponseWriter, r *http.Request) {
	var orderSizeRequest int
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading request body: %v", err)
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if err := json.Unmarshal(body, &orderSizeRequest); err != nil {
		log.Printf("Error unmarshaling order size request: %v", err)
		http.Error(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	packageSizes, err := h.app.GetPackages()
	if err != nil {
		log.Printf("Error getting packages for calculation: %v", err)
		http.Error(w, "Failed to get packages: "+err.Error(), http.StatusInternalServerError)
		return
	}

	result, err := h.app.CalculatePacksNeeded(orderSizeRequest, packageSizes)
	if err != nil {
		log.Printf("Error calculating packs needed (order size: %d): %v", orderSizeRequest, err)
		http.Error(w, "Failed to calculate packs needed: "+err.Error(), http.StatusInternalServerError)
		return
	}

	responseBody, err := json.Marshal(result)
	if err != nil {
		log.Printf("Error marshaling calculation result: %v", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseBody)
}
