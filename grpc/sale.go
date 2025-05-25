package grpc

import (
	"context"
	iims_pb "github.com/igntnk/stocky_iims/proto/pb"
	"github.com/igntnk/stocky_iims/service"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type saleServer struct {
	iims_pb.UnimplementedSaleServiceServer
	Logger      zerolog.Logger
	SaleService service.SaleService
}

func RegisterSaleServer(server *grpc.Server, logger zerolog.Logger, saleService service.SaleService) {
	iims_pb.RegisterSaleServiceServer(server, &saleServer{Logger: logger, SaleService: saleService})
}

func (s *saleServer) InsertOne(ctx context.Context, req *iims_pb.InsertSaleRequest) (*iims_pb.InsertSaleResponse, error) {
	s.Logger.Debug().Msg("Insert Sale")

	result, err := s.SaleService.InsertOne(ctx, req)
	if err != nil {
		s.Logger.Error().Err(err).Msg("Sale Insert Error")
		return nil, err
	}

	return result, nil
}

func (s *saleServer) Get(ctx context.Context, req *iims_pb.GetSalesRequest) (*iims_pb.GetSalesResponse, error) {
	s.Logger.Debug().Msg("Get Sale")

	result, err := s.SaleService.Get(ctx, req)
	if err != nil {
		s.Logger.Error().Err(err).Msg("SaleService Get error")
		return nil, err
	}

	return result, nil
}

func (s *saleServer) Delete(ctx context.Context, req *iims_pb.DeleteSaleRequest) (*emptypb.Empty, error) {
	s.Logger.Debug().Msg("Delete Sale")

	err := s.SaleService.Delete(ctx, req)
	if err != nil {
		s.Logger.Error().Err(err).Msg("SaleService Delete error")
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *saleServer) Update(ctx context.Context, req *iims_pb.UpdateSaleRequest) (*emptypb.Empty, error) {
	s.Logger.Debug().Msg("Update Sale")

	err := s.SaleService.Update(ctx, req)
	if err != nil {
		s.Logger.Error().Err(err).Msg("SaleService Update error")
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *saleServer) BlockSale(ctx context.Context, req *iims_pb.BlockSaleOperationMessage) (*emptypb.Empty, error) {
	s.Logger.Debug().Msg("Block Sale")

	err := s.SaleService.BlockSale(ctx, req)
	if err != nil {
		s.Logger.Error().Err(err).Msg("SaleService BlockSale error")
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *saleServer) UnblockSale(ctx context.Context, req *iims_pb.BlockSaleOperationMessage) (*emptypb.Empty, error) {
	s.Logger.Debug().Msg("Unblock Sale")

	err := s.SaleService.UnblockSale(ctx, req)
	if err != nil {
		s.Logger.Error().Err(err).Msg("SaleService UnblockSale error")
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
