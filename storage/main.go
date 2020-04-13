package storage

import (
	"bitbucket.org/alien_soft/courier_service/storage/postgres"
	"bitbucket.org/alien_soft/courier_service/storage/repo"
	"github.com/jmoiron/sqlx"
)

// StorageI ...
type StorageI interface {
	Courier() repo.CourierStorageI
	Distributor() repo.DistributorStorageI
}

type storagePg struct {
	db              *sqlx.DB
	courierRepo     repo.CourierStorageI
	distributorRepo repo.DistributorStorageI
}

// NewStoragePg ...
func NewStoragePg(db *sqlx.DB) StorageI {
	return &storagePg{
		db:              db,
		courierRepo:     postgres.NewCourierRepo(db),
		distributorRepo: postgres.NewDistributorRepo(db),
	}
}

func (s storagePg) Courier() repo.CourierStorageI {
	return s.courierRepo
}

func (s storagePg) Distributor() repo.DistributorStorageI {
	return s.distributorRepo
}
