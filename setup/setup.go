package setup

import (
	"context"
	"github.com/igntnk/stocky_iims/config"
	grpcapp "github.com/igntnk/stocky_iims/grpc"
	mongorepo "github.com/igntnk/stocky_iims/repository/mongo"
	"github.com/igntnk/stocky_iims/service"
	"github.com/rs/zerolog"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

var grpcServer *grpc.Server

func GRPCServer() *grpc.Server {
	return grpcServer
}

func Init(ctx context.Context, db *mongo.Database, isReplicaSet bool, logger zerolog.Logger, cfg *config.Config) error {
	var (
		saleRepo    = mongorepo.NewSaleRepository(ctx, db, isReplicaSet, logger)
		productRepo = mongorepo.NewProductRepository(ctx, db, isReplicaSet, logger)

		saleService    = service.NewSaleService(logger, saleRepo)
		productService = service.NewProductService(logger, productRepo)
	)

	grpcServer = grpc.NewServer()
	grpcapp.RegisterSaleServer(grpcServer, logger, saleService)
	grpcapp.RegisterProductServer(grpcServer, logger, productService)

	return nil
}
