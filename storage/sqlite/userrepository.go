package sqlite

import (
	"context"

	"github.com/Negat1v9/bot-core/models"
)

// TODO: Add to all methods context.
type UserRepository struct {
	storage *Storage
}

// Method for start programm progress once.
func (r *UserRepository) CreateTable() error {
	stmt, err := r.storage.db.Prepare(`
	CREATE TABLE IF NOT EXISTS users(
		id INTEGER PRIMARY KEY, 
		user_name TEXT, 
		chat_id INTEGER NOT NULL);
	`)
	if err != nil {
		return err
	}
	_, err = stmt.Exec()
	if err != nil {
		return err
	}
	return nil
}

// Save user in database.
func (r *UserRepository) SaveUser(_ context.Context, u *models.User) error {
	stmt, err := r.storage.db.Prepare(`
	INSERT INTO users (id, user_name, chat_id) VALUES (?, ?, ?);
	`)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(u.Id, u.Name, u.ChatID)
	if err != nil {
		return err
	}
	return nil
}

// Find user by id.
func (r *UserRepository) Find(_ context.Context, ID int) (*models.User, error) {
	stmt, err := r.storage.db.Prepare("SELECT * FROM users WHERE id = ?;")
	if err != nil {
		return nil, err
	}
	u := &models.User{}
	err = stmt.QueryRow(ID).Scan(&u.Id, &u.Name, &u.ChatID)
	if err != nil {
		return nil, err
	}
	return u, nil
}

// TODO: implemente this method...
// Change user parameters.
func (r *UserRepository) ChangeUser(_ context.Context, ID int) error {
	panic("not implemented")
}
