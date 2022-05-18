package domain

import (
	"database/sql"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
	"github.com/stakkato95/service-engineering-go-lib/logger"
	"github.com/stakkato95/twitter-service-users/config"
)

type UserRepo interface {
	Create(*User) (*User, error)
	Authenticate(*User) (string, error)
}

type defaultUserRepo struct {
	db *sql.DB
}

func NewUserRepo() UserRepo {
	db, err := sql.Open(config.AppConfig.DbDriver, config.GetConnectionString())

	if err != nil {
		logger.Panic("can not open database connection: " + err.Error())
	}
	if err = db.Ping(); err != nil {
		logger.Panic("can not ping database: " + err.Error())
	}

	repo := &defaultUserRepo{db: db}
	// repo.migrate()

	return repo
}

func (r *defaultUserRepo) Create(user *User) (*User, error) {
	return nil, nil
}

func (r *defaultUserRepo) Authenticate(user *User) (string, error) {
	return "password", nil
}

func (r *defaultUserRepo) migrate() {
	driver, _ := mysql.WithInstance(r.db, &mysql.Config{})
	m, _ := migrate.NewWithDatabaseInstance(
		"file:///migrations",
		config.AppConfig.DbName,
		driver,
	)
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		logger.Fatal("can not migrate up: " + err.Error())
	}
}
