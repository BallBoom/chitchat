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

// Get the session for an existing user
func (user User) Session() (s Session, err error) {
	statment := "select uuid, user_id, email, created_at from session where user_id=?"
	stout, err := Db.Prepare(statment)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	defer stout.Close()

	err = stout.QueryRow(user.ID).Scan(&s.UUID, &s.UserID, &s.Email, &s.CreatedAt)
	return
}

// create new user
func (u *User) Create() (user User, err error) {
	statment := "insert into user(uuid, name, email, password, created_at) values(?,?,?,?,?)"
	stin, err := Db.Prepare(statment)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	defer stin.Close()

	uuid := createUUID()
	stin.Exec(uuid, u.Name, u.Email, Encrypt(u.Password), time.Now())

	stout, err := Db.Prepare("select id, uuid, email, name, password, create_at from user where uuid=?")
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	defer stout.Close()

	err = stout.QueryRow(uuid).Scan(&user.ID, &user.UUID, &user.Email, &user.Name, &user.Password,
		&user.CreatedAt)
	return
}

// Delete user from database
func (user *User) Delete() (err error) {
	statement := "delete from users where id = ?"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.ID)
	return
}

// Update user information in the database
func (user *User) Update() (err error) {
	statement := "update users set name = ?, email = ? where id = ?"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Name, user.Email, user.ID)
	return
}

// Delete all users from database
func UserDeleteAll() (err error) {
	statement := "delete from users"
	_, err = Db.Exec(statement)
	return
}

// Get all users in the database and returns it
func Users() (users []User, err error) {
	rows, err := Db.Query("SELECT id, uuid, name, email, password, created_at FROM users")
	if err != nil {
		return
	}
	for rows.Next() {
		user := User{}
		if err = rows.Scan(&user.ID, &user.UUID, &user.Name, &user.Email, &user.Password, &user.CreatedAt); err != nil {
			return
		}
		users = append(users, user)
	}
	rows.Close()
	return
}

// Get a single user given the email
func UserByEmail(email string) (user User, err error) {
	user = User{}
	err = Db.QueryRow("SELECT id, uuid, name, email, password, created_at FROM users WHERE email = ?", email).
		Scan(&user.ID, &user.UUID, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	return
}

// Get a single user given the UUID
func UserByUUID(uuid string) (user User, err error) {
	user = User{}
	err = Db.QueryRow("SELECT id, uuid, name, email, password, created_at FROM users WHERE uuid = ?", uuid).
		Scan(&user.ID, &user.UUID, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	return
}
