package serializer

import (
	"coconut/model"
	"time"
)

type OptionSerializer struct {
	model.Option
}

type OptionsSerializer struct {
	Options []model.Option
}

type OptionResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Position  int       `json:"position"`
	ProductID uint      `json:"product_id"`
	Values    []string  `json:"values"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (os *OptionSerializer) Response() OptionResponse {
	response := OptionResponse{
		ID:        os.ID,
		Name:      os.Name,
		Position:  os.Position,
		ProductID: os.ProductID,
		Values:    os.ValuesArr(),
		CreatedAt: os.CreatedAt,
		UpdatedAt: os.UpdatedAt,
	}

	return response
}

func (s *OptionsSerializer) Response() []OptionResponse {
	response := []OptionResponse{}
	for _, option := range s.Options {
		serializer := OptionSerializer{option}
		response = append(response, serializer.Response())
	}
	return response
}
