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

// CourierService ...
type CourierService struct {
	storage storage.StorageI
	logger  l.Logger
	client  *grpc_client.GrpcClient
}

// NewCourierService ...
func NewCourierService(db *sqlx.DB, client *grpc_client.GrpcClient, log l.Logger) *CourierService {
	return &CourierService{
		storage: storage.NewStoragePg(db),
		logger:  log,
		client:  client,
	}
}

// Courier
func (s *CourierService) Create(ctx context.Context, req *pb.Courier) (*pb.CreateCourierResponse, error) {
	var err error

	courier, err := s.storage.Courier().Create(req)
	if err != nil {
		s.logger.Error("Error while creating courier", l.Error(err), l.Any("req", req))
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.CreateCourierResponse{
		Courier: courier,
	}, nil
}

func (s *CourierService) Update(ctx context.Context, req *pb.Courier) (*pb.UpdateCourierResponse, error) {
	courier, err := s.storage.Courier().Update(req)

	if err == sql.ErrNoRows {
		s.logger.Error("Error while updating courier, Not Found", l.Any("req", req))
		return nil, status.Error(codes.NotFound, "Not found")
	} else if err != nil {
		s.logger.Error("Error while updating courier", l.Error(err), l.Any("req", req))
		return nil, status.Error(codes.Internal, "Internal server error")
	}

	return &pb.UpdateCourierResponse{
		Courier: courier,
	}, nil
}

func (s *CourierService) GetCourier(ctx context.Context, req *pb.GetCourierRequest) (*pb.GetCourierResponse, error) {
	var courier *pb.Courier
	courier, err := s.storage.Courier().GetCourier(req.Id)
	if err == sql.ErrNoRows {
		s.logger.Error("Error while getting an courier, Not found", l.Any("req", req))
		return nil, status.Error(codes.NotFound, "Not found")
	} else if err != nil {
		s.logger.Error("Error while getting courier", l.Error(err), l.Any("req", req))
		return nil, status.Error(codes.Internal, "Internal server error")
	}

	return &pb.GetCourierResponse{
		Courier: courier,
	}, nil
}

func (s *CourierService) GetAllCouriers(ctx context.Context, req *pb.GetAllCouriersRequest) (*pb.GetAllCouriersResponse, error) {
	var couriers []*pb.Courier

	couriers, count, err := s.storage.Courier().GetAllCouriers(req.Page, req.Limit)
	if err == sql.ErrNoRows {
		s.logger.Error("Error while getting all couriers, Not found", l.Any("req", req))
		return nil, status.Error(codes.NotFound, "Not found")
	} else if err != nil {
		s.logger.Error("Error while getting all couriers", l.Error(err), l.Any("req", req))
		return nil, status.Error(codes.Internal, "Internal server error")
	}

	return &pb.GetAllCouriersResponse{
		Couriers: couriers,
		Count:    count,
	}, nil
}

func (s *CourierService) Delete(ctx context.Context, req *pb.DeleteCourierRequest) (*gpb.Empty, error) {
	err := s.storage.Courier().Delete(req.Id)
	if err == sql.ErrNoRows {
		s.logger.Error("Error while deleting courier, Not found", l.Any("req", req))
		return nil, status.Error(codes.NotFound, "Not found")
	} else if err != nil {
		s.logger.Error("Error while deleting courier", l.Error(err), l.Any("req", req))
		return nil, status.Error(codes.Internal, "Internal server error")
	}
	return &gpb.Empty{}, nil
}

func (s *CourierService) BlockCourier(ctx context.Context, req *pb.BlockCourierRequest) (*gpb.Empty, error) {
	err := s.storage.Courier().Delete(req.Id)
	if err == sql.ErrNoRows {
		s.logger.Error("Error while blocking courier, Not found", l.Any("req", req))
		return nil, status.Error(codes.NotFound, "Not found")
	} else if err != nil {
		s.logger.Error("Error while blocking courier", l.Error(err), l.Any("req", req))
		return nil, status.Error(codes.Internal, "Internal server error")
	}
	return &gpb.Empty{}, nil
}

func (s *CourierService) UnblockCourier(ctx context.Context, req *pb.UnblockCourierRequest) (*gpb.Empty, error) {
	err := s.storage.Courier().Delete(req.Id)
	if err == sql.ErrNoRows {
		s.logger.Error("Error while unblocking courier, Not found", l.Any("req", req))
		return nil, status.Error(codes.NotFound, "Not found")
	} else if err != nil {
		s.logger.Error("Error while unblocking courier", l.Error(err), l.Any("req", req))
		return nil, status.Error(codes.Internal, "Internal server error")
	}
	return &gpb.Empty{}, nil
}

func (s *CourierService) GetAllDistributorCouriers(ctx context.Context, req *pb.GetAllDistributorCouriersRequest) (*pb.GetAllDistributorCouriersResponse, error) {
	var couriers []*pb.Courier

	couriers, count, err := s.storage.Courier().GetAllDistributorCouriers(req.DistributorId, req.Page, req.Limit)
	if err == sql.ErrNoRows {
		s.logger.Error("Error while getting all distributor's couriers, Not found", l.Any("req", req))
		return nil, status.Error(codes.NotFound, "Not found")
	} else if err != nil {
		s.logger.Error("Error while getting all distrubutor's couriers", l.Error(err), l.Any("req", req))
		return nil, status.Error(codes.Internal, "Internal server error")
	}

	return &pb.GetAllDistributorCouriersResponse{
		Couriers: couriers,
		Count:    count,
	}, nil
}

// CourierDetails
func (s *CourierService) CreateCourierDetails(ctx context.Context, req *pb.CourierDetails) (*pb.CreateCourierDetailsResponse, error) {

	cd, err := s.storage.Courier().CreateCourierDetails(req)
	if err != nil {
		s.logger.Error("Error while creating courier details", l.Error(err), l.Any("req", req))
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.CreateCourierDetailsResponse{
		CourierDetails: cd,
	}, nil
}

func (s *CourierService) UpdateCourierDetails(ctx context.Context, req *pb.CourierDetails) (*pb.UpdateCourierDetailsResponse, error) {
	cd, err := s.storage.Courier().UpdateCourierDetails(req)

	if err == sql.ErrNoRows {
		s.logger.Error("Error while updating courier details, Not Found", l.Any("req", req))
		return nil, status.Error(codes.NotFound, "Not found")
	} else if err != nil {
		s.logger.Error("Error while updating courier details", l.Error(err), l.Any("req", req))
		return nil, status.Error(codes.Internal, "Internal server error")
	}

	return &pb.UpdateCourierDetailsResponse{
		CourierDetails: cd,
	}, nil
}

func (s *CourierService) GetCourierDetails(ctx context.Context, req *pb.GetCourierDetailsRequest) (*pb.GetCourierDetailsResponse, error) {
	var cd *pb.CourierDetails
	cd, err := s.storage.Courier().GetCourierDetails(req.CourierId)
	if err == sql.ErrNoRows {
		s.logger.Error("Error while getting an courier details, Not found", l.Any("req", req))
		return nil, status.Error(codes.NotFound, "Not found")
	} else if err != nil {
		s.logger.Error("Error while getting courier details", l.Error(err), l.Any("req", req))
		return nil, status.Error(codes.Internal, "Internal server error")
	}

	return &pb.GetCourierDetailsResponse{
		CourierDetails: cd,
	}, nil
}

// CourierVehicle
func (s *CourierService) CreateCourierVehicle(ctx context.Context, req *pb.CourierVehicle) (*pb.CreateCourierVehicleResponse, error) {
	cv, err := s.storage.Courier().CreateCourierVehicle(req)
	if err != nil {
		s.logger.Error("Error while creating courier vehicle", l.Error(err), l.Any("req", req))
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.CreateCourierVehicleResponse{
		CourierVehicle: cv,
	}, nil
}

func (s *CourierService) UpdateCourierVehicle(ctx context.Context, req *pb.CourierVehicle) (*pb.UpdateCourierVehicleResponse, error) {
	cv, err := s.storage.Courier().UpdateCourierVehicle(req)

	if err == sql.ErrNoRows {
		s.logger.Error("Error while updating courier vehicle, Not Found", l.Any("req", req))
		return nil, status.Error(codes.NotFound, "Not found")
	} else if err != nil {
		s.logger.Error("Error while updating courier vehicle", l.Error(err), l.Any("req", req))
		return nil, status.Error(codes.Internal, "Internal server error")
	}

	return &pb.UpdateCourierVehicleResponse{
		CourierVehicle: cv,
	}, nil
}

func (s *CourierService) GetCourierVehicle(ctx context.Context, req *pb.GetCourierVehicleRequest) (*pb.GetCourierVehicleResponse, error) {
	var cv *pb.CourierVehicle
	cv, err := s.storage.Courier().GetCourierVehicle(req.Id)
	if err == sql.ErrNoRows {
		s.logger.Error("Error while getting an courier vehicle, Not found", l.Any("req", req))
		return nil, status.Error(codes.NotFound, "Not found")
	} else if err != nil {
		s.logger.Error("Error while getting courier vehicle", l.Error(err), l.Any("req", req))
		return nil, status.Error(codes.Internal, "Internal server error")
	}

	return &pb.GetCourierVehicleResponse{
		CourierVehicle: cv,
	}, nil
}

func (s *CourierService) GetAllCourierVehicles(ctx context.Context, req *pb.GetAllCourierVehiclesRequest) (*pb.GetAllCourierVehiclesResponse, error) {
	var cv []*pb.CourierVehicle

	cv, err := s.storage.Courier().GetAllCourierVehicles(req.CourierId)
	if err == sql.ErrNoRows {
		s.logger.Error("Error while getting all courier vehicle, Not found", l.Any("req", req))
		return nil, status.Error(codes.NotFound, "Not found")
	} else if err != nil {
		s.logger.Error("Error while getting all courier vehicle", l.Error(err), l.Any("req", req))
		return nil, status.Error(codes.Internal, "Internal server error")
	}

	return &pb.GetAllCourierVehiclesResponse{
		CourierVehicles: cv,
	}, nil
}

func (s *CourierService) DeleteCourierVehicle(ctx context.Context, req *pb.DeleteCourierVehicleRequest) (*gpb.Empty, error) {
	err := s.storage.Courier().DeleteCourierVehicle(req.Id)
	if err == sql.ErrNoRows {
		s.logger.Error("Error while deleting courier vehicle, Not found", l.Any("req", req))
		return nil, status.Error(codes.NotFound, "Not found")
	} else if err != nil {
		s.logger.Error("Error while deleting courier vehicle", l.Error(err), l.Any("req", req))
		return nil, status.Error(codes.Internal, "Internal server error")
	}
	return &gpb.Empty{}, nil
}
