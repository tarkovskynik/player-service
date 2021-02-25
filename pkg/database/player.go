package database

import (
	"database/sql"
	"github.com/sirupsen/logrus"
	"player"
)

type PlayerRepository struct {
	db *sql.DB
}

func NewInvoiceRepository(db *sql.DB) *PlayerRepository {
	return &PlayerRepository{db: db}
}

func (r *PlayerRepository) Create(user player.User) error {
	var id int

	row := r.db.QueryRow("INSERT INTO users(id, balance, token) VALUES( $1, $2, $3) RETURNING id",
		int(user.Id), user.Balance, user.Token)

	err := row.Scan(&id)
	if err != nil {
		return err
	}
	return nil
}
// I created this method but not use because in instruction I shouldn't have called it from the base
func (r *PlayerRepository) GetById(id int) player.User {
	var user player.User

	row := r.db.QueryRow("SELECT id, balance, token FROM users WHERE id=$1", id)
	err := row.Scan(&user.Id, &user.Balance, &user.Token)
	if err != nil {
		if err == sql.ErrNoRows {
			logrus.WithField("repository", "getbyID").Errorf("error: %s", err.Error())
		}
	}
	return user
}

func (r *PlayerRepository) Update(id int, user player.User) error {
	_, err := r.db.Exec("UPDATE users SET balance=$1 WHERE id=$2",
		user.Balance,id)
	return err
}
