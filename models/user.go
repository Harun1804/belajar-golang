package models

import (
	"errors"

	"example.com/rest-api/db"
	"example.com/rest-api/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) CreateUser() error {
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

func (u *User) Authenticate() error {
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

	// if !utils.CheckPasswordHash(u.Password, storedPassword) {
	// 	return "", utils.ErrInvalidCredentials
	// }

	// token, err := utils.GenerateJWT(u.ID)
	// if err != nil {
	// 	return "", err
	// }
}
