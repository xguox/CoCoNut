package model

type Variant struct {
	Price    float32
	Sku      string
	Stock    int
	Position int
	Product  Product
	Option1  *string
	Option2  *string
	Option3  *string
}
