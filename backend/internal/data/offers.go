package data

import "eneba-backend/internal/domain"

var Offers  = []domain.Offer{
{
	// GTA VI
		ID:           "offer-1",
		ProductID:    "prod-1",
		SellerName:   "Geimink.lt",
		SellerRating: 99.4,
		PriceCents:   7299,
		Currency:     "EUR",
		Stock:        15,
	},
	{
		ID:           "offer-2",
		ProductID:    "prod-1",
		SellerName:   "Eneba",
		SellerRating: 99.8,
		PriceCents:   6597,
		Currency:     "EUR",
		Stock:        6,
	},
	{
		ID:           "offer-3",
		ProductID:    "prod-1",
		SellerName:   "GamingKings",
		SellerRating: 86.9,
		PriceCents:   6299,
		Currency:     "EUR",
		Stock:        3,
	},
	// Elden Ring
	{
		ID:           "offer-4",
		ProductID:    "prod-2",
		SellerName:   "Geimink.lt",
		SellerRating: 99.4,
		PriceCents:   3499,
		Currency:     "EUR",
		Stock:        20,
	},
	{
		ID:           "offer-5",
		ProductID:    "prod-2",
		SellerName:   "GamingKings",
		SellerRating: 86.9,
		PriceCents:   3620,
		Currency:     "EUR",
		Stock:        11,
	},

	// FH6
{
		ID:           "offer-6",
		ProductID:    "prod-3",
		SellerName:   "Eneba",
		SellerRating: 99.8,
		PriceCents:   4088,
		Currency:     "EUR",
		Stock:        14,
	},
	{
		ID:           "offer-7",
		ProductID:    "prod-3",
		SellerName:   "GamingKings",
		SellerRating: 96.7,
		PriceCents:   3750,
		Currency:     "EUR",
		Stock:        6,
	},

}