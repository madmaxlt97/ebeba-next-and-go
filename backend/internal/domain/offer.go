package domain

type Offer struct {
	ID           string  `json:"id"`
	ProductID    string  `json:"productId"`
	SellerName   string  `json:"sellerName"`
	SellerRating float64 `json:"sellerRating"`
	PriceCents   int64   `json:"priceCents"`
	Currency     string  `json:"currency"`
	Stock        int     `json:"stock"`
}
