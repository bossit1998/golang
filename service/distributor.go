package service

import (
	"context"
	"database/sql"

	gpb "github.com/golang/protobuf/ptypes/empty"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "bitbucket.org/alien_soft/courier_service/genproto/courier_service"
	l "bitbucket.org/alien_soft/courier_service/pkg/logger"
	"bitbucket.org/alien_soft/courier_service/service/grpc_client"
	"bitbucket.org/alien_soft/courier_service/storage"
)

// EventService ...
type DistributorService struct {
	storage storage.StorageI
	logger  l.Logger
	client  *grpc_client.GrpcClient
}

// NewEventService ...
func NewDistributorService(db *sqlx.DB, client *grpc_client.GrpcClient, log l.Logger) *DistributorService {
	return &DistributorService{
		storage: storage.NewStoragePg(db),
		logger:  log,
		client:  client,
	}
}

// Create ...
func (s *DistributorService) Create(ctx context.Context, req *pb.Distributor) (*pb.CreateDistributorResponse, error) {
	var err error

	Distributor, err := s.storage.Distributor().Create(req)
	if err != nil {
		s.logger.Error("Error while creating event", l.Error(err), l.Any("req", req))
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.CreateDistributorResponse{
		Distributor: Distributor,
	}, nil
}

func (s *DistributorService) Update(ctx context.Context, req *pb.Distributor) (*pb.UpdateDistributorResponse, error) {
	Distributor, err := s.storage.Distributor().Update(req)

	if err == sql.ErrNoRows {
		s.logger.Error("Error while updating event, Not Found", l.Any("req", req))
		return nil, status.Error(codes.NotFound, "Not found")
	} else if err != nil {
		s.logger.Error("Error while updating event", l.Error(err), l.Any("req", req))
		return nil, status.Error(codes.Internal, "Internal server error")
	}

	return &pb.UpdateDistributorResponse{
		Distributor: Distributor,
	}, nil
}

func (s *DistributorService) GetDistributor(ctx context.Context, req *pb.GetDistributorRequest) (*pb.GetDistributorResponse, error) {
	var Distributor *pb.Distributor
	Distributor, err := s.storage.Distributor().GetDistributor(req.Id)
	if err == sql.ErrNoRows {
		s.logger.Error("Error while getting an event, Not found", l.Any("req", req))
		return nil, status.Error(codes.NotFound, "Not found")
	} else if err != nil {
		s.logger.Error("Error while getting event", l.Error(err), l.Any("req", req))
		return nil, status.Error(codes.Internal, "Internal server error")
	}

	return &pb.GetDistributorResponse{
		Distributor: Distributor,
	}, nil
}

func (s *DistributorService) GetAllDistributors(ctx context.Context, req *pb.GetAllDistributorsRequest) (*pb.GetAllDistributorsResponse, error) {
	var Distributors []*pb.Distributor

	Distributors, count, err := s.storage.Distributor().GetAllDistributors(req.Page, req.Limit)
	if err == sql.ErrNoRows {
		s.logger.Error("Error while getting all events, Not found", l.Any("req", req))
		return nil, status.Error(codes.NotFound, "Not found")
	} else if err != nil {
		s.logger.Error("Error while getting all events", l.Error(err), l.Any("req", req))
		return nil, status.Error(codes.Internal, "Internal server error")
	}

	return &pb.GetAllDistributorsResponse{
		Distributors: Distributors,
		Count:        count,
	}, nil
}

func (s *DistributorService) Delete(ctx context.Context, req *pb.DeleteDistributorRequest) (*gpb.Empty, error) {
	err := s.storage.Distributor().Delete(req.Id)
	if err == sql.ErrNoRows {
		s.logger.Error("Error while deleting event, Not found", l.Any("req", req))
		return nil, status.Error(codes.NotFound, "Not found")
	} else if err != nil {
		s.logger.Error("Error while deleting event", l.Error(err), l.Any("req", req))
		return nil, status.Error(codes.Internal, "Internal server error")
	}
	return &gpb.Empty{}, nil
}
