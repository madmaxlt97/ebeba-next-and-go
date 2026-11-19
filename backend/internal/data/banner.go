package data

import "eneba-backend/internal/domain"

var Banners = []domain.Banner{
	{
		ID:         "banner-1",
		Title:      "Grand Theft Auto VI",
		Subtitle:   "Instant digital key delivery. Compare offers from trusted sellers.",
		ImageURL:   "/images/banners/gta-6-hero.jpg",
		TargetSlug: "grand-theft-auto-vi-ps5-global",
		BadgeText:  "PRE-ORDER / HOT",
		Source:     "promo-service",
	},
	{
		ID:         "banner-2",
		Title:      "Weekend Key Deals",
		Subtitle:   "Best prices across PlayStation, Xbox, and Nintendo.",
		ImageURL:   "/images/banners/weekend-deals.jpg",
		TargetSlug: "",
		BadgeText:  "DEALS",
		Source:     "deals-service",
	},
}