package repository

import (
	"context"
	"github.com/igntnk/stocky_iims/models"
)

const (
	ProductCollection = "products"
)

type ProductRepository interface {
	InsertOne(context.Context, *models.Product) (string, error)
	Get(context.Context, int64, int64) ([]models.Product, error)
	GetById(context.Context, string) (models.Product, error)
	GetByProductCode(context.Context, string) (models.Product, error)
	Delete(context.Context, string) error
	Update(context.Context, *models.Product) error
	BlockProduct(context.Context, string) error
	UnblockProduct(context.Context, string) error
}
