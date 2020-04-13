package repo

import (
	pb "bitbucket.org/alien_soft/courier_service/genproto/courier_service"
	_ "github.com/lib/pq"
)

type (
	CourierStorageI interface {
		Create(courier *pb.Courier) (*pb.Courier, error)
		Update(courier *pb.Courier) (*pb.Courier, error)
		GetCourier(id string) (*pb.Courier, error)
		GetAllCouriers(page, limit uint64) ([]*pb.Courier, uint64, error)
		Delete(id string) error

		CreateCourierDetails(cd *pb.CourierDetails) (*pb.CourierDetails, error)
		UpdateCourierDetails(cd *pb.CourierDetails) (*pb.CourierDetails, error)
		GetCourierDetails(courierId string) (*pb.CourierDetails, error)

		CreateCourierVehicle(cv *pb.CourierVehicle) (*pb.CourierVehicle, error)
		UpdateCourierVehicle(cv *pb.CourierVehicle) (*pb.CourierVehicle, error)
		GetCourierVehicle(id string) (*pb.CourierVehicle, error)
		GetAllCourierVehicles(courierId string) ([]*pb.CourierVehicle, error)
		DeleteCourierVehicle(id string) error
	}
)
