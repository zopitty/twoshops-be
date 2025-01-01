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

	// shops: and max_distance:
	var req models.ShopRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	apiKey := config.GetGoogleAPIKey()
	shops := make(map[string][]models.Outlet)

	fmt.Println("searching for shops:", req.Shops)

	for _, shop := range req.Shops {
		fmt.Println("fetching shop: ", shop)
		outlets, err := google.FetchOutlets(apiKey, shop)
		fmt.Println("fetched outlets: ", outlets)
		if err != nil {
			http.Error(w, "Failed to fetch outlets", http.StatusInternalServerError)
			fmt.Println(err)
			return
		}
		shops[shop] = outlets
	}

	// fixed distance range
	// within 100, within 500, within 1000 of each other
	distanceRanges := []float64{100, 250, 500, 750, 1000} // meters

	// Find clusters and group them by distance
	groupedClusters := core.FindClustersByDistance(shops, distanceRanges, req.MaxDistance)

	// debug
	fmt.Println("Final grouped clusters:", groupedClusters)

	// Respond with grouped clusters
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(groupedClusters)
}
