package data

import (
	"log"
	"time"
)

type Thread struct {
	Id        int
	Uuid      string
	Topic     string
	UserId    int
	CreatedAt time.Time
}

type Post struct {
	Id        int
	Uuid      string
	Body      string
	UserId    int
	ThreadId  int
	CreatedAt time.Time
}

func Threads() (threads []Thread, err error) {
	rows, err := Db.Query("SELECT id,uuid,topic,user_id,created_at FROM threads ORDER BY created_at DESC")
	if err != nil {
		log.Fatal(err)
		return
	}
	for rows.Next() {
		th := Thread{}
		if err = rows.Scan(&th.Id, &th.Uuid, &th.Topic, &th.UserId, &th.CreatedAt); err != nil {
			log.Fatal(err)
			return
		}
		threads = append(threads, th)
	}
	rows.Close()
	return
}

func (thread *Thread) NumberReplies() (count int) {
	rows, err := Db.Query("SELECT count(*) FROM posts where thread_id = $1", thread.Id)
	if err != nil {
		log.Fatal(err)
		return
	}
	for rows.Next() {
		if err = rows.Scan(&count); err != nil {
			log.Fatal(err)
			return
		}
	}
	rows.Close()
	return
}

func (user *User) CreateThread(topic string) (conv Thread, err error) {
	statement := "INSERT INTO threads(uuid,topic,user_id,created_at) values($1,$2,$3,$4) returning id,uuid,topic,user_id,created_at"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(createUUID(), topic, user.Id, time.Now()).Scan(&conv.Id, &conv.Uuid, &conv.Topic, &conv.UserId, &conv.CreatedAt)
	return
}

func ThreadByUUID(uuid string) (conv Thread, err error) {
	err = Db.QueryRow("SELECT id,uuid,topic,user_id,created_at FROM threads WHERE uuid = $1", uuid).
		Scan(&conv.Id, &conv.Uuid, &conv.Topic, &conv.UserId, &conv.CreatedAt)
	return
}

func (user *User) CreatePost(conv Thread, body string) (post Post, err error) {
	statement := "insert into posts (uuid, body, user_id, thread_id, created_at) values ($1, $2, $3, $4, $5) returning id, uuid, body, user_id, thread_id, created_at"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	// use QueryRow to return a row and scan the returned id into the Session struct
	err = stmt.QueryRow(createUUID(), body, user.Id, conv.Id, time.Now()).Scan(&post.Id, &post.Uuid, &post.Body, &post.UserId, &post.ThreadId, &post.CreatedAt)
	return
}
