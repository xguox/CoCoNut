package serializer

import (
	"coconut/model"
	"time"
)

type ProductSerializer struct {
	model.Product
}

type ProductsSerializer struct {
	Products []model.Product
}

type ProductResponse struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (s *ProductSerializer) Response() ProductResponse {
	response := ProductResponse{
		ID:        s.ID,
		Title:     s.Title,
		CreatedAt: s.CreatedAt,
		UpdatedAt: s.UpdatedAt,
	}
	return response
}

func (s *ProductsSerializer) Response() []ProductResponse {
	response := []ProductResponse{}
	for _, product := range s.Products {
		serializer := ProductSerializer{product}
		response = append(response, serializer.Response())
	}
	return response
}
