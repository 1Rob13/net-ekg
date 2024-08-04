package db

import (
	"database/sql"
	"fmt"
	"log/slog"
	"os"

	_ "github.com/mattn/go-sqlite3"

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

	slog.Debug("init sql db")

	//TODO do not delete old one by creating check for file first
	file, err := os.Create("sqLite-db-subscribers.db")

	if err != nil {
		slog.Error("cant create db file")
	}

	file.Close()
	slog.Debug("success on db write")

	sqLiteDB, err := sql.Open("sqlite3", "./sqLite-db-subscribers.db")

	if err != nil {
		slog.Error("cant create db connection")
		fmt.Printf("db connection to file failed because: %v", err)
	}

	//TODO figure out where this func can run and not die right after create
	// defer sqLiteDB.Close()

	//adding the base table
	createTable(sqLiteDB)
	return &SQLiteClient{DB: sqLiteDB}
}

func (s *SQLiteClient) Save(u models.User) error {

	//s is nil here...

	return s.insertSubscriber(u.Name, u.Email.String())

}
func (s *SQLiteClient) RetrieveAllUsers() (models.UserArray, error) {
	return models.UserArray{{Name: "hans", Email: "dontgetTheFlammenwerfer@problems.de"}, {Name: "joergdeherg", Email: "dontgetTheFlammenwerfer@problems.de"}}, nil
}

func (s *SQLiteClient) Delete(models.User) error {
	panic("not implemented")
}

func createTable(s *sql.DB) {

	createSubscribersTableSQL := `CREATE TABLE subscribers (
		"idSubscribers" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"name" TEXT,
		"email" TEXT		
	  );`

	slog.Debug("start create table string")
	st, err := s.Prepare(createSubscribersTableSQL)

	if err != nil {
		slog.Error("cant prepare sql statement")
		fmt.Printf("cant prepare sql statement because: %v", err)
	}

	res, err := st.Exec()

	if err != nil {
		slog.Error("cant exec sql statement")
		fmt.Printf("cant exec sql statement because: %v", err)
		fmt.Println(res)

	}

	// slog.Debug(res.RowsAffected)
	slog.Debug("success creating subscribers table")
}

func (s *SQLiteClient) insertSubscriber(name string, email string) error {
	slog.Debug("inserting subscriber")

	insertStudentSQL := `INSERT INTO subscribers( name, email) VALUES (?, ?)`
	st, err := s.DB.Prepare(insertStudentSQL) // Prepare statement.
	// This is good to avoid SQL injections
	if err != nil {
		slog.Error("cant prepare sql statement")
		return fmt.Errorf("cant prepare sql statement because: %v", err)
	}
	res, err := st.Exec(name, email)

	if err != nil {
		slog.Error("cant exec sql statement")
		fmt.Println(res)
		return fmt.Errorf("cant exec sql statement because: %v", err)

	}

	return nil
}
