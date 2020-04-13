package repo

import (
	pb "bitbucket.org/alien_soft/courier_service/genproto/courier_service"
	_ "github.com/lib/pq"
)

type (
	DistributorStorageI interface {
		Create(Distributor *pb.Distributor) (*pb.Distributor, error)
		Update(Distributor *pb.Distributor) (*pb.Distributor, error)
		GetDistributor(id string) (*pb.Distributor, error)
		GetAllDistributors(page, limit uint64) ([]*pb.Distributor, uint64, error)
		Delete(id string) error
	}
)
