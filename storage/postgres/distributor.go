package postgres

import (
	pb "bitbucket.org/alien_soft/courier_service/genproto/courier_service"
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

func (cm *distributorRepo) Create(distributor *pb.Distributor) (*pb.Distributor, error) {

	tx, err := cm.db.Begin()
	if err != nil {
		return nil, err
	}

	distributorID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	insertNew :=
		`INSERT INTO
		distributor
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
		distributor.Phone,
	)

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	c, err := cm.GetDistributor(distributorID.String())
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
		`UPDATE distributor
		(
			phone,
			first_name,
			last_name,
		)
		VALUES
		($1, $2, $3)`

	_, err = tx.Exec(
		updateQuery,
		distributor.Phone,
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
	var distributor pb.Distributor

	_, err := uuid.Parse(id)

	row := cm.db.QueryRow(`
		SELECT  id,
				phone,
				first_name,
				last_name,
				created_at,
		FROM distributor
		WHERE id=$1`, id,
	)

	err = row.Scan(
		&distributor.Id,
		&distributor.Phone,
		&distributor.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &distributor, nil
}

func (cm *distributorRepo) GetAllDistributors(page, limit uint64) ([]*pb.Distributor, uint64, error) {
	var distributors []*pb.Distributor

	rows, err := cm.db.Queryx(`
		SELECT  id,
				phone,
				first_name,
				last_name,
				created_at,
		FROM distributor`)

	if err != nil {
		return nil, 0, err
	}

	for rows.Next() {
		var c pb.Distributor
		err = rows.Scan(
			&c.Id,
			&c.Phone,
		)

		if err != nil {
			return nil, 0, err
		}

		distributors = append(distributors, &c)
	}

	return distributors, 10, nil
}

func (cm *distributorRepo) Delete(id string) error {
	_, err := cm.db.Exec(`
		UPDATE distributor SET status = false where id = $1`, id,
	)
	if err != nil {
		return err
	}

	return nil
}
