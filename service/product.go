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
	GetById(context.Context, *pb.GetByIdProductRequest) (*pb.GetProductMessage, error)
	GetByProductCode(context.Context, *pb.GetByProductCodeRequest) (*pb.GetProductMessage, error)
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
	id, err := p.repo.InsertOne(ctx, &models.Product{
		Name:         request.Name,
		Description:  request.Description,
		Price:        float64(request.Price),
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
		price := strconv.FormatFloat(product.Price, 'f', -1, 64)

		productsMessage = append(productsMessage, &pb.GetProductMessage{
			Id:           product.Id,
			Name:         product.Name,
			Description:  product.Description,
			CreationDate: product.CreationDate,
			Price:        price,
		})
	}

	return &pb.GetProductsResponse{
		Products: productsMessage,
	}, nil
}

func (p productService) GetById(ctx context.Context, request *pb.GetByIdProductRequest) (*pb.GetProductMessage, error) {
	res, err := p.repo.GetById(ctx, request.GetId())
	if err != nil {
		return nil, err
	}
	price := strconv.FormatFloat(res.Price, 'f', -1, 64)

	return &pb.GetProductMessage{
		Id:           res.Id,
		Name:         res.Name,
		Description:  res.Description,
		CreationDate: res.CreationDate,
		Price:        price,
	}, nil
}

func (p productService) GetByProductCode(ctx context.Context, request *pb.GetByProductCodeRequest) (*pb.GetProductMessage, error) {
	res, err := p.repo.GetByProductCode(ctx, request.GetCode())
	if err != nil {
		return nil, err
	}
	price := strconv.FormatFloat(res.Price, 'f', -1, 64)

	return &pb.GetProductMessage{
		Id:           res.Id,
		Name:         res.Name,
		Description:  res.Description,
		CreationDate: res.CreationDate,
		Price:        price,
	}, nil
}

func (p productService) Delete(ctx context.Context, request *pb.DeleteProductRequest) error {
	return p.repo.Delete(ctx, request.GetId())
}

func (p productService) Update(ctx context.Context, request *pb.UpdateProductRequest) error {
	return p.repo.Update(ctx, &models.Product{
		Id:           request.Id,
		Name:         request.Name,
		Description:  request.Description,
		Price:        float64(request.Price),
		CreationDate: time.Now().String(),
	})
}

func (p productService) BlockProduct(ctx context.Context, message *pb.BlockProductOperationMessage) error {
	return p.repo.BlockProduct(ctx, message.Id)
}

func (p productService) UnblockProduct(ctx context.Context, message *pb.BlockProductOperationMessage) error {
	return p.repo.UnblockProduct(ctx, message.Id)
}
