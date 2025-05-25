package service

import (
	"context"
	"github.com/igntnk/stocky_iims/models"
	"github.com/igntnk/stocky_iims/proto/pb"
	"github.com/igntnk/stocky_iims/repository"
	"github.com/rs/zerolog"
)

type SaleService interface {
	InsertOne(context.Context, *pb.InsertSaleRequest) (*pb.InsertSaleResponse, error)
	Get(context.Context, *pb.GetSalesRequest) (*pb.GetSalesResponse, error)
	Delete(context.Context, *pb.DeleteSaleRequest) error
	Update(context.Context, *pb.UpdateSaleRequest) error
	BlockSale(context.Context, *pb.BlockSaleOperationMessage) error
	UnblockSale(context.Context, *pb.BlockSaleOperationMessage) error
}

type saleService struct {
	Logger zerolog.Logger
	repo   repository.SaleRepository
}

func NewSaleService(logger zerolog.Logger, repo repository.SaleRepository) SaleService {
	return &saleService{
		Logger: logger,
		repo:   repo,
	}
}

func (s saleService) InsertOne(ctx context.Context, request *pb.InsertSaleRequest) (*pb.InsertSaleResponse, error) {
	result, err := s.repo.InsertOne(ctx, &models.Sale{
		Name:        request.Name,
		Description: request.Description,
		SaleSize:    int(request.SaleSize),
		ProductId:   request.Product,
	})
	if err != nil {
		return nil, err
	}

	return &pb.InsertSaleResponse{
		Id: result,
	}, nil
}

func (s saleService) Get(ctx context.Context, request *pb.GetSalesRequest) (*pb.GetSalesResponse, error) {
	sales, err := s.repo.Get(ctx, request.GetLimit(), request.GetOffset())
	if err != nil {
		return nil, err
	}

	resultSales := make([]*pb.GetSaleMessage, len(sales))

	for i, sale := range sales {
		resultSales[i] = &pb.GetSaleMessage{
			Id:          sale.Id,
			Name:        sale.Name,
			Description: sale.Description,
			SaleSize:    int32(sale.SaleSize),
			Product:     sale.ProductId,
		}
	}

	return &pb.GetSalesResponse{
		Sales: resultSales,
	}, nil
}

func (s saleService) Delete(ctx context.Context, request *pb.DeleteSaleRequest) error {
	return s.repo.Delete(ctx, request.GetId())
}

func (s saleService) Update(ctx context.Context, request *pb.UpdateSaleRequest) error {
	return s.repo.Update(ctx, &models.Sale{
		Id:          request.GetId(),
		Name:        request.GetName(),
		Description: request.GetDescription(),
		SaleSize:    int(request.SaleSize),
	})
}

func (s saleService) BlockSale(ctx context.Context, message *pb.BlockSaleOperationMessage) error {
	return s.repo.BlockSale(ctx, message.Id)
}

func (s saleService) UnblockSale(ctx context.Context, message *pb.BlockSaleOperationMessage) error {
	return s.repo.UnblockSale(ctx, message.Id)
}
