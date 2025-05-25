package models

type Sale struct {
	Id          string `json:"id" bson:"id"`
	Name        string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`
	SaleSize    int    `json:"sale_size" bson:"sale_size"`
	ProductId   string `json:"product_id" bson:"product_id"`
}
