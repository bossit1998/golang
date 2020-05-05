package repo

import (
	pb "genproto/courier_service"

	_ "github.com/lib/pq"
)

type (
	DistributorStorageI interface {
		Create(Distributor *pb.Distributor) (*pb.Distributor, error)
		Update(Distributor *pb.Distributor) (*pb.Distributor, error)
		GetDistributor(id string) (*pb.Distributor, error)
		GetAllDistributors(page, limit uint64) ([]*pb.Distributor, uint64, error)
		Delete(id string) error

		CreatePark(Park *pb.Park) (*pb.Park, error)
		UpdatePark(Park *pb.Park) (*pb.Park, error)
		GetPark(id string) (*pb.Park, error)
		GetAllDistributorParks(page, limit uint64) ([]*pb.Park, uint64, error)
		GetAllParks(page, limit uint64) ([]*pb.Park, uint64, error)
		DeletePark(id string) error
	}
)
