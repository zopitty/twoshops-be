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

	distanceRanges := []float64{100, 500, 1000} // meters

	// Find clusters and group them by distance
	groupedClusters := core.FindClustersByDistance(shops, distanceRanges, req.MaxDistance)

	// Debug output
	fmt.Println("Final grouped clusters:", groupedClusters)

	// Respond with grouped clusters
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(groupedClusters)
}
