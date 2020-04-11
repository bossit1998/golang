package service

import (
	"context"

	"github.com/jmoiron/sqlx"

	pb "bitbucket.org/alien_soft/courier_service/genproto/courier_service"

	l "bitbucket.org/alien_soft/courier_service/pkg/logger"
	"bitbucket.org/alien_soft/courier_service/storage"
)

// EventService ...
type CourierService struct {
	storage storage.StorageI
	logger  l.Logger
}

// NewEventService ...
func NewCourierService(db *sqlx.DB, log l.Logger) *CourierService {
	return &CourierService{
		storage: storage.NewStoragePg(db),
		logger:  log,
	}
}

// Create ...
func (s *CourierService) Create(ctx context.Context, req *pb.Courier) (*pb.CreateCourierResponse, error) {
	// var err error

	// courier, err := s.storage.Courier().Create(req)
	// if err != nil {
	// 	s.logger.Error("Error while creating event", l.Error(err), l.Any("req", req))
	// 	return nil, status.Error(codes.Internal, err.Error())
	// }
	// return &pb.CreateCourierResponse{
	// 	Courier: courier,
	// }, nil

	_, _ = s.storage.Courier().Create(req)

	return nil, nil
}

func (s *CourierService) Update(ctx context.Context, req *pb.Courier) (*pb.UpdateCourierResponse, error) {
	// var err error

	// courier, err := s.storage.Courier().Create(req)
	// if err != nil {
	// 	s.logger.Error("Error while creating event", l.Error(err), l.Any("req", req))
	// 	return nil, status.Error(codes.Internal, err.Error())
	// }
	// return &pb.CreateCourierResponse{
	// 	Courier: courier,
	// }, nil

	_, _ = s.storage.Courier().Create(req)

	return nil, nil
}
