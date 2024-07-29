package db

import (
	"database/sql"

	"github.com/1Rob13/net-ekg/models"
)

type Saver interface {
	Save(models.User) error
	RetrieveAllUsers() ([]models.User, error)
	Delete(models.User) error
}

type SQLiteClient struct {
	DB *sql.DB
}

func New() *SQLiteClient {
	//TODO
	//init process sqlite
	// ensure that necessary db and table are present
	return &SQLiteClient{}
}

func (s *SQLiteClient) Save(models.User) error {
	panic("not implemented")
}
func (s *SQLiteClient) RetrieveAllUsers() ([]models.User, error) {
	panic("not implemented")
}

func (s *SQLiteClient) Delete(models.User) error {
	panic("not implemented")
}
