package postgres

import (
	"database/sql"
	"time"

	pb "genproto/courier_service"

	"bitbucket.org/alien_soft/courier_service/pkg/etc"
	"bitbucket.org/alien_soft/courier_service/storage/repo"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type distributorRepo struct {
	db *sqlx.DB
}

// NewdistributorRepo ...
func NewDistributorRepo(db *sqlx.DB) repo.DistributorStorageI {
	return &distributorRepo{db: db}
}

//Distributor
func (cm *distributorRepo) Create(distributor *pb.Distributor) (*pb.Distributor, error) {

	tx, err := cm.db.Begin()
	if err != nil {
		return nil, err
	}

	insertNew :=
		`INSERT INTO
		distributors
		(
			id,
			access_token,
			name,
			phone
		)
		VALUES
		($1, $2, $3, $4)`

	_, err = tx.Exec(
		insertNew,
		distributor.GetId(),
		distributor.GetAccessToken(),
		distributor.GetName(),
		distributor.GetPhone(),
	)

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	c, err := cm.GetDistributor(distributor.GetId())
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (cm *distributorRepo) Update(distributor *pb.Distributor) (*pb.Distributor, error) {

	tx, err := cm.db.Begin()
	if err != nil {
		return nil, err
	}

	updateQuery :=
		`UPDATE distributors
		SET
			phone=$1,
			name=$2
		WHERE id=$3`

	_, err = tx.Exec(
		updateQuery,
		distributor.GetName(),
		distributor.GetPhone(),
		distributor.GetId(),
	)

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	c, err := cm.GetDistributor(distributor.GetId())
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (cm *distributorRepo) GetDistributor(id string) (*pb.Distributor, error) {
	var (
		createdAt   time.Time
		layoutDate  string = "2006-01-02 15:04:05"
		distributor pb.Distributor
	)

	row := cm.db.QueryRow(`
		SELECT  id,
				access_token,
				name,
				phone,
				created_at,
				is_active
		FROM distributors
		WHERE id=$1`, id,
	)

	err := row.Scan(
		&distributor.Id,
		&distributor.AccessToken,
		&distributor.Name,
		&distributor.Phone,
		&createdAt,
		&distributor.IsActive,
	)

	distributor.CreatedAt = createdAt.Format(layoutDate)
	if err != nil {
		return nil, err
	}

	return &distributor, nil
}

func (cm *distributorRepo) GetAllDistributors(page, limit uint64) ([]*pb.Distributor, uint64, error) {
	var (
		count        uint64
		createdAt    time.Time
		layoutDate   string = "2006-01-02 15:04:05"
		distributors []*pb.Distributor
	)

	offset := (page - 1) * limit
	query := `
		SELECT  id,
				access_token,
				name,
				phone,
				created_at,
				is_active
		FROM distributors
		WHERE deleted_at IS NULL
		ORDER BY created_at DESC 
		LIMIT $1 OFFSET $2`
	rows, err := cm.db.Queryx(query, limit, offset)

	if err != nil {
		return nil, 0, err
	}

	for rows.Next() {
		var d pb.Distributor
		err = rows.Scan(
			&d.Id,
			&d.AccessToken,
			&d.Name,
			&d.Phone,
			&createdAt,
			&d.IsActive,
		)

		if err != nil {
			return nil, 0, err
		}
		d.CreatedAt = createdAt.Format(layoutDate)
		distributors = append(distributors, &d)
	}

	row := cm.db.QueryRow(`
		SELECT count(1) 
		FROM distributors
		WHERE deleted_at IS NULL`,
	)
	err = row.Scan(
		&count,
	)

	return distributors, count, nil
}

func (cm *distributorRepo) Delete(id string) error {
	_, err := cm.db.Exec(`
		UPDATE distributors SET deleted_at=CURRENT_TIMESTAMP where id=$1`, id,
	)
	if err != nil {
		return err
	}

	return nil
}

//Park
func (cm *distributorRepo) CreatePark(park *pb.Park) (*pb.Park, error) {
	var address sql.NullString = etc.NullString(park.Address)

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
		parks
		(
			id,
			distributor_id,
			name,
			location,
			address
		)
		VALUES
		($1, $2, $3, st_makepoint($4, $5), $6)`

	_, err = tx.Exec(
		insertNew,
		ID,
		park.GetDistributorId(),
		park.GetName(),
		park.GetLocation().GetLong(),
		park.GetLocation().GetLat(),
		address,
	)

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	p, err := cm.GetPark(ID.String())
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (cm *distributorRepo) UpdatePark(park *pb.Park) (*pb.Park, error) {
	var address sql.NullString = etc.NullString(park.Address)

	tx, err := cm.db.Begin()
	if err != nil {
		return nil, err
	}

	updateQuery :=
		`UPDATE parks
		SET
			name=$1,
			location=st_makepoint($2, $3),
			address=$4
		WHERE id=$5`

	_, err = tx.Exec(
		updateQuery,
		park.GetName(),
		park.GetLocation().GetLong(),
		park.GetLocation().GetLat(),
		address,
		park.GetId(),
	)

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	p, err := cm.GetPark(park.GetId())
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (cm *distributorRepo) GetPark(id string) (*pb.Park, error) {
	var (
		createdAt  time.Time
		layoutDate string = "2006-01-02 15:04:05"
		park       pb.Park
		location   pb.Location
		address    sql.NullString
	)

	row := cm.db.QueryRow(`
		SELECT  id,
				distributor_id,
				name,
				st_x(location),
				st_y(location),	
				address,
				created_at,
				is_active
		FROM parks
		WHERE id=$1`, id,
	)

	err := row.Scan(
		&park.Id,
		&park.DistributorId,
		&park.Name,
		&location.Long,
		&location.Lat,
		&address,
		&createdAt,
		&park.IsActive,
	)

	if err != nil {
		return nil, err
	}
	park.Address = etc.StringValue(address)
	park.Location = &location
	park.CreatedAt = createdAt.Format(layoutDate)

	return &park, nil
}

func (cm *distributorRepo) GetAllDistributorParks(page, limit uint64) ([]*pb.Park, uint64, error) {
	var (
		count      uint64
		createdAt  time.Time
		layoutDate string = "2006-01-02 15:04:05"
		parks      []*pb.Park
		location   pb.Location
		address    sql.NullString
	)

	offset := (page - 1) * limit
	query := `
		SELECT  id,
				distributor_id,
				name,
				st_x(location),
				st_y(location),	
				address,
				created_at,
				is_active
		FROM parks
		WHERE deleted_at IS NULL
		ORDER BY created_at DESC 
		LIMIT $1 OFFSET $2`
	rows, err := cm.db.Queryx(query, limit, offset)

	if err != nil {
		return nil, 0, err
	}

	for rows.Next() {
		var park pb.Park
		err = rows.Scan(
			&park.Id,
			&park.DistributorId,
			&park.Name,
			&location.Long,
			&location.Lat,
			&address,
			&createdAt,
			&park.IsActive,
		)

		if err != nil {
			return nil, 0, err
		}
		park.Address = etc.StringValue(address)
		park.Location = &location
		park.CreatedAt = createdAt.Format(layoutDate)
		parks = append(parks, &park)
	}

	row := cm.db.QueryRow(`
		SELECT count(1) 
		FROM parks
		WHERE deleted_at IS NULL`,
	)
	err = row.Scan(
		&count,
	)

	return parks, count, nil
}

func (cm *distributorRepo) GetAllParks(page, limit uint64) ([]*pb.Park, uint64, error) {
	var (
		count      uint64
		createdAt  time.Time
		layoutDate string = "2006-01-02 15:04:05"
		parks      []*pb.Park
	)

	offset := (page - 1) * limit
	query := `
		SELECT  id,
				access_token,
				name,
				phone,
				created_at,
				is_active
		FROM parks
		WHERE deleted_at=NULL
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2`
	rows, err := cm.db.Queryx(query, limit, offset)

	if err != nil {
		return nil, 0, err
	}

	for rows.Next() {
		var p pb.Park
		err = rows.Scan(
			&p.Id,

			&p.Name,

			&createdAt,
			&p.IsActive,
		)

		if err != nil {
			return nil, 0, err
		}
		p.CreatedAt = createdAt.Format(layoutDate)
		parks = append(parks, &p)
	}

	row := cm.db.QueryRow(`
		SELECT count(1)
		FROM parks
		WHERE deleted_at=NULL`,
	)
	err = row.Scan(
		&count,
	)

	return parks, count, nil
}

func (cm *distributorRepo) DeletePark(id string) error {
	_, err := cm.db.Exec(`
		UPDATE parks SET deleted_at=CURRENT_TIMESTAMP where id=$1`, id,
	)
	if err != nil {
		return err
	}

	return nil
}
