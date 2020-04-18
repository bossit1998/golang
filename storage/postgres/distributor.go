package postgres

import (
	"time"

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

	ID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	insertNew :=
		`INSERT INTO
		distributors
		(
			id,
			name,
			phone
		)
		VALUES
		($1, $2, $3)`

	_, err = tx.Exec(
		insertNew,
		ID,
		distributor.GetName(),
		distributor.GetPhone(),
	)

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	c, err := cm.GetDistributor(ID.String())
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
				name,
				phone,
				created_at
		FROM distributors
		WHERE id=$1`, id,
	)

	err := row.Scan(
		&distributor.Id,
		&distributor.Name,
		&distributor.Phone,
		&createdAt,
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
				name,
				phone,
				created_at
		FROM distributors
		WHERE status=true
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
			&d.Name,
			&d.Phone,
			&createdAt,
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
		WHERE status=true`,
	)
	err = row.Scan(
		&count,
	)

	return distributors, count, nil
}

func (cm *distributorRepo) Delete(id string) error {
	_, err := cm.db.Exec(`
		UPDATE distributors SET status = false where id = $1`, id,
	)
	if err != nil {
		return err
	}

	return nil
}
