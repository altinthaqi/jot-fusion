package db

import (
	"github.com/altinthaqi/jot-fusion/model"
)

type Storage interface {
	CreateAccount(*model.Account) error
	DeleteAccount(int) error
	UpdateAccount(*model.Account) error
	GetAccounts() ([]*model.Account, error)
	GetAccountByID(int) (*model.Account, error)
}

func (s *PostgresStore) CreateAccount(acc *model.Account) error {
	query := `insert into account (
		first_name, 
		last_name, 
		number, 
		balance, 
		created_at ) 
		values ($1, $2, $3, $4, $5)`

	_, err := s.db.Query(query,
		acc.FitstName,
		acc.LastName,
		acc.Number,
		acc.Balance,
		acc.CreatedAt)

	if err != nil {
		return err
	}
	return nil
}

func (s *PostgresStore) UpdateAccount(*model.Account) error {
	return nil
}

func (s *PostgresStore) DeleteAccount(id int) error {
	return nil
}

func (s *PostgresStore) GetAccountByID(id int) (*model.Account, error) {
	rows, err := s.db.Query(`select * from account where id = $1`, id)

	if err != nil {
		return nil, err
	}

	account := &model.Account{}

	for rows.Next() {
		err := rows.Scan(&account.ID,
			&account.FitstName,
			&account.LastName,
			&account.Number,
			&account.Balance,
			&account.CreatedAt)

		if err != nil {
			return nil, err
		}
	}

	return account, nil

}

func (s *PostgresStore) GetAccounts() ([]*model.Account, error) {
	rows, err := s.db.Query(`select * from account`)

	if err != nil {
		return nil, nil
	}

	accounts := []*model.Account{}
	for rows.Next() {
		account := new(model.Account)

		err := rows.Scan(&account.ID,
			&account.FitstName,
			&account.LastName,
			&account.Number,
			&account.Balance,
			&account.CreatedAt)

		if err != nil {
			return nil, err
		}

		accounts = append(accounts, account)

	}

	return accounts, nil
}
