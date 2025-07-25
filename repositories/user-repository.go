package repositories

import (
	"errors"

	"example.com/rest-api/db"
	"example.com/rest-api/models"
	"example.com/rest-api/utils"
)

type UserRepository struct{}

func (r *UserRepository) CreateUser(u *models.User) error {
	query := `
	INSERT INTO users (email, password)
	VALUES (?, ?)`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	hashPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Email, hashPassword)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	u.ID = id
	return nil
}

func (r *UserRepository) Authenticate(u *models.User) error {
	query := `
	SELECT id, password FROM users WHERE email = ?`

	row := db.DB.QueryRow(query, u.Email)
	var storedID int64
	var storedPassword string
	err := row.Scan(&storedID, &storedPassword)

	if err != nil {
		return errors.New("user not found")
	}

	passwordIsValid := utils.CheckPasswordHash(u.Password, storedPassword)
	if !passwordIsValid {
		return errors.New("invalid password")
	}

	u.ID = storedID

	return nil
}