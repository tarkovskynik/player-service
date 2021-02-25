package database

import (
"database/sql"
"player"
)

type PlayerRepository struct {
	db *sql.DB
}

func NewInvoiceRepository(db *sql.DB) *PlayerRepository {
	return &PlayerRepository{db: db}
}

func (r *PlayerRepository) Create(user player.User) (int, error) {
	var id int
	row := r.db.QueryRow("INSERT INTO users(balance, token) VALUES( $1, $2 ) RETURNING id",
		user.Balance, user.Token)
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *PlayerRepository) GetById(id int) (player.User, error) {
	var user player.User

	row := r.db.QueryRow("SELECT id, balance, token FROM users WHERE id=$1", id)
	err := row.Scan(&user.Id, &user.Balance, &user.Token)
	if err != nil {
		if err == sql.ErrNoRows {
			return player.User{}, player.NewErrorUserNotFound(id)
		}

		return player.User{}, err
	}

	return user, nil
}

func (r *PlayerRepository) Update(id int, user player.User) error {
	_, err := r.db.Exec("UPDATE invoices SET balance=$1, token=$2 WHERE id=$3",
		user.Balance, user.Token, id)
	return err
}