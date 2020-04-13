package postgres

import (
	"database/sql"
	"time"

	"bitbucket.org/alien_soft/courier_service/pkg/etc"

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

//courier
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
		courierID,
		courier.GetPhone(),
		courier.GetFirstName(),
		courier.GetLastName(),
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
		($1, $2, $3)
		WHERE id = $4`

	_, err = tx.Exec(
		updateQuery,
		courier.GetPhone(),
		courier.GetFirstName(),
		courier.GetLastName(),
		courier.GetId(),
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
	var (
		createdAt  time.Time
		layoutDate string = "2006-01-02 15:04:05"
		courier    pb.Courier
	)

	row := cm.db.QueryRow(`
		SELECT  id,
				phone,
				first_name,
				last_name,
				created_at,
		FROM courier
		WHERE id=$1`, id,
	)

	err := row.Scan(
		&courier.Id,
		&courier.Phone,
		&courier.FirstName,
		&courier.LastName,
		&createdAt,
	)

	courier.CreatedAt = createdAt.Format(layoutDate)
	if err != nil {
		return nil, err
	}

	return &courier, nil
}

func (cm *courierRepo) GetAllCouriers(page, limit uint64) ([]*pb.Courier, uint64, error) {
	var (
		count      uint64
		createdAt  time.Time
		layoutDate string = "2006-01-02 15:04:05"
		couriers   []*pb.Courier
	)

	rows, err := cm.db.Queryx(`
		SELECT  id,
				phone,
				first_name,
				last_name,
				created_at,
		FROM courier
		WHERE status=true`)

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
			&createdAt,
		)

		if err != nil {
			return nil, 0, err
		}
		c.CreatedAt = createdAt.Format(layoutDate)
		couriers = append(couriers, &c)
	}

	row := cm.db.QueryRow(`
		SELECT count(1) 
		FROM courier
		WHERE status=true`,
	)
	err = row.Scan(
		&count,
	)

	return couriers, count, nil
}

func (cm *courierRepo) Delete(id string) error {
	_, err := cm.db.Exec(`
		UPDATE courier SET status=false where id=$1`, id,
	)
	if err != nil {
		return err
	}

	return nil
}

//courier details
func (cm *courierRepo) CreateCourierDetails(cd *pb.CourierDetails) (*pb.CourierDetails, error) {
	var (
		gender  sql.NullString = etc.NullString(cd.Gender)
		address sql.NullString = etc.NullString(cd.Address)
		img     sql.NullString = etc.NullString(cd.Img)
	)

	tx, err := cm.db.Begin()
	if err != nil {
		return nil, err
	}

	ID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	insertNew :=
		`INSERT INTO
		courier_details
		(
			courier_id,
			passport_number,
			gender,
			birth_date,
			address,
			img,
			lisense_number,
			lisense_given_date,
			lisense_expiry_date
		)
		VALUES
		($1, $2, $3, $4, $5, $6, $7, $8, $9)`

	_, err = tx.Exec(
		insertNew,
		cd.GetCourierId(),
		cd.GetPassportNumber(),
		gender,
		cd.GetBirthDate(),
		address,
		img,
		cd.GetLisenseNumber(),
		cd.GetLisenseGivenDate(),
		cd.GetLisenseExpiryDate(),
	)

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	c, err := cm.GetCourierDetails(ID.String())
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (cm *courierRepo) UpdateCourierDetails(cd *pb.CourierDetails) (*pb.CourierDetails, error) {
	var (
		gender  sql.NullString = etc.NullString(cd.Gender)
		address sql.NullString = etc.NullString(cd.Address)
		img     sql.NullString = etc.NullString(cd.Img)
	)

	tx, err := cm.db.Begin()
	if err != nil {
		return nil, err
	}

	updateQuery :=
		`UPDATE courier_details
		(
			passport_number,
			gender,
			birth_date,
			address,
			img,
			lisense_number,
			lisense_given_date,
			lisense_expiry_date
		)
		VALUES
		($1, $2, $3, $4, $5, $6, $7, $8) 
		WHERE courier_id=$9`

	_, err = tx.Exec(
		updateQuery,
		cd.GetPassportNumber(),
		gender,
		cd.GetBirthDate(),
		address,
		img,
		cd.GetLisenseNumber(),
		cd.GetLisenseGivenDate(),
		cd.GetLisenseExpiryDate(),
		cd.GetCourierId(),
	)

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	c, err := cm.GetCourierDetails(cd.GetCourierId())
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (cm *courierRepo) GetCourierDetails(courierId string) (*pb.CourierDetails, error) {
	var (
		gender, address, img sql.NullString
		layout               string = "2006-01-02"
		bDate, gDate, eDate  time.Time
		cd                   pb.CourierDetails
	)

	_, err := uuid.Parse(courierId)

	row := cm.db.QueryRow(`
		SELECT  courier_id
				passport_number,
				gender,
				birth_date,
				address,
				img,
				lisense_number,
				lisense_given_date,
				lisense_expiry_date,
		FROM courier_details
		WHERE courier_id=$1`, courierId)

	err = row.Scan(
		&cd.CourierId,
		&cd.PassportNumber,
		&gender,
		&bDate,
		&address,
		&img,
		&cd.LisenseNumber,
		&gDate,
		&eDate,
	)

	cd.Gender = etc.StringValue(gender)
	cd.BirthDate = bDate.Format(layout)
	cd.Address = etc.StringValue(address)
	cd.Img = etc.StringValue(img)
	cd.LisenseGivenDate = gDate.Format(layout)
	cd.LisenseExpiryDate = eDate.Format(layout)
	if err != nil {
		return nil, err
	}

	return &cd, nil
}

//courier vehicle
func (cm *courierRepo) CreateCourierVehicle(cv *pb.CourierVehicle) (*pb.CourierVehicle, error) {

	tx, err := cm.db.Begin()
	if err != nil {
		return nil, err
	}

	ID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	insertNew :=
		`INSERT INTO
		courier_vehicle
		(
			id,
			courier_id,
			model,
			vehicle_number,
		)
		VALUES
		($1, $2, $3, $4)`

	_, err = tx.Exec(
		insertNew,
		ID,
		cv.GetCourierId(),
		cv.GetModel(),
		cv.GetVehicleNumber(),
	)

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	c, err := cm.GetCourierVehicle(ID.String())
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (cm *courierRepo) UpdateCourierVehicle(cv *pb.CourierVehicle) (*pb.CourierVehicle, error) {

	tx, err := cm.db.Begin()
	if err != nil {
		return nil, err
	}

	updateQuery :=
		`UPDATE courier_vehicle
		(
			courier_id,
			model,
			vehicle_number,
		)
		VALUES
		($1, $2, $3)
		WHERE id = $4`

	_, err = tx.Exec(
		updateQuery,
		cv.GetCourierId(),
		cv.GetModel(),
		cv.GetVehicleNumber(),
		cv.GetId(),
	)

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	c, err := cm.GetCourierVehicle(cv.GetId())
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (cm *courierRepo) GetCourierVehicle(id string) (*pb.CourierVehicle, error) {

	var (
		layout    string = "2006-01-02 15:04:05"
		createdAt time.Time
		cv        pb.CourierVehicle
	)
	_, err := uuid.Parse(id)

	row := cm.db.QueryRow(`
		SELECT  id,
				courier_id,
				model,
				vehicle_number,
				created_at,
		FROM courier_vehicle
		WHERE id=$1`, id,
	)

	err = row.Scan(
		&cv.Id,
		&cv.CourierId,
		&cv.Model,
		&cv.VehicleNumber,
		&createdAt,
	)
	if err != nil {
		return nil, err
	}

	cv.CreatedAt = createdAt.Format(layout)

	return &cv, nil
}

func (cm *courierRepo) GetAllCourierVehicles(courierId string) ([]*pb.CourierVehicle, error) {
	var (
		courierVehicles []*pb.CourierVehicle
		layout          string = "2006-01-02 15:04:05"
		createdAt       time.Time
	)

	rows, err := cm.db.Queryx(`
		SELECT  id,
				courier_id,
				model,
				vehicle_number,
				created_at,
		FROM courier_vehicle`)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var cv pb.CourierVehicle
		err = rows.Scan(
			&cv.Id,
			&cv.CourierId,
			&cv.Model,
			&cv.VehicleNumber,
			&cv.CreatedAt,
		)

		if err != nil {
			return nil, err
		}
		cv.CreatedAt = createdAt.Format(layout)

		courierVehicles = append(courierVehicles, &cv)
	}

	return courierVehicles, nil
}

func (cm *courierRepo) DeleteCourierVehicle(id string) error {
	_, err := cm.db.Exec(`
		UPDATE courier_vehicle SET status = false where id = $1`, id,
	)
	if err != nil {
		return err
	}

	return nil
}
