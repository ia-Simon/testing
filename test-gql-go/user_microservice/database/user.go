package database

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

// UserEntity represents entities from the table ad_user.
type UserEntity struct {
	Id   string
	Name string
	Age  int
}

// FindUserById retrieves one user from the database. If an error occurs it will be returned, and the output value can be disregarded. If no user is found,
// both err and output will be nil.
func (dbSvc *DatabaseService) FindUserById(ctx context.Context, id string) (output *UserEntity, err error) {
	logger := logrus.WithContext(ctx).WithFields(logrus.Fields{
		"action": "DatabaseService#FindUserById",
		"id":     id,
	})
	logger.Info("Begin FindUserById")

	const query = "select ad_user_id, name, age from ad_user where ad_user_id = $1;"

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

	output = &UserEntity{}

	row := tx.QueryRow(query, id)
	err = row.Scan(
		&(output.Id),
		&(output.Name),
		&(output.Age),
	)
	if err != nil {
		if err == sql.ErrNoRows {
			output = nil
			err = nil
			logger.Info("User not found")
			return
		}
		logger.WithError(err).Error("Failed to retrieve user.")
		return
	}

	logger.Info("User retrieved successfully")
	return
}

// FindUsers lists users from the database. If an error occurs it will be returned, and the output value can be disregarded.
func (dbSvc *DatabaseService) FindUsers(ctx context.Context, page, size int) (output []*UserEntity, total int, err error) {
	logger := logrus.WithContext(ctx).WithFields(logrus.Fields{
		"action": "DatabaseService#FindUsers",
		"page":   page,
		"size":   size,
	})
	logger.Info("Begin FindUsers")

	const countQuery = "select count(*) from ad_user"
	const query = "select ad_user_id, name, age from ad_user offset $1 limit $2"

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
		logger.WithError(err).Error("Failed to count users.")
		return
	}

	rows, err := tx.Query(query, (page-1)*size, size)
	if err != nil {
		logger.WithError(err).Error("Failed to retrieve users.")
		return
	}

	output = make([]*UserEntity, 0, size)
	for rows.Next() {
		tmp := &UserEntity{}
		err = rows.Scan(
			&(tmp.Id),
			&(tmp.Name),
			&(tmp.Age),
		)
		output = append(output, tmp)
	}

	logger.Info("Users retrieved successfully")
	return
}

// InsertUserInput is used as input on DatabseService#InsertUser.
type InsertUserInput struct {
	Name string
	Age  int
}

// InsertUser inserts a new user into the database. If an error occurs it will be returned, and the output value can be disregarded.
func (dbSvc *DatabaseService) InsertUser(ctx context.Context, input InsertUserInput) (output *UserEntity, err error) {
	logger := logrus.WithContext(ctx).WithFields(logrus.Fields{
		"action": "DatabaseService#InsertUser",
		"input":  input,
	})
	logger.Info("Begin InsertUser")

	const query = "insert into ad_user (ad_user_id, name, age) values ($1, $2, $3)"

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
	_, err = tx.Exec(query, id, input.Name, input.Age)
	if err != nil {
		logger.WithError(err).Error("Failed to insert user.")
		return
	}

	output = &UserEntity{
		Id:   id,
		Name: input.Name,
		Age:  input.Age,
	}

	logger.Info("Inserted user successfully.")
	return
}
