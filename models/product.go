package models

type Product struct {
	Id           string  `json:"id" bson:"id"`
	Name         string  `json:"name" bson:"name"`
	Description  string  `json:"description" bson:"description"`
	Price        float64 `json:"price" bson:"price"`
	CreationDate string  `json:"creation_date" bson:"creation_date"`
}
