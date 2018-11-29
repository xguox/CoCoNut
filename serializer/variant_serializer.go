package serializer

import (
	"coconut/model"
	"time"
)

type VariantSerializer struct {
	model.Variant
}

type VariantsSerializer struct {
	Variants []model.Variant
}

type VariantResponse struct {
	ID        uint      `json:"id"`
	Price     float32   `json:"price"`
	Sku       string    `json:"sku"`
	Stock     int       `json:"stock"`
	Position  int       `json:"position"`
	ProductID uint      `json:"product_id"`
	IsDefault bool      `json:"is_default"`
	Option1   string    `json:"option1"`
	Option2   string    `json:"option2"`
	Option3   string    `json:"option3"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (vs *VariantSerializer) Response() VariantResponse {
	response := VariantResponse{
		ID:        vs.ID,
		Price:     vs.Price,
		Sku:       vs.Sku,
		Stock:     vs.Stock,
		Position:  vs.Position,
		ProductID: vs.ProductID,
		IsDefault: vs.IsDefault,
		Option1:   vs.Option1,
		Option2:   vs.Option2,
		Option3:   vs.Option3,
		CreatedAt: vs.CreatedAt,
		UpdatedAt: vs.UpdatedAt,
	}

	return response
}

func (s *VariantsSerializer) Response() []VariantResponse {
	response := []VariantResponse{}
	for _, variant := range s.Variants {
		serializer := VariantSerializer{variant}
		response = append(response, serializer.Response())
	}

	return response
}
