package service

import (
	"context"
	"github.com/igntnk/stocky_iims/models"
	"github.com/igntnk/stocky_iims/proto/pb"
	"github.com/igntnk/stocky_iims/repository"
	"github.com/rs/zerolog"
	"strconv"
	"time"
)

type ProductService interface {
	InsertOne(context.Context, *pb.InsertProductRequest) (*pb.InsertProductResponse, error)
	Get(context.Context, *pb.GetProductsRequest) (*pb.GetProductsResponse, error)
	Delete(context.Context, *pb.DeleteProductRequest) error
	Update(context.Context, *pb.UpdateProductRequest) error
	BlockProduct(context.Context, *pb.BlockProductOperationMessage) error
	UnblockProduct(context.Context, *pb.BlockProductOperationMessage) error
}

type productService struct {
	Logger zerolog.Logger
	repo   repository.ProductRepository
}

func NewProductService(logger zerolog.Logger, repo repository.ProductRepository) ProductService {
	return &productService{
		Logger: logger,
		repo:   repo,
	}
}

func (p productService) InsertOne(ctx context.Context, request *pb.InsertProductRequest) (*pb.InsertProductResponse, error) {
	price := strconv.FormatFloat(float64(request.Price), 'g', -1, 64)

	id, err := p.repo.InsertOne(ctx, &models.Product{
		Name:         request.Name,
		Description:  request.Description,
		Price:        price,
		CreationDate: time.Now().String(),
	})
	if err != nil {
		return nil, err
	}

	return &pb.InsertProductResponse{Id: id}, nil
}

func (p productService) Get(ctx context.Context, request *pb.GetProductsRequest) (*pb.GetProductsResponse, error) {
	products, err := p.repo.Get(ctx, request.GetLimit(), request.GetOffset())
	if err != nil {
		return nil, err
	}

	productsMessage := []*pb.GetProductMessage{}
	for _, product := range products {
		productsMessage = append(productsMessage, &pb.GetProductMessage{
			Id:           product.Id,
			Name:         product.Name,
			Description:  product.Description,
			CreationDate: product.CreationDate,
			Price:        product.Price,
		})
	}

	return &pb.GetProductsResponse{
		Products: productsMessage,
	}, nil
}

func (p productService) Delete(ctx context.Context, request *pb.DeleteProductRequest) error {
	return p.repo.Delete(ctx, request.GetId())
}

func (p productService) Update(ctx context.Context, request *pb.UpdateProductRequest) error {
	price := strconv.FormatFloat(float64(request.Price), 'g', -1, 64)

	return p.repo.Update(ctx, &models.Product{
		Id:           request.Id,
		Name:         request.Name,
		Description:  request.Description,
		Price:        price,
		CreationDate: time.Now().String(),
	})
}

func (p productService) BlockProduct(ctx context.Context, message *pb.BlockProductOperationMessage) error {
	return p.repo.BlockProduct(ctx, message.Id)
}

func (p productService) UnblockProduct(ctx context.Context, message *pb.BlockProductOperationMessage) error {
	return p.repo.UnblockProduct(ctx, message.Id)
}
