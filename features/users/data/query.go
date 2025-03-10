package data

import (
	"database/sql"
	"errors"
	"presensi/features/users"
	"presensi/utils"
)

type UserData struct {
	DB *sql.DB
}

func NewUserDataRepository(db *sql.DB) users.UserDataInterface {
	return &UserData{
		DB: db,
	}
}

func (d *UserData) Register(user users.User) error {
	isEmail := d.IsEmailExist(user.Email)
	if isEmail {
		return errors.New("email ada")
	}
	query := `INSERT INTO users (id, nim, email, password) VALUES (?, ?, ?, ?)`
	_, err := d.DB.Exec(query, user.ID, user.NIM, user.Email, user.Password)
	if err != nil {
		return err
	}
	return nil
}

func (d *UserData) Login(user users.User) (*users.Login, error) {
	var dataTable User
	query := `SELECT id, nim, email, password FROM users WHERE email = ?`
	err := d.DB.QueryRow(query, user.Email).Scan(&dataTable.ID, &dataTable.NIM, &dataTable.Email, &dataTable.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("email tidak ada")
		}
		return nil, err
	}
	if !utils.CheckPasswordHash(user.Password, dataTable.Password) {
		return nil, errors.New("password salah")
	}

	userLogin := &users.Login{
		ID:    dataTable.ID,
		Email: dataTable.Email,
		Token: "",
	}

	return userLogin, nil
}

func (d *UserData) IsEmailExist(email string) bool {
	var userEmail string
	query := `SELECT email FROM users WHERE email = $1`
	err := d.DB.QueryRow(query, email).Scan(&userEmail)
	if err != nil {
		if err == sql.ErrNoRows {
			return false
		}
		return false
	}
	return true
}
