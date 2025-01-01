package core

import "github.com/zopitty/twoshops-be/internal/models"

func FindClusters(shops map[string][]models.Outlet, maxDistance float64) []map[string]interface{} {
	// results := make(map[string][]map[string]interface{})
	var results []map[string]interface{}

	// Convert shops to map into a slice of shop names
	shopNames := make([]string, 0, len(shops))
	for name := range shops {
		shopNames = append(shopNames, name)
	}

	// generate all combinations: one outlet for each shop
	outletCombinations := generateOutletCombinations(shops, shopNames)

	// check proximity
	for _, combination := range outletCombinations {
		if isCombinationInProximity(combination, maxDistance) {
			results = append(results, combination)

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

func generateCombinationHelper(shops map[string][]models.Outlet, shopNames []string, index int, currentCombination map[string]interface{}, combinations *[]map[string]interface{}) {
	if index == len(shopNames) {
		combination := make(map[string]interface{})
		for key, value := range currentCombination {
			combination[key] = value
		}
		*combinations = append(*combinations, combination)
		return
	}

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

	for i := 0; i < len(outlets)-1; i++ {
		for j := i + 1; j < len(outlets); j++ {
			distance := CalculateDistance(outlets[i].Latitude, outlets[i].Longitude, outlets[j].Latitude, outlets[j].Longitude)
			if distance > maxDistance {
				return false
			}
		}
	}
	return true
}
