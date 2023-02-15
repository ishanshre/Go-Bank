package storage

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"

	"github.com/ishanshre/Go-Bank/pkg/models"
	"github.com/joho/godotenv"
)

type Storage interface {
	CreateAccount(*models.Account) error
	DeleteAccount(int) error
	UpdateAccount(int, *models.Account) error
	GetAccounts() ([]*models.Account, error)
	GetAccountById(int) (*models.Account, error)
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatalf("Error in loading the environment files: %v", err)
	}
	db, err := sql.Open("postgres", os.Getenv("POSTGRES_CONN_STRING"))
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &PostgresStore{
		db: db,
	}, nil

}

func (s *PostgresStore) Init() error {
	return s.createAccountTable()
}

func (s *PostgresStore) createAccountTable() error {
	query := `CREATE TABLE IF NOT EXISTS account (
		id SERIAL PRIMARY KEY,
		first_name VARCHAR(50),
		last_name VARCHAR(50),
		account_number BIGINT,
		balance BIGINT,
		created_at TIMESTAMPTZ
	);`
	_, err := s.db.Exec(query)
	return err
}
func (s *PostgresStore) CreateAccount(account *models.Account) error {
	query := `INSERT INTO account 
			(first_name, last_name, account_number, balance, created_at) 
			VALUES ($1, $2, $3, $4, $5)
			`
	_, err := s.db.Query(
		query,
		account.FirstName,
		account.LastName,
		account.AccountNumber,
		account.Balance,
		account.CreatedAt,
	)
	if err != nil {
		return err
	}
	log.Printf("Account Created Successfull")
	return nil
}

func (s *PostgresStore) DeleteAccount(id int) error {
	query := `DELETE FROM account where id = $1`
	s.db.Exec("COMMIT")
	_, err := s.db.Query(query, id)
	return err
}

func (s *PostgresStore) UpdateAccount(id int, account *models.Account) error {
	query := `
		UPDATE account 
		SET first_name = $1, last_name = $2
		WHERE id = $3
	`
	s.db.Exec("COMMIT")
	_, err := s.db.Query(query, account.FirstName, account.LastName, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *PostgresStore) GetAccounts() ([]*models.Account, error) {
	rows, err := s.db.Query("SELECT * FROM account")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	accounts := []*models.Account{}
	for rows.Next() {
		account, err := scanAccounts(rows)
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, account)
	}
	return accounts, nil
}

func (s *PostgresStore) GetAccountById(id int) (*models.Account, error) {
	query := `SELECT * FROM account where id = $1`
	rows, err := s.db.Query(query, id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanAccounts(rows)
	}
	return nil, fmt.Errorf("account with id %d not found", id)
}

func scanAccounts(rows *sql.Rows) (*models.Account, error) {
	account := new(models.Account)
	err := rows.Scan(
		&account.ID,
		&account.FirstName,
		&account.LastName,
		&account.AccountNumber,
		&account.Balance,
		&account.CreatedAt,
	)
	return account, err
}
