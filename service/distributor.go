package service

import (
	"context"
	"database/sql"

	gpb "github.com/golang/protobuf/ptypes/empty"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "genproto/courier_service"

	l "bitbucket.org/alien_soft/courier_service/pkg/logger"
	"bitbucket.org/alien_soft/courier_service/storage"
)

// Distributor Service ...
type DistributorService struct {
	storage storage.StorageI
	logger  l.Logger
}

// New Distributor Service ...
func NewDistributorService(db *sqlx.DB, log l.Logger) *DistributorService {
	return &DistributorService{
		storage: storage.NewStoragePg(db),
		logger:  log,
	}
}

//Create ...
func (s *DistributorService) Create(ctx context.Context, req *pb.Distributor) (*pb.CreateDistributorResponse, error) {
	var err error

	Distributor, err := s.storage.Distributor().Create(req)
	if err != nil {
		s.logger.Error("Error while creating distributor", l.Error(err), l.Any("req", req))
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.CreateDistributorResponse{
		Distributor: Distributor,
	}, nil
}

func (s *DistributorService) Update(ctx context.Context, req *pb.Distributor) (*pb.UpdateDistributorResponse, error) {
	Distributor, err := s.storage.Distributor().Update(req)

	if err == sql.ErrNoRows {
		s.logger.Error("Error while updating distributor, Not Found", l.Any("req", req))
		return nil, status.Error(codes.NotFound, "Not found")
	} else if err != nil {
		s.logger.Error("Error while updating distributor", l.Error(err), l.Any("req", req))
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
		s.logger.Error("Error while getting an distributor, Not found", l.Any("req", req))
		return nil, status.Error(codes.NotFound, "Not found")
	} else if err != nil {
		s.logger.Error("Error while getting distributor", l.Error(err), l.Any("req", req))
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
		s.logger.Error("Error while getting all distributors, Not found", l.Any("req", req))
		return nil, status.Error(codes.NotFound, "Not found")
	} else if err != nil {
		s.logger.Error("Error while getting all distributors", l.Error(err), l.Any("req", req))
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
		s.logger.Error("Error while deleting distributor, Not found", l.Any("req", req))
		return nil, status.Error(codes.NotFound, "Not found")
	} else if err != nil {
		s.logger.Error("Error while deleting distributor", l.Error(err), l.Any("req", req))
		return nil, status.Error(codes.Internal, "Internal server error")
	}
	return &gpb.Empty{}, nil
}

// Park
func (s *DistributorService) CreatePark(ctx context.Context, req *pb.Park) (*pb.CreateParkResponse, error) {
	park, err := s.storage.Distributor().CreatePark(req)
	if err != nil {
		s.logger.Error("Error while creating distributor", l.Error(err), l.Any("req", req))
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.CreateParkResponse{
		Park: park,
	}, nil
}

func (s *DistributorService) UpdatePark(ctx context.Context, req *pb.Park) (*pb.UpdateParkResponse, error) {
	Park, err := s.storage.Distributor().UpdatePark(req)

	if err == sql.ErrNoRows {
		s.logger.Error("Error while updating park, Not Found", l.Any("req", req))
		return nil, status.Error(codes.NotFound, "Not found")
	} else if err != nil {
		s.logger.Error("Error while updating park", l.Error(err), l.Any("req", req))
		return nil, status.Error(codes.Internal, "Internal server error")
	}

	return &pb.UpdateParkResponse{
		Park: Park,
	}, nil
}

func (s *DistributorService) GetPark(ctx context.Context, req *pb.GetParkRequest) (*pb.GetParkResponse, error) {
	var Park *pb.Park
	Park, err := s.storage.Distributor().GetPark(req.Id)
	if err == sql.ErrNoRows {
		s.logger.Error("Error while getting a park, Not found", l.Any("req", req))
		return nil, status.Error(codes.NotFound, "Not found")
	} else if err != nil {
		s.logger.Error("Error while getting a park", l.Error(err), l.Any("req", req))
		return nil, status.Error(codes.Internal, "Internal server error")
	}

	return &pb.GetParkResponse{
		Park: Park,
	}, nil
}

func (s *DistributorService) GetAllDistributorParks(ctx context.Context, req *pb.GetAllDistributorParksRequest) (*pb.GetAllDistributorParksResponse, error) {
	var Parks []*pb.Park

	Parks, count, err := s.storage.Distributor().GetAllDistributorParks(req.Page, req.Limit)
	if err == sql.ErrNoRows {
		s.logger.Error("Error while getting all distributors, Not found", l.Any("req", req))
		return nil, status.Error(codes.NotFound, "Not found")
	} else if err != nil {
		s.logger.Error("Error while getting all distributors", l.Error(err), l.Any("req", req))
		return nil, status.Error(codes.Internal, "Internal server error")
	}

	return &pb.GetAllDistributorParksResponse{
		Parks: Parks,
		Count: count,
	}, nil
}

func (s *DistributorService) GetAllParks(ctx context.Context, req *pb.GetAllParksRequest) (*pb.GetAllParksResponse, error) {
	// var Distributors []*pb.Distributor

	// Distributors, count, err := s.storage.Distributor().GetAllDistributors(req.Page, req.Limit)
	// if err == sql.ErrNoRows {
	// 	s.logger.Error("Error while getting all distributors, Not found", l.Any("req", req))
	// 	return nil, status.Error(codes.NotFound, "Not found")
	// } else if err != nil {
	// 	s.logger.Error("Error while getting all distributors", l.Error(err), l.Any("req", req))
	// 	return nil, status.Error(codes.Internal, "Internal server error")
	// }

	// return &pb.GetAllDistributorsResponse{
	// 	Distributors: Distributors,
	// 	Count:        count,
	// }, nil
	return nil, nil
}

func (s *DistributorService) DeletePark(ctx context.Context, req *pb.DeleteParkRequest) (*gpb.Empty, error) {
	err := s.storage.Distributor().DeletePark(req.Id)
	if err == sql.ErrNoRows {
		s.logger.Error("Error while deleting park, Not found", l.Any("req", req))
		return nil, status.Error(codes.NotFound, "Not found")
	} else if err != nil {
		s.logger.Error("Error while deleting park", l.Error(err), l.Any("req", req))
		return nil, status.Error(codes.Internal, "Internal server error")
	}
	return &gpb.Empty{}, nil
}
