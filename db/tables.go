package db

type UserAccount struct {
	ID        string `db:"id" json:"id"`
	FirstName string `db:"first_name" json:"first_name"`
	LastName  string `db:"last_name" json:"last_name"`
	BirthDate string `db:"date_of_birth" json:"date_of_birth"`
	Gender    string `db:"gender" json:"gender"`
}

type UserLoginData struct {
	ID           string `db:"id" json:"id"`
	UserId       string `db:"user_id" json:"user_id"`
	Username     string `db:"user_name" json:"user_name"`
	PasswordHash string `db:"psw_hash" json:"psw_hash"`
}
