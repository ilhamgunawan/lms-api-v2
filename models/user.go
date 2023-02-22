package models

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/ilhamgunawan/lms-api-v2/db"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	UserID       string `db:"user_id" json:"user_id"`
	Username     string `db:"user_name" json:"user_name"`
	PasswordHash string `db:"psw_hash" json:"-"`
	FirstName    string `db:"first_name" json:"first_name"`
	LastName     string `db:"last_name" json:"last_name"`
	BirthDate    string `db:"date_of_birth" json:"date_of_birth"`
	Gender       string `db:"gender" json:"gender"`
}

func GetUserByUsername(username string) (user User, err error) {
	err = db.GetDB().SelectOne(&user,
		"SELECT ul.user_id, ul.user_name, ul.psw_hash, ua.first_name, ua.last_name, ua.date_of_birth, ua.gender FROM user_login_data ul LEFT JOIN user_account ua ON ul.user_id=ua.id WHERE user_name=$1",
		username)

	if err != nil {
		return user, err
	}

	return user, nil
}

type UserAccount struct {
	ID        string `db:"id" json:"id"`
	FirstName string `db:"first_name" json:"first_name"`
	LastName  string `db:"last_name" json:"last_name"`
	BirthDate string `db:"date_of_birth" json:"date_of_birth"`
	Gender    string `db:"gender" json:"gender"`
}

func GetUsers(offset, limit int) (users []UserAccount, err error) {
	_, err = db.GetDB().Select(&users, "SELECT * FROM user_account OFFSET $1 LIMIT $2", offset, limit)

	if err != nil {
		fmt.Println("err", err)
		return users, err
	}

	return users, nil
}

func CountUsers() (count int64, err error) {
	count, err = db.GetDB().SelectInt("SELECT COUNT(*) FROM user_account")

	if err != nil {
		return count, err
	}

	return count, nil
}

func GetUserById(id string) (user UserAccount, err error) {
	err = db.GetDB().SelectOne(&user,
		"SELECT ua.id, ua.first_name, ua.last_name, ua.date_of_birth, ua.gender FROM user_account ua WHERE id=$1",
		id)

	if err != nil {
		return user, err
	}

	return user, nil
}

type CreateUserBody struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Gender    string `json:"gender"`
	BirthDate string `json:"date_of_birth"`
	Username  string `json:"user_name"`
	Password  string `json:"password"`
}

func CreateUser(ua db.UserAccount, ul db.UserLoginData, plainPassword string) (user db.UserAccount, err error) {
	ua.ID = uuid.NewString()

	trx, err := db.GetDB().Begin()

	if err != nil {
		return user, err
	}

	// Insert into user_Account
	err = trx.Insert(&ua)

	if err != nil {
		trx.Rollback()
		return user, err
	}

	// Generate password hash
	plainPsw := []byte(plainPassword)
	hashedPsw, err := bcrypt.GenerateFromPassword(plainPsw, 4)

	if err != nil {
		return user, err
	}

	ul.ID = uuid.NewString()
	ul.UserId = ua.ID
	ul.PasswordHash = string(hashedPsw)

	// Insert into user_login_data
	err = trx.Insert(&ul)

	if err != nil {
		trx.Rollback()
		return user, err
	}

	err = trx.Commit()

	if err != nil {
		trx.Rollback()
		return user, err
	}

	user.ID = ua.ID
	user.FirstName = ua.FirstName
	user.LastName = ua.LastName
	user.Gender = ua.Gender
	user.BirthDate = ua.BirthDate

	return user, nil
}

func DeleteUser(userId string) (user db.UserAccount, err error) {
	trx, err := db.GetDB().Begin()

	if err != nil {
		return user, err
	}

	err = trx.SelectOne(&user, "SELECT * FROM user_account WHERE id=$1", userId)

	if err != nil {
		trx.Rollback()
		return user, err
	}

	userLoginData := db.UserLoginData{}

	err = trx.SelectOne(&userLoginData, "SELECT * FROM user_login_data WHERE user_id=$1", userId)

	if err != nil {
		trx.Rollback()
		return user, err
	}

	// Delete user login data
	_, err = trx.Exec("DELETE FROM user_login_data WHERE user_id=$1", userId)

	if err != nil {
		trx.Rollback()
		return user, err
	}

	// Delete user account
	_, err = trx.Exec("DELETE FROM user_account WHERE id=$1", userId)

	if err != nil {
		trx.Rollback()
		return user, err
	}

	err = trx.Commit()

	if err != nil {
		trx.Rollback()
		return user, err
	}

	return user, nil
}
