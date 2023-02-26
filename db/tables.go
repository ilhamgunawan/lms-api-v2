package db

type UserAccount struct {
	ID        string `db:"id, primarykey" json:"id"`
	FirstName string `db:"first_name" json:"first_name"`
	LastName  string `db:"last_name" json:"last_name"`
	BirthDate string `db:"date_of_birth" json:"date_of_birth"`
	Gender    string `db:"gender" json:"gender"`
}

type UserLoginData struct {
	ID           string `db:"id, primarykey" json:"id"`
	UserId       string `db:"user_id" json:"user_id"`
	Username     string `db:"user_name" json:"user_name"`
	PasswordHash string `db:"psw_hash" json:"psw_hash"`
}

type Course struct {
	ID          string  `db:"id, primarykey" json:"id"`
	Title       string  `db:"title" json:"title"`
	Slug        string  `db:"slug" json:"slug"`
	Description string  `db:"description" json:"description"`
	Status      string  `db:"status" json:"status"`
	Created     string  `db:"created" json:"created"`
	Updated     *string `db:"updated" json:"updated"`
	Deleted     *string `db:"deleted" json:"deleted"`
	AuthorId    string  `db:"author_id" json:"author_id"`
}

type Topic struct {
	ID       string  `db:"id, primarykey" json:"id"`
	CourseId string  `db:"course_id" json:"course_id"`
	Title    string  `db:"title" json:"title"`
	Status   string  `db:"status" json:"status"`
	Created  string  `db:"created" json:"created"`
	Updated  *string `db:"updated" json:"updated"`
	Deleted  *string `db:"deleted" json:"deleted"`
	AuthorId string  `db:"author_id" json:"author_id"`
}

type Mission struct {
	ID         string  `db:"id, primarykey" json:"id"`
	TopicId    string  `db:"topic_id" json:"topic_id"`
	TypeId     string  `db:"type_id" json:"type_id"`
	Title      string  `db:"title" json:"title"`
	Status     string  `db:"status" json:"status"`
	ContentUrl string  `db:"content_url" json:"content_url"`
	Created    string  `db:"created" json:"created"`
	Updated    *string `db:"updated" json:"updated"`
	Deleted    *string `db:"deleted" json:"deleted"`
	AuthorId   string  `db:"author_id" json:"author_id"`
}

type MissionType struct {
	ID          string `db:"id, primarykey" json:"id"`
	Name        string `db:"name" json:"name"`
	Description string `db:"description" json:"description"`
}

type MissionUser struct {
	ID        string  `db:"id, primarykey" json:"id"`
	UserId    string  `db:"user_id" json:"user_id"`
	MissionId string  `db:"mission_id" json:"mission_id"`
	Status    string  `db:"status" json:"status"`
	Created   string  `db:"created" json:"created"`
	Updated   *string `db:"updated" json:"updated"`
	Duration  int     `db:"duration" json:"duration"`
}

type TopicUser struct {
	ID       string  `db:"id, primarykey" json:"id"`
	UserId   string  `db:"user_id" json:"user_id"`
	TopicId  string  `db:"topic_id" json:"topic_id"`
	Status   string  `db:"status" json:"status"`
	Created  string  `db:"created" json:"created"`
	Updated  *string `db:"updated" json:"updated"`
	Duration int     `db:"duration" json:"duration"`
}

type CourseUser struct {
	ID       string  `db:"id, primarykey" json:"id"`
	UserId   string  `db:"user_id" json:"user_id"`
	CourseId string  `db:"course_id" json:"course_id"`
	Status   string  `db:"status" json:"status"`
	Created  string  `db:"created" json:"created"`
	Updated  *string `db:"updated" json:"updated"`
	Duration int     `db:"duration" json:"duration"`
}
