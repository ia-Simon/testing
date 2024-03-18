package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

// DatabaseService holds the database connection and exposes methods to interact with it.
type DatabaseService struct {
	// Database connection
	db *sql.DB
}

// Config is used to configure the database connection inside the New DatabaseService constructor.
type Config struct {
	// Database host
	Host string
	// Database port
	Port string
	// Database username
	Username string
	// Database password
	Password string
	// Database name
	Name string
}

// New creates a new DatabaseService based on the given configuration.
func New(config Config) *DatabaseService {
	logger := logrus.WithFields(logrus.Fields{
		"action": "(package database)#New",
		"config": config,
	})
	logger.Info("Create new DatabaseService")

	db, err := sql.Open("postgres", fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		config.Username,
		config.Password,
		config.Name,
		config.Host,
		config.Port,
	))
	if err != nil {
		logrus.WithError(err).Fatal("Error while creating database connection.")
	}
	err = db.Ping()
	if err != nil {
		logrus.WithError(err).Fatal("Error while connecting to database.")
	}

	return &DatabaseService{db}
}
