package postgres

import (
	pb "bitbucket.org/alien_soft/courier_service/genproto/courier_service"
	"bitbucket.org/alien_soft/courier_service/storage/repo"
	"github.com/google/uuid"
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

	tx, err := cm.db.Begin()
	if err != nil {
		return nil, err
	}

	courierID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	insertNew :=
		`INSERT INTO
		courier
		(
			id,
			phone,
			first_name,
			last_name,
		)
		VALUES
		($1, $2, $3, $4)`

	_, err = tx.Exec(
		insertNew,
		courier.Phone,
		courier.FirstName,
		courier.LastName,
	)

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	c, err := cm.GetCourier(courierID.String())
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (cm *courierRepo) Update(courier *pb.Courier) (*pb.Courier, error) {

	tx, err := cm.db.Begin()
	if err != nil {
		return nil, err
	}

	updateQuery :=
		`UPDATE courier
		(
			phone,
			first_name,
			last_name,
		)
		VALUES
		($1, $2, $3)`

	_, err = tx.Exec(
		updateQuery,
		courier.Phone,
		courier.FirstName,
		courier.LastName,
	)

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	c, err := cm.GetCourier(courier.GetId())
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (cm *courierRepo) GetCourier(id string) (*pb.Courier, error) {
	var courier pb.Courier

	_, err := uuid.Parse(id)

	row := cm.db.QueryRow(`
		SELECT  id,
				phone,
				first_name,
				last_name,
				created_at,
		FROM courier
		WHERE id=$1`, id,
	)

	err = row.Scan(
		&courier.Id,
		&courier.Phone,
		&courier.FirstName,
		&courier.LastName,
		&courier.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &courier, nil
}

func (cm *courierRepo) GetAllCouriers(page, limit uint64) ([]*pb.Courier, uint64, error) {
	var couriers []*pb.Courier

	rows, err := cm.db.Queryx(`
		SELECT  id,
				phone,
				first_name,
				last_name,
				created_at,
		FROM courier`)

	if err != nil {
		return nil, 0, err
	}

	for rows.Next() {
		var c pb.Courier
		err = rows.Scan(
			&c.Id,
			&c.Phone,
			&c.FirstName,
			&c.LastName,
		)

		if err != nil {
			return nil, 0, err
		}

		couriers = append(couriers, &c)
	}

	return couriers, 10, nil
}

func (cm *courierRepo) Delete(id string) error {
	_, err := cm.db.Exec(`
		UPDATE courier SET status = false where id = $1`, id,
	)
	if err != nil {
		return err
	}

	return nil
}
