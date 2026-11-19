package data

import (
	"eneba-backend/internal/domain"
)

var Products = []domain.Product{
	{
		ID:          "prod-1",
		Slug:        "gta-vi-xbox-global",
		Title:       "GTA VI",
		Description: "Open-world action adventure set in Leonida. Digital game key.",
		CoverURL:    "/images/gtavi.jpg",
		Platform:    "Xbox",
		Region:      "Global",
		Tags:        []string{"gta", "rockstar", "action", "open-world"},
	},
	{
		ID:          "prod-2",
		Slug:        "elden-ring-ps5-europe",
		Title:       "Elden Ring",
		Description: "Dark fantasy action RPG by FromSoftware.",
		CoverURL:    "/images/elden-ring.jpg",
		Platform:    "PlayStation",
		Region:      "Europe",
		Tags:        []string{"soulslike", "rpg", "fantasy"},
	},
	{
		ID:          "prod-3",
		Slug:        "forza-horizon-6-xbox-global",
		Title:       "Forza Horizon 6",
		Description: "Open-world racing adventure across Japan.",
		CoverURL:    "/images/forza-horizon-6.jpg",
		Platform:    "Xbox",
		Region:      "Global",
		Tags:        []string{"racing", "cars", "open-world"},
	},
}