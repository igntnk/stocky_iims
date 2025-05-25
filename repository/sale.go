package repository

import (
	"context"
	"github.com/igntnk/stocky_iims/models"
)

const (
	SaleCollection = "products"
)

type SaleRepository interface {
	InsertOne(context.Context, *models.Sale) (string, error)
	Get(context.Context, int64, int64) ([]models.Sale, error)
	Delete(context.Context, string) error
	Update(context.Context, *models.Sale) error
	BlockSale(context.Context, string) error
	UnblockSale(context.Context, string) error
}
