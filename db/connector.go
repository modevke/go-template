package db

import (

	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Repositories struct {
	db *gorm.DB
}

func NewRepositories(DbUser, DbPassword, DbPort, DbHost, DbName string) (*Repositories, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", DbHost, DbUser, DbPassword, DbName, DbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return &Repositories{
		db: db,
	}, nil
}

func (s *Repositories) Close() error {
	sqlDb, err := s.db.DB()
	if err != nil {
		return sqlDb.Close()
	}

	return err
}