package moudels

import (
	"log"
	"time"
)

type User struct {
	ID        int
	UUID      string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
}

// 将生成的用户生成一条session
func (user User) CreateSession() (session Session, err error) {
	statment := "insert into session (uuid, email, user_id,created_at) values(?,?,?,?)"
	stin, err := Db.Prepare(statment)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	defer stin.Close()

	// 生成UUID
	uuid := createUUID()
	// 插入session
	stin.Exec(uuid, user.Email, user.ID, time.Now())

	// 将插入的session查询出来
	stout, err := Db.Prepare("select id, uuid, email, user_id, created_at from session where uuid=?")
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	defer stout.Close()

	err = stout.QueryRow(uuid).Scan(&session.ID, &session.UUID, &session.Email, &session.UserID, &session.CreatedAt)
	return

}
