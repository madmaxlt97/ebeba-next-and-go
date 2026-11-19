package service

import (
	"eneba-backend/internal/data"
	"eneba-backend/internal/domain"
)

func GetBanners() []domain.Banner {
	return data.Banners
}