package storage

import (
	"bitbucket.org/alien_soft/courier_service/storage/postgres"
	"bitbucket.org/alien_soft/courier_service/storage/repo"
	"github.com/jmoiron/sqlx"
)

// StorageI ...
type StorageI interface {
	Courier() repo.CourierStorageI
}

type storagePg struct {
	db          *sqlx.DB
	courierRepo repo.CourierStorageI
}

// NewStoragePg ...
func NewStoragePg(db *sqlx.DB) StorageI {
	return &storagePg{
		db:          db,
		courierRepo: postgres.NewCourierRepo(db),
	}
}

func (s storagePg) Courier() repo.CourierStorageI {
	return s.courierRepo
}
