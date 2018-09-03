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
	ID          uint       `json:"id"`
	Title       string     `json:"title"`
	BodyHTML    *string    `json:"body_html"`
	PublishedAt *time.Time `json:"published_at"`
	Vendor      *string    `json:"vendor"`
	Keywords    *string    `json:"keywords"`
	Price       float32    `json:"price"`
	Slug        *string    `json:"slug"`
	StockQty    int        `json:"stock_qty"`
	Status      int        `json:"status"`
	HotSale     bool       `json:"hot_sale"`
	NewArrival  bool       `json:"new_arrival"`
	CategoryID  int        `json:"category_id"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (s *ProductSerializer) Response() ProductResponse {
	response := ProductResponse{
		ID:          s.ID,
		Title:       s.Title,
		Price:       s.Price,
		BodyHTML:    s.BodyHTML,
		PublishedAt: s.PublishedAt,
		Vendor:      s.Vendor,
		Keywords:    s.Keywords,
		Slug:        s.Slug,
		StockQty:    s.StockQty,
		Status:      s.Status,
		HotSale:     s.HotSale,
		NewArrival:  s.NewArrival,
		CategoryID:  s.CategoryID,
		CreatedAt:   s.CreatedAt,
		UpdatedAt:   s.UpdatedAt,
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
