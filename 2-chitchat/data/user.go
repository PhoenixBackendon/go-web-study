package data

import (
	"time"
)

type User struct {
	Id        int
	Uuid      string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
}

type Session struct {
	Id        int
	Uuid      string
	Email     string
	UserId    int
	CreatedAt time.Time
}

func UserByEmail(email string) (user User, err error) {
	rows, err := Db.Query("SELECT id,uuid,name,email,password,created_at FROM users WHERE email = $1", email)
	if err != nil {
		return
	}
	for rows.Next() {
		if err = rows.Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt); err != nil {
			return
		}
	}
	return
}

func (user *User) CreateSession() (session Session, err error) {
	statement := "INSERT INTO sessions(uuid,email,user_id,created_at) values ($1,$2,$3,$4) returning id,uuid,email,user_id,created_at"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(createUUID(), user.Email, user.Id, time.Now()).Scan(&session.Id, &session.Uuid, &session.Email, &session.UserId, &session.CreatedAt)
	return
}

func (session *Session) Check() (valid bool, err error) {
	err = Db.QueryRow("SELECT id,uuid,email,user_id,created_at FROM sessions WHERE uuid = $1",
		session.Uuid).Scan(&session.Id, &session.Uuid, &session.Email, &session.UserId, &session.CreatedAt)
	if err != nil {
		valid = false
		return
	}
	if session.Id != 0 {
		valid = true
	}
	return
}

func (session *Session) DeleteByUuid() {
	statement := "DELETE FROM sessions WHERE uuid = $1"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(session.Uuid)
	return
}

func (user *User) Create() (err error) {
	statement := "INSERT INTO users(uuid,name,email,password,created_at) values ($1,$2,$3,$4,$5) returning id ,uuid,created_at"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(createUUID(), user.Name, user.Email, Encrypt(user.Password), time.Now()).Scan(&user.Id, &user.Uuid, &user.CreatedAt)
	return
}

func (session *Session) User() (user User, err error) {
	user = User{}
	err = Db.QueryRow("SELECT id,uuid,name,email,created_at FROM users WHERE id = $1", session.UserId).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.CreatedAt)
	return
}
