package model

type Variant struct {
	Price    float32
	Sku      string
	Stock    int "stock"
	Position int "position"
	Product  Product
	Option1  *string "option1"
	Option2  *string "option2"
	Option3  *string "option3"
}
