package service

import (
	"eneba-backend/internal/data"
	"eneba-backend/internal/domain"
	"errors"
)

type ProductCard struct {
	Product domain.Product `json:"product"`
	LowestPrice int64 `json:"lowestPrice"`
	OfferNumber int `json:"offerNumber"`
}
type ProductDetails struct {
	Product domain.Product `json:"product"`
	Offers []domain.Offer `json:"offers"`
}

func GetProductCardByID(productID string) (ProductCard, error) {

	var foundProduct domain.Product
	var productExists bool

	
	for _, p := range data.Products {
		if p.ID == productID {
			foundProduct = p
			productExists = true
			break
		}
	}

	if !productExists {
		return ProductCard{}, errors.New("product not found")
	}

	var lowestPrice int64
	var offerCount int

	for _, offer := range data.Offers {
		if offer.ProductID == foundProduct.ID {
			if offerCount == 0 || offer.PriceCents < lowestPrice {
				lowestPrice = offer.PriceCents
			}
			offerCount++
		}
	}

	return ProductCard{
		Product:          foundProduct,
		LowestPrice: lowestPrice,
		OfferNumber:       offerCount,
	}, nil
}

func GetAllProductCards() []ProductCard {
	var cards []ProductCard

	for _, p := range data.Products {
		card, err := GetProductCardByID(p.ID)
		if err != nil {
			continue
		}
		cards = append(cards, card)
	}
	return cards
}

func GetProductDetailsByID(productID string) (ProductDetails, error) {

	var foundProduct domain.Product
	var productExists bool

	for _, p := range data.Products {
		if p.ID == productID {
			foundProduct = p
			productExists = true
			break
			}
		}
		if !productExists {
			return ProductDetails{}, errors.New("Product not found")
		}

		var offers []domain.Offer
		for _, offer := range data.Offers {
			if offer.ProductID == foundProduct.ID {
				offers = append(offers, offer)
			}
		}
		return ProductDetails{
			Product: foundProduct,
			Offers: offers,
		}, nil
	
}