package setup

import (
	"context"
	"github.com/igntnk/stocky_iims/config"
	grpcapp "github.com/igntnk/stocky_iims/grpc"
	"github.com/igntnk/stocky_iims/repository"
	mongorepo "github.com/igntnk/stocky_iims/repository/mongo"
	"github.com/igntnk/stocky_iims/service"
	"github.com/rs/zerolog"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

var grpcServer *grpc.Server

func GRPCServer() *grpc.Server {
	return grpcServer
}

func SetupDefaultData(ctx context.Context, db *mongo.Database) error {
	saleRepo := db.Collection(repository.SaleCollection)

	_, err := saleRepo.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys:    bson.D{{Key: " name", Value: 1}},
		Options: options.Index().SetUnique(true),
	})
	if err != nil {
		return err
	}

	productRepo := db.Collection(repository.ProductCollection)
	_, err = productRepo.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys:    bson.D{{Key: " name", Value: 1}},
		Options: options.Index().SetUnique(true),
	})

	if err != nil {
		return err
	}
	return nil
}

func Init(ctx context.Context, db *mongo.Database, isReplicaSet bool, logger zerolog.Logger, cfg *config.Config) error {
	var (
		saleRepo    = mongorepo.NewSaleRepository(ctx, db, isReplicaSet, logger)
		productRepo = mongorepo.NewProductRepository(ctx, db, isReplicaSet, logger)

		saleService    = service.NewSaleService(logger, saleRepo)
		productService = service.NewProductService(logger, productRepo)
	)

	grpcapp.RegisterSaleServer(grpcServer, logger, saleService)
	grpcapp.RegisterProductServer(grpcServer, logger, productService)

	return nil
}
