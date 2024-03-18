package database

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

// DogEntity represents entities from the table c_dog.
type DogEntity struct {
	Id     string
	Name   string
	Age    int
	UserId string
}

// FindDogById retrieves one dog from the database. If an error occurs it will be returned, and the output value can be disregarded. If no dog is found,
// both err and output will be nil.
func (dbSvc *DatabaseService) FindDogById(ctx context.Context, id string) (output *DogEntity, err error) {
	logger := logrus.WithContext(ctx).WithFields(logrus.Fields{
		"action": "DatabaseService#FindDogById",
		"id":     id,
	})
	logger.Info("Begin FindDogById")

	const query = "select c_dog_id, name, age, ad_user_id from c_dog where c_dog_id = $1;"

	tx, err := dbSvc.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		logger.WithError(err).Error("Failed to begin transaction.")
		return
	}
	defer (func() {
		if err != nil {
			tx.Rollback()
			logger.Info("Transaction rolled back.")
			return
		}
		tx.Commit()
		logger.Info("Transaction commited.")
	})()

	output = &DogEntity{}

	row := tx.QueryRow(query, id)
	err = row.Scan(
		&(output.Id),
		&(output.Name),
		&(output.Age),
		&(output.UserId),
	)
	if err != nil {
		if err == sql.ErrNoRows {
			output = nil
			err = nil
			logger.Info("Dog not found")
			return
		}
		logger.WithError(err).Error("Failed to retrieve dog.")
		return
	}

	logger.Info("Dog retrieved successfully")
	return
}

// FindDogs lists dogs from the database. If an error occurs it will be returned, and the output value can be disregarded.
func (dbSvc *DatabaseService) FindDogs(ctx context.Context, page, size int) (output []*DogEntity, total int, err error) {
	logger := logrus.WithContext(ctx).WithFields(logrus.Fields{
		"action": "DatabaseService#FindDogs",
		"page":   page,
		"size":   size,
	})
	logger.Info("Begin FindDogs")

	const countQuery = "select count(*) from c_dog"
	const query = "select c_dog_id, name, age, ad_user_id from c_dog offset $1 limit $2"

	tx, err := dbSvc.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		logger.WithError(err).Error("Failed to begin transaction.")
		return
	}
	defer (func() {
		if err != nil {
			tx.Rollback()
			logger.Info("Transaction rolled back.")
			return
		}
		tx.Commit()
		logger.Info("Transaction commited.")
	})()

	row := tx.QueryRow(countQuery)
	err = row.Scan(&total)
	if err != nil {
		logger.WithError(err).Error("Failed to count dogs.")
		return
	}

	rows, err := tx.Query(query, (page-1)*size, size)
	if err != nil {
		logger.WithError(err).Error("Failed to retrieve dogs.")
		return
	}

	output = make([]*DogEntity, 0, size)
	for rows.Next() {
		tmp := &DogEntity{}
		err = rows.Scan(
			&(tmp.Id),
			&(tmp.Name),
			&(tmp.Age),
			&(tmp.UserId),
		)
		output = append(output, tmp)
	}

	logger.Info("Dogs retrieved successfully")
	return
}

// InsertDogInput is used as input on DatabseService#InsertDog.
type InsertDogInput struct {
	Name   string
	Age    int
	UserId string
}

// InsertDog inserts a new dog into the database. If an error occurs it will be returned, and the output value can be disregarded.
func (dbSvc *DatabaseService) InsertDog(ctx context.Context, input InsertDogInput) (output *DogEntity, err error) {
	logger := logrus.WithContext(ctx).WithFields(logrus.Fields{
		"action": "DatabaseService#InsertDog",
		"input":  input,
	})
	logger.Info("Begin InsertDog")

	const query = "insert into c_dog (c_dog_id, name, age, ad_user_id) values ($1, $2, $3, $4)"

	tx, err := dbSvc.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		logger.WithError(err).Error("Failed to begin transaction.")
		return
	}
	defer (func() {
		if err != nil {
			tx.Rollback()
			logger.Info("Transaction rolled back.")
			return
		}
		tx.Commit()
		logger.Info("Transaction commited.")
	})()

	id := uuid.NewString()
	_, err = tx.Exec(query, id, input.Name, input.Age, input.UserId)
	if err != nil {
		logger.WithError(err).Error("Failed to insert dog.")
		return
	}

	output = &DogEntity{
		Id:     id,
		Name:   input.Name,
		Age:    input.Age,
		UserId: input.UserId,
	}

	logger.Info("Inserted dog successfully.")
	return
}
