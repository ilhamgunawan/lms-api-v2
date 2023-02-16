package models

import "github.com/ilhamgunawan/lms-api-v2/db"

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
