package core

import "github.com/zopitty/twoshops-be/internal/models"

// todo: remove if not required
func FilterByProximity(pairs []models.ClosestPair, maxDistance float64) []models.ClosestPair {
	if maxDistance <= 0 {
		return pairs
	}

	var filtered []models.ClosestPair
	for _, pair := range pairs {
		if pair.Distance <= maxDistance {
			filtered = append(filtered, pair)
		}
	}
	return filtered
}
