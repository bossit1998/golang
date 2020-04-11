package postgres

import (
	pb "bitbucket.org/alien_soft/courier_service/genproto/courier_service"
	"bitbucket.org/alien_soft/courier_service/storage/repo"
	"github.com/jmoiron/sqlx"
)

type courierRepo struct {
	db *sqlx.DB
}

// NewCourierRepo ...
func NewCourierRepo(db *sqlx.DB) repo.CourierStorageI {
	return &courierRepo{db: db}
}

func (cm *courierRepo) Create(courier *pb.Courier) (*pb.Courier, error) {
	// var (
	// 	phone     sql.NullString = etc.NullString(courier.Phone)
	// 	firstName sql.NullString = etc.NullString(courier.FirstName)
	// 	lastName  sql.NullString = etc.NullString(courier.LastName)
	// )

	// tx, err := cm.db.Begin()
	// if err != nil {
	// 	return nil, err
	// }

	// courierID, err := uuid.NewRandom()
	// if err != nil {
	// 	return nil, err
	// }

	// insertNew :=
	// 	`INSERT INTO
	// 	courier
	// 	(
	// 		id,
	// 		phone,
	// 		first_name,
	// 		last_name,
	// 	)
	// 	VALUES
	// 	($1, $2, $3, $4)`

	// _, err = tx.Exec(
	// 	phone,
	// 	firstName,
	// 	lastName,
	// )

	// if err != nil {
	// 	tx.Rollback()
	// 	return nil, err
	// }

	// tx.Commit()

	// courier, err := cm.GetCourier(courierID.String())
	// if err != nil {
	// 	return nil, err
	// }

	return nil, nil
}
