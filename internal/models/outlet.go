package models

type Outlet struct {
	Name           string  `json:"name"`
	Address        string  `json:"address"`
	Latitude       float64 `json:"latitude"`
	Longitude      float64 `json:"longitude"`
	BusinessStatus string  `json:"business_status"`
}

type ClosestPair struct {
	ShopA    string  `json:"shop_a"`
	OutletA  Outlet  `json:"outlet_a"`
	ShopB    string  `json:"shop_b"`
	OutletB  Outlet  `json:"outlet_b"`
	Distance float64 `json:"distance"`
}

type ShopRequest struct {
	Shops       []string `json:"shops"`
	MaxDistance float64  `json:"max_distance"`
}
