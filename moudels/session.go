package moudels

import (
	"fmt"
	"time"
)

type Session struct {
	ID        int
	UUID      string
	UserID    int
	Email     string
	CreatedAt time.Time
}

// 检查session是否存在
func (s *Session) Check() (valid bool, err error) {
	err = Db.QueryRow("select id, uuid, email, created_at from sessions where uuid=?", s.UUID).
		Scan(&s.ID, &s.UUID, &s.Email, &s.CreatedAt)
	if err != nil {
		valid = false
		// log.Fatal(err.Error())
		fmt.Println(err.Error())
		return
	}
	if s.ID != 0 {
		valid = true
	}
	return
}

// UUID删除session
func (s *Session) DeleteByUUID() (err error) {
	stment := "delete from sessions where uuid=?"
	st, err := Db.Prepare(stment)
	if err != nil {
		// log.Fatal(err.Error())
		fmt.Println(err.Error())
		return
	}
	defer st.Close()

	_, err = st.Exec(s.UUID)
	return
}

// 获取某个session的用户
func (s *Session) User() (user User, err error) {
	user = User{}
	err = Db.QueryRow("select id,uuid, email, created_at from user where id =?", s.UserID).
		Scan(&user.ID, &user.UUID, &user.Email, &user.CreatedAt)
	if err != nil {
		// log.Fatal(err.Error())
		fmt.Println(err.Error())
	}
	return
}

// 删除所有session
func DeleteAllSession() (err error) {
	stmt := "delete from sessions"
	_, err = Db.Prepare(stmt)

	return
}
