package repo

import (
	pb "bitbucket.org/alien_soft/courier_service/genproto/courier_service"
	_ "github.com/lib/pq"
)

type (
	CourierStorageI interface {
		Create(courier *pb.Courier) (*pb.Courier, error)
	}
)
