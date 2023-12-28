package db

import (
	"fmt"
	"log"

	"github.com/altinthaqi/jot-fusion/model"
)

func (s *PostgresStore) Init() error {

	// s.seedAccounts()
	return s.createAccountTable()

}

func (s *PostgresStore) seedAccount(fname, lname, pw string) *model.Account {
	acc, err := model.NewAccount(fname, lname, pw)

	if err != nil {
		log.Fatal(err)
	}

	if err := s.CreateAccount(acc); err != nil {
		log.Fatal(err)
	}
	fmt.Println("new account", acc)

	return acc
}

func (s *PostgresStore) seedAccounts() error {
	s.seedAccount("testMe", "testMe", "123") // 7700

	return nil
}

func (s *PostgresStore) createAccountTable() error {
	query := `create table if not exists account (
		id serial primary key,
		first_name varchar(50),
		last_name varchar(50),
		encrypted_password,
		number serial,
		balance serial,
		created_at timestamp 
	)`

	_, err := s.db.Exec(query)

	return err
}
