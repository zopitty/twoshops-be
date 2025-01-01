package core

import "github.com/zopitty/twoshops-be/internal/models"

// todo: remove if not required
func ExtractUniqueOutlets(pairs []models.ClosestPair) map[string]models.Outlet {
	selectedOutlets := make(map[string]models.Outlet)

	for _, pair := range pairs {

		// Select outlet for shop_a
		if _, exists := selectedOutlets[pair.ShopA]; !exists {
			selectedOutlets[pair.ShopA] = pair.OutletA
		}
		// Select outlet for shop_b
		if _, exists := selectedOutlets[pair.ShopB]; !exists {
			selectedOutlets[pair.ShopB] = pair.OutletB
		}

	}

	return selectedOutlets
}
