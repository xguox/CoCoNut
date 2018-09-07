package model

type Variant struct {
	Price    float32
	Sku      string
	Stock    int `json:"stock"`
	Position int `json:"position"`
	Product  Product
	Option1  *string `json:"option1"`
	Option2  *string `json:"option2"`
	Option3  *string `json:"option3"`
}
