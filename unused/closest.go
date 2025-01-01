package core

import (
	"math"

	"github.com/zopitty/twoshops-be/internal/models"
)

// todo: remove if not required
func FindClosestPair(shopA, shopB string, outletsA, outletsB []models.Outlet) []models.ClosestPair {
	// var closestPair models.ClosestPair
	// minDistance := math.MaxFloat64

	// for _, outletA := range outletsA {
	// 	for _, outletB := range outletsB {
	// 		distance := CalculateDistance(outletA.Latitude, outletA.Longitude, outletB.Latitude, outletB.Longitude)
	// 		if distance < minDistance {
	// 			minDistance = distance
	// 			closestPair = models.ClosestPair{
	// 				OutletA:  outletA,
	// 				OutletB:  outletB,
	// 				Distance: distance,
	// 			}
	// 		}
	// 	}
	// }
	// fmt.Println(closestPair)
	// return closestPair, minDistance

	var results []models.ClosestPair

	for _, outletA := range outletsA {
		for _, outletB := range outletsB {
			distance := CalculateDistance(outletA.Latitude, outletA.Longitude, outletB.Latitude, outletB.Longitude)
			results = append(results, models.ClosestPair{
				ShopA:    shopA,
				ShopB:    shopB,
				OutletA:  outletA,
				OutletB:  outletB,
				Distance: distance,
			})
		}
	}

	return results

}
func CalculateDistance(lat1, lon1, lat2, lon2 float64) float64 {
	const earthRadius = 6371 // in km
	dLat := (lat2 - lat1) * math.Pi / 180.0
	dLon := (lon2 - lon1) * math.Pi / 180.0

	a := math.Sin(dLat/2)*math.Sin(dLat/2) + math.Cos(lat1*math.Pi/180.0)*math.Cos(lat2*math.Pi/180.0)*math.Sin(dLon/2)*math.Sin(dLon/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return earthRadius * c
}

func CalculateDistances(shops map[string][]models.Outlet) []models.ClosestPair {
	var results []models.ClosestPair
	keys := make([]string, 0, len(shops))
	for key := range shops {
		keys = append(keys, key)
	}

	for i := 0; i < len(keys)-1; i++ {
		for j := i + 1; j < len(keys); j++ {
			shopA, shopB := keys[i], keys[j]
			outletsA, outletsB := shops[shopA], shops[shopB]

			pairs := FindClosestPair(shopA, shopB, outletsA, outletsB)
			results = append(results, pairs...)
		}
	}

	return results
}
