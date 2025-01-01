package core

import (
	"fmt"

	"github.com/zopitty/twoshops-be/internal/models"
)

func FindClustersByDistance(
	shops map[string][]models.Outlet, // fetched data
	distanceRanges []float64, // preset range
	maxDistance float64, // user input
) map[string][]map[string]interface{} {

	fmt.Println("outlets:", shops)

	results := make(map[string][]map[string]interface{})

	// convert the shops map into a slice of shop names
	shopNames := make([]string, 0, len(shops))
	for name := range shops {
		shopNames = append(shopNames, name)
	}

	// generate all combinations: one outlet from each shop
	outletCombinations := generateOutletCombinations(shops, shopNames)

	// check proximity and group by distance
	for _, combination := range outletCombinations {
		maxDistanceInCombination := calculateMaxDistance(combination)

		// only consider clusters within the max distance allowed
		if maxDistanceInCombination > maxDistance/1000.0 { // convert maxDistance from meters to kilometers
			continue
		}

		// assign to appropriate distance group
		for _, rangeLimit := range distanceRanges {
			if rangeLimit > maxDistance { // skip ranges beyond the input max distance
				break
			}
			rangeKey := fmt.Sprintf("Within %.0fm", rangeLimit)
			if maxDistanceInCombination <= rangeLimit/1000.0 { // convert rangeLimit to kilometers
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
		fmt.Println("BASE CASE:")
		combination := make(map[string]interface{})
		for key, value := range currentCombination {
			combination[key] = value
		}
		fmt.Println("Generated combination:", combination) // debug
		*combinations = append(*combinations, combination)
		return
	}

	// Recursive case: iterate over all outlets for the current shop
	shopName := shopNames[index]
	for _, outlet := range shops[shopName] {
		currentCombination[shopName] = outlet
		fmt.Printf("currentIndex: %v \n", index)
		fmt.Printf("currentCombination: %v \n", currentCombination)
		generateCombinationHelper(shops, shopNames, index+1, currentCombination, combinations)
	}
}

// calculateMaxDistance calculates the maximum distance between any two outlets in the combination.
func calculateMaxDistance(combination map[string]interface{}) float64 {
	// extract all outlets from the combination into a slice
	outlets := []models.Outlet{}
	for _, outlet := range combination {
		outlets = append(outlets, outlet.(models.Outlet))
	}

	// initialize maxDistance to 0
	maxDistance := 0.0

	// iterate over all pairs of outlets to calculate distances
	for i := 0; i < len(outlets)-1; i++ {
		for j := i + 1; j < len(outlets); j++ {
			distance := CalculateDistance(outlets[i].Latitude, outlets[i].Longitude, outlets[j].Latitude, outlets[j].Longitude)
			if distance > maxDistance {
				maxDistance = distance // update maxDistance if a larger distance is found
			}
		}
	}

	return maxDistance
}
