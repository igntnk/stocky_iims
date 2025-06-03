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

type productRepository struct {
	Logger            zerolog.Logger
	ProductCollection *mongo.Collection
	Tx                Tx
}

func NewProductRepository(ctx context.Context, database *mongo.Database, trxImpl bool, logger zerolog.Logger) repository.ProductRepository {
	tx := noTxImpl
	if trxImpl {
		tx = txImpl
	}

	return &productRepository{
		Logger:            logger.With().Str("repository", repository.ProductCollection).Logger(),
		ProductCollection: database.Collection(repository.ProductCollection),
		Tx:                tx,
	}
}

func getPipeline(limit, offset int64) mongo.Pipeline {
	pipeline := mongo.Pipeline{}
	if offset > 0 {
		pipeline = append(pipeline, bson.D{{
			"$skip",
			offset,
		}})
	}

	if limit > 0 {
		pipeline = append(pipeline, bson.D{{
			"$limit",
			limit,
		}})
	}
	return pipeline
}

func (r *productRepository) InsertOne(ctx context.Context, product *models.Product) (string, error) {
	res, err := r.ProductCollection.InsertOne(ctx, product)
	if err != nil {
		return "", err
	}

	return res.InsertedID.(string), nil
}

func (r *productRepository) Get(ctx context.Context, limit, offset int64) ([]models.Product, error) {
	products := []models.Product{}
	pipeline := getPipeline(limit, offset)
	res, err := r.ProductCollection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer res.Close(ctx)

	err = res.All(ctx, &products)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (r *productRepository) GetById(ctx context.Context, id string) (models.Product, error) {
	product := models.Product{}

	idObj, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return product, err
	}

	err = r.ProductCollection.FindOne(ctx, bson.M{"_id": idObj}).Decode(&product)
	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *productRepository) GetByProductCode(ctx context.Context, code string) (models.Product, error) {
	product := models.Product{}

	err := r.ProductCollection.FindOne(ctx, bson.M{"product_code": code}).Decode(&product)
	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *productRepository) Delete(ctx context.Context, id string) error {
	idObj, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = r.ProductCollection.DeleteOne(ctx, bson.M{"_id": idObj})
	if err != nil {
		return err
	}

	return nil
}

func (r *productRepository) Update(ctx context.Context, product *models.Product) error {
	id, err := primitive.ObjectIDFromHex(product.Id)
	if err != nil {
		return err
	}

	_, err = r.ProductCollection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": product})
	if err != nil {
		return err
	}

	return nil
}

func (r *productRepository) BlockProduct(ctx context.Context, id string) error {
	idObj, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = r.ProductCollection.UpdateOne(ctx, bson.M{"_id": idObj}, bson.M{"$set": bson.M{"blocked": true}})
	if err != nil {
		return err
	}

	return nil
}

func (r *productRepository) UnblockProduct(ctx context.Context, id string) error {
	idObj, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = r.ProductCollection.UpdateOne(ctx, bson.M{"_id": idObj}, bson.M{"$set": bson.M{"blocked": false}})
	if err != nil {
		return err
	}

	return nil
}
