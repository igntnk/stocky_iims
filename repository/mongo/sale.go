package mongo

import (
	"context"
	"github.com/igntnk/stocky_iims/models"
	"github.com/igntnk/stocky_iims/repository"
	"github.com/rs/zerolog"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type saleRepository struct {
	Logger         zerolog.Logger
	SaleCollection *mongo.Collection
	Tx             Tx
}

func NewSaleRepository(ctx context.Context, database *mongo.Database, trxImpl bool, logger zerolog.Logger) repository.SaleRepository {
	tx := noTxImpl
	if trxImpl {
		tx = txImpl
	}

	return &saleRepository{
		Logger:         logger.With().Str("repository", repository.SaleCollection).Logger(),
		SaleCollection: database.Collection(repository.SaleCollection),
		Tx:             tx,
	}
}

func (r *saleRepository) InsertOne(ctx context.Context, sale *models.Sale) (string, error) {
	res, err := r.SaleCollection.InsertOne(ctx, sale)
	if err != nil {
		return "", err
	}

	return res.InsertedID.(string), nil
}

func (r *saleRepository) Get(ctx context.Context, limit, offset int64) ([]models.Sale, error) {
	sales := []models.Sale{}
	pipeline := getPipeline(limit, offset)
	res, err := r.SaleCollection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer res.Close(ctx)

	err = res.All(ctx, &sales)
	if err != nil {
		return nil, err
	}

	return sales, nil
}

func (r *saleRepository) Delete(ctx context.Context, id string) error {
	idObj, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = r.SaleCollection.DeleteOne(ctx, bson.M{"_id": idObj})
	if err != nil {
		return err
	}

	return nil
}

func (r *saleRepository) Update(ctx context.Context, Sale *models.Sale) error {
	id, err := primitive.ObjectIDFromHex(Sale.Id)
	if err != nil {
		return err
	}

	_, err = r.SaleCollection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": Sale})
	if err != nil {
		return err
	}

	return nil
}

func (r *saleRepository) BlockSale(ctx context.Context, id string) error {
	idObj, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = r.SaleCollection.UpdateOne(ctx, bson.M{"_id": idObj}, bson.M{"$set": bson.M{"blocked": true}})
	if err != nil {
		return err
	}

	return nil
}

func (r *saleRepository) UnblockSale(ctx context.Context, id string) error {
	idObj, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = r.SaleCollection.UpdateOne(ctx, bson.M{"_id": idObj}, bson.M{"$set": bson.M{"blocked": false}})
	if err != nil {
		return err
	}

	return nil
}
