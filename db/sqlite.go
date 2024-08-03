package db

import (
	"database/sql"
	"fmt"

	"github.com/1Rob13/net-ekg/models"
)

type Saver interface {
	Save(models.User) error
	RetrieveAllUsers() (models.UserArray, error)
	Delete(models.User) error
}

type SQLiteClient struct {
	DB *sql.DB
}

func NewSQClient() *SQLiteClient {
	//TODO
	//init process sqlite
	// ensure that necessary db and table are present
	return &SQLiteClient{}
}

func (s *SQLiteClient) Save(models.User) error {
	return fmt.Errorf("cant create user cause db not implemented yet")

}
func (s *SQLiteClient) RetrieveAllUsers() (models.UserArray, error) {
	return models.UserArray{{Name: "hans", Email: "dontgetTheFlammenwerfer@problems.de"}, {Name: "joergdeherg", Email: "dontgetTheFlammenwerfer@problems.de"}}, nil
}

func (s *SQLiteClient) Delete(models.User) error {
	panic("not implemented")
}
