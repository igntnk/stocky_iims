package grpc

import (
	"context"
	iims_pb "github.com/igntnk/stocky_iims/proto/pb"
	"github.com/igntnk/stocky_iims/service"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type productServer struct {
	iims_pb.UnimplementedProductServiceServer
	Logger         zerolog.Logger
	ProductService service.ProductService
}

func RegisterProductServer(server *grpc.Server, logger zerolog.Logger, productService service.ProductService) {
	iims_pb.RegisterProductServiceServer(server, &productServer{Logger: logger, ProductService: productService})
}

func (s *productServer) InsertOne(ctx context.Context, req *iims_pb.InsertProductRequest) (*iims_pb.InsertProductResponse, error) {
	s.Logger.Debug().Msg("Insert Product")

	result, err := s.ProductService.InsertOne(ctx, req)
	if err != nil {
		s.Logger.Error().Err(err).Msg("ProductService InsertOne error")
		return nil, err
	}

	return result, nil
}

func (s *productServer) Get(ctx context.Context, req *iims_pb.GetProductsRequest) (*iims_pb.GetProductsResponse, error) {
	s.Logger.Debug().Msg("Get Product")

	result, err := s.ProductService.Get(ctx, req)
	if err != nil {
		s.Logger.Error().Err(err).Msg("ProductService Get error")
		return nil, err
	}

	return result, nil
}

func (s *productServer) GetById(ctx context.Context, req *iims_pb.GetByIdProductRequest) (*iims_pb.GetProductMessage, error) {
	s.Logger.Debug().Msg("Get Product")

	result, err := s.ProductService.GetById(ctx, req)
	if err != nil {
		s.Logger.Error().Err(err).Msg("ProductService GetById error")
		return nil, err
	}

	return result, nil
}

func (s *productServer) GetByProductCode(ctx context.Context, req *iims_pb.GetByProductCodeRequest) (*iims_pb.GetProductMessage, error) {
	s.Logger.Debug().Msg("Get Product")

	result, err := s.ProductService.GetByProductCode(ctx, req)
	if err != nil {
		s.Logger.Error().Err(err).Msg("ProductService GetById error")
		return nil, err
	}

	return result, nil
}

func (s *productServer) Delete(ctx context.Context, req *iims_pb.DeleteProductRequest) (*emptypb.Empty, error) {
	s.Logger.Debug().Msg("Delete Product")

	err := s.ProductService.Delete(ctx, req)
	if err != nil {
		s.Logger.Error().Err(err).Msg("ProductService Delete error")
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *productServer) Update(ctx context.Context, req *iims_pb.UpdateProductRequest) (*emptypb.Empty, error) {
	s.Logger.Debug().Msg("Update Product")

	err := s.ProductService.Update(ctx, req)
	if err != nil {
		s.Logger.Error().Err(err).Msg("ProductService Update error")
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *productServer) BlockProduct(ctx context.Context, req *iims_pb.BlockProductOperationMessage) (*emptypb.Empty, error) {
	s.Logger.Debug().Msg("Block Product")

	err := s.ProductService.BlockProduct(ctx, req)
	if err != nil {
		s.Logger.Error().Err(err).Msg("ProductService BlockProduct error")
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *productServer) UnblockProduct(ctx context.Context, req *iims_pb.BlockProductOperationMessage) (*emptypb.Empty, error) {
	s.Logger.Debug().Msg("Unblock Product")

	err := s.ProductService.UnblockProduct(ctx, req)
	if err != nil {
		s.Logger.Error().Err(err).Msg("ProductService UnblockProduct error")
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
