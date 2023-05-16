package storage

import "database/sql"

type Storage interface{}

type dataBase struct {
	db *sql.DB
}

func NewStorage(driver, source string) (Storage, error) {
	db, err := sql.Open(driver, source)
	if err != nil {
		return nil, err
	}
	return dataBase{
		db: db,
	}, nil
}
