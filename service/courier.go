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
	"bitbucket.org/alien_soft/courier_service/storage"
)

// CourierService ...
type CourierService struct {
	storage storage.StorageI
	logger  l.Logger
}

// NewCourierService ...
func NewCourierService(db *sqlx.DB, log l.Logger) *CourierService {
	return &CourierService{
		storage: storage.NewStoragePg(db),
		logger:  log,
	}
}

// Courier
func (s *CourierService) Create(ctx context.Context, req *pb.Courier) (*pb.CreateCourierResponse, error) {
	var err error

	courier, err := s.storage.Courier().Create(req)
	if err != nil {
		s.logger.Error("Error while creating event", l.Error(err), l.Any("req", req))
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.CreateCourierResponse{
		Courier: courier,
	}, nil
}

func (s *CourierService) Update(ctx context.Context, req *pb.Courier) (*pb.UpdateCourierResponse, error) {
	courier, err := s.storage.Courier().Update(req)

	if err == sql.ErrNoRows {
		s.logger.Error("Error while updating event, Not Found", l.Any("req", req))
		return nil, status.Error(codes.NotFound, "Not found")
	} else if err != nil {
		s.logger.Error("Error while updating event", l.Error(err), l.Any("req", req))
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
		s.logger.Error("Error while getting an event, Not found", l.Any("req", req))
		return nil, status.Error(codes.NotFound, "Not found")
	} else if err != nil {
		s.logger.Error("Error while getting event", l.Error(err), l.Any("req", req))
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
		s.logger.Error("Error while getting all events, Not found", l.Any("req", req))
		return nil, status.Error(codes.NotFound, "Not found")
	} else if err != nil {
		s.logger.Error("Error while getting all events", l.Error(err), l.Any("req", req))
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
		s.logger.Error("Error while deleting event, Not found", l.Any("req", req))
		return nil, status.Error(codes.NotFound, "Not found")
	} else if err != nil {
		s.logger.Error("Error while deleting event", l.Error(err), l.Any("req", req))
		return nil, status.Error(codes.Internal, "Internal server error")
	}
	return &gpb.Empty{}, nil
}

// CourierDetails
func (s *CourierService) CreateCourierDetails(ctx context.Context, req *pb.CourierDetails) (*pb.CreateCourierDetailsResponse, error) {
	// var err error

	// courier, err := s.storage.Courier().Create(req)
	// if err != nil {
	// 	s.logger.Error("Error while creating event", l.Error(err), l.Any("req", req))
	// 	return nil, status.Error(codes.Internal, err.Error())
	// }
	// return &pb.CreateCourierResponse{
	// 	Courier: courier,
	// }, nil
	return nil, nil
}

func (s *CourierService) UpdateCourierDetails(ctx context.Context, req *pb.CourierDetails) (*pb.UpdateCourierDetailsResponse, error) {
	// courier, err := s.storage.Courier().Update(req)

	// if err == sql.ErrNoRows {
	// 	s.logger.Error("Error while updating event, Not Found", l.Any("req", req))
	// 	return nil, status.Error(codes.NotFound, "Not found")
	// } else if err != nil {
	// 	s.logger.Error("Error while updating event", l.Error(err), l.Any("req", req))
	// 	return nil, status.Error(codes.Internal, "Internal server error")
	// }

	// return &pb.UpdateCourierResponse{
	// 	Courier: courier,
	// }, nil
	return nil, nil
}

func (s *CourierService) GetCourierDetails(ctx context.Context, req *pb.GetCourierDetailsRequest) (*pb.GetCourierDetailsResponse, error) {
	// var courier *pb.Courier
	// courier, err := s.storage.Courier().GetCourier(req.Id)
	// if err == sql.ErrNoRows {
	// 	s.logger.Error("Error while getting an event, Not found", l.Any("req", req))
	// 	return nil, status.Error(codes.NotFound, "Not found")
	// } else if err != nil {
	// 	s.logger.Error("Error while getting event", l.Error(err), l.Any("req", req))
	// 	return nil, status.Error(codes.Internal, "Internal server error")
	// }

	// return &pb.GetCourierResponse{
	// 	Courier: courier,
	// }, nil
	return nil, nil
}

// CourierVehicle
func (s *CourierService) CreateCourierVehicle(ctx context.Context, req *pb.CourierVehicle) (*pb.CreateCourierVehicleResponse, error) {
	// var err error

	// courier, err := s.storage.Courier().Create(req)
	// if err != nil {
	// 	s.logger.Error("Error while creating event", l.Error(err), l.Any("req", req))
	// 	return nil, status.Error(codes.Internal, err.Error())
	// }
	// return &pb.CreateCourierResponse{
	// 	Courier: courier,
	// }, nil
	return nil, nil
}

func (s *CourierService) UpdateCourierVehicle(ctx context.Context, req *pb.CourierVehicle) (*pb.UpdateCourierVehicleResponse, error) {
	// courier, err := s.storage.Courier().Update(req)

	// if err == sql.ErrNoRows {
	// 	s.logger.Error("Error while updating event, Not Found", l.Any("req", req))
	// 	return nil, status.Error(codes.NotFound, "Not found")
	// } else if err != nil {
	// 	s.logger.Error("Error while updating event", l.Error(err), l.Any("req", req))
	// 	return nil, status.Error(codes.Internal, "Internal server error")
	// }

	// return &pb.UpdateCourierResponse{
	// 	Courier: courier,
	// }, nil
	return nil, nil
}

func (s *CourierService) GetCourierVehicle(ctx context.Context, req *pb.GetCourierVehicleRequest) (*pb.GetCourierVehicleResponse, error) {
	// var courier *pb.Courier
	// courier, err := s.storage.Courier().GetCourier(req.Id)
	// if err == sql.ErrNoRows {
	// 	s.logger.Error("Error while getting an event, Not found", l.Any("req", req))
	// 	return nil, status.Error(codes.NotFound, "Not found")
	// } else if err != nil {
	// 	s.logger.Error("Error while getting event", l.Error(err), l.Any("req", req))
	// 	return nil, status.Error(codes.Internal, "Internal server error")
	// }

	// return &pb.GetCourierResponse{
	// 	Courier: courier,
	// }, nil
	return nil, nil
}

func (s *CourierService) GetAllCourierVehicles(ctx context.Context, req *pb.GetAllCourierVehiclesRequest) (*pb.GetAllCourierVehiclesResponse, error) {
	// var couriers []*pb.Courier

	// couriers, count, err := s.storage.Courier().GetAllCouriers(req.Page, req.Limit)
	// if err == sql.ErrNoRows {
	// 	s.logger.Error("Error while getting all events, Not found", l.Any("req", req))
	// 	return nil, status.Error(codes.NotFound, "Not found")
	// } else if err != nil {
	// 	s.logger.Error("Error while getting all events", l.Error(err), l.Any("req", req))
	// 	return nil, status.Error(codes.Internal, "Internal server error")
	// }

	// return &pb.GetAllCouriersResponse{
	// 	Couriers: couriers,
	// 	Count:    count,
	// }, nil
	return nil, nil
}

func (s *CourierService) DeleteCourierVehicle(ctx context.Context, req *pb.DeleteCourierVehicleRequest) (*gpb.Empty, error) {
	// err := s.storage.Courier().Delete(req.Id)
	// if err == sql.ErrNoRows {
	// 	s.logger.Error("Error while deleting event, Not found", l.Any("req", req))
	// 	return nil, status.Error(codes.NotFound, "Not found")
	// } else if err != nil {
	// 	s.logger.Error("Error while deleting event", l.Error(err), l.Any("req", req))
	// 	return nil, status.Error(codes.Internal, "Internal server error")
	// }
	return &gpb.Empty{}, nil
}
