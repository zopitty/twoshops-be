package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/zopitty/twoshops-be/config"
	"github.com/zopitty/twoshops-be/core"
	"github.com/zopitty/twoshops-be/internal/google"
	"github.com/zopitty/twoshops-be/internal/models"
)

func HandleFindClosest(w http.ResponseWriter, r *http.Request) {
	var req models.ShopRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	apiKey := config.GetGoogleAPIKey()
	shops := make(map[string][]models.Outlet)

	fmt.Println("searching for shops:", req.Shops)

	for _, shop := range req.Shops {
		outlets, err := google.FetchOutlets(apiKey, shop)
		fmt.Println("outlets: \n", outlets)
		if err != nil {
			http.Error(w, "Failed to fetch outlets", http.StatusInternalServerError)
			fmt.Println(err)
			return
		}
		shops[shop] = outlets
	}

	// Find the closest pair
	allPairs := core.CalculateDistances(shops)

	// Responsd with the closest pair
	maxDistanceKm := req.MaxDistance / 1000.0
	filteredPairs := core.FilterByProximity(allPairs, maxDistanceKm)

	// Extract unique outlets for each shop
	uniqueOutlets := core.ExtractUniqueOutlets(filteredPairs)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(uniqueOutlets)
}
