package core

import (
	"fmt"

	"github.com/zopitty/twoshops-be/internal/models"
)

func FindClustersByDistance(shops map[string][]models.Outlet, distanceRanges []float64, maxDistance float64) map[string][]map[string]interface{} {
	results := make(map[string][]map[string]interface{})

	// Convert the shops map into a slice of shop names
	shopNames := make([]string, 0, len(shops))
	for name := range shops {
		shopNames = append(shopNames, name)
	}

	// Generate all combinations: one outlet from each shop
	outletCombinations := generateOutletCombinations(shops, shopNames)

	// Check proximity and group by distance
	for _, combination := range outletCombinations {
		maxDistanceInCombination := calculateMaxDistance(combination)

		// Only consider clusters within the max distance allowed
		if maxDistanceInCombination > maxDistance/1000.0 { // Convert maxDistance from meters to kilometers
			continue
		}

		// Assign to appropriate distance group
		for _, rangeLimit := range distanceRanges {
			if rangeLimit > maxDistance { // Skip ranges beyond the input max distance
				break
			}
			rangeKey := fmt.Sprintf("Within %.0fm", rangeLimit)
			if maxDistanceInCombination <= rangeLimit/1000.0 { // Convert rangeLimit to kilometers
				results[rangeKey] = append(results[rangeKey], combination)
				break
			}
		}
	}

	return results
}

func generateOutletCombinations(shops map[string][]models.Outlet, shopNames []string) []map[string]interface{} {
	if len(shopNames) == 0 {
		return nil
	}

	var combinations []map[string]interface{}
	generateCombinationHelper(shops, shopNames, 0, map[string]interface{}{}, &combinations)
	return combinations
}

func generateCombinationHelper(
	shops map[string][]models.Outlet,
	shopNames []string,
	index int,
	currentCombination map[string]interface{},
	combinations *[]map[string]interface{},
) {
	// Base case: all shops have been processed
	if index == len(shopNames) {
		combination := make(map[string]interface{})
		for key, value := range currentCombination {
			combination[key] = value
		}
		fmt.Println("Generated combination:", combination) // Debug output
		*combinations = append(*combinations, combination)
		return
	}

	// Recursive case: iterate over all outlets for the current shop
	shopName := shopNames[index]
	for _, outlet := range shops[shopName] {
		currentCombination[shopName] = outlet
		generateCombinationHelper(shops, shopNames, index+1, currentCombination, combinations)
	}
}

func isCombinationInProximity(combination map[string]interface{}, maxDistance float64) bool {
	outlets := []models.Outlet{}
	for _, outlet := range combination {
		outlets = append(outlets, outlet.(models.Outlet))
	}

	// Debug output
	fmt.Println("Checking combination for proximity:", combination)

	for i := 0; i < len(outlets)-1; i++ {
		for j := i + 1; j < len(outlets); j++ {
			distance := CalculateDistance(outlets[i].Latitude, outlets[i].Longitude, outlets[j].Latitude, outlets[j].Longitude)
			fmt.Printf("Distance between %v and %v: %.2f km\n", outlets[i], outlets[j], distance) // Debug
			if distance > maxDistance {
				return false
			}
		}
	}
	return true
}

// calculateMaxDistance calculates the maximum distance between any two outlets in the combination.
func calculateMaxDistance(combination map[string]interface{}) float64 {
	// Extract all outlets from the combination into a slice
	outlets := []models.Outlet{}
	for _, outlet := range combination {
		outlets = append(outlets, outlet.(models.Outlet))
	}

	// Initialize maxDistance to 0
	maxDistance := 0.0

	// Iterate over all pairs of outlets to calculate distances
	for i := 0; i < len(outlets)-1; i++ {
		for j := i + 1; j < len(outlets); j++ {
			distance := CalculateDistance(outlets[i].Latitude, outlets[i].Longitude, outlets[j].Latitude, outlets[j].Longitude)
			if distance > maxDistance {
				maxDistance = distance // Update maxDistance if a larger distance is found
			}
		}
	}

	return maxDistance
}
