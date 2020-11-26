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
	statment := "insert into sessions (uuid, email, user_id,created_at) values(?,?,?,?)"
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
	stout, err := Db.Prepare("select id, uuid, email, user_id, created_at from sessions where uuid=?")
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	defer stout.Close()

	err = stout.QueryRow(uuid).Scan(&session.ID, &session.UUID, &session.Email, &session.UserID, &session.CreatedAt)
	return
}

// Get the session for an existing user 检查某个用户是否存在session
func (user User) Session() (s Session, err error) {
	statment := "select uuid, user_id, email, created_at from sessions where user_id=?"
	stout, err := Db.Prepare(statment)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	defer stout.Close()

	err = stout.QueryRow(user.ID).Scan(&s.UUID, &s.UserID, &s.Email, &s.CreatedAt)
	return
}

// create new user 创建新用户
func (u *User) Create() (user User, err error) {
	statment := "insert into users(uuid, name, email, password, created_at) values(?,?,?,?,?)"
	stin, err := Db.Prepare(statment)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	defer stin.Close()

	uuid := createUUID()
	stin.Exec(uuid, u.Name, u.Email, Encrypt(u.Password), time.Now())

	stout, err := Db.Prepare("select id, uuid, email, name, password, created_at from users where uuid=?")
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	defer stout.Close()

	err = stout.QueryRow(uuid).Scan(&user.ID, &user.UUID, &user.Email, &user.Name, &user.Password,
		&user.CreatedAt)
	return
}

// Delete user from database  删除单个用户
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

// Update user information in the database  更新用户名 邮箱
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

// Delete all users from database  删除所有的用户
func UserDeleteAll() (err error) {
	statement := "delete from users"
	_, err = Db.Exec(statement)
	return
}

// Get all users in the database and returns it 获取所有的用户
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

// Get a single user given the email   使用邮箱获取用户
func UserByEmail(email string) (user User, err error) {
	user = User{}
	err = Db.QueryRow("SELECT id, uuid, name, email, password, created_at FROM users WHERE email = ?", email).
		Scan(&user.ID, &user.UUID, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	return
}

// Get a single user given the UUID 获取用户 使用UUID
func UserByUUID(uuid string) (user User, err error) {
	user = User{}
	err = Db.QueryRow("SELECT id, uuid, name, email, password, created_at FROM users WHERE uuid = ?", uuid).
		Scan(&user.ID, &user.UUID, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	return
}

// Create a new thread  该用户创建新主题
func (user *User) CreateThread(topic string) (conv Thread, err error) {
	statement := "insert into threads (uuid, topic, user_id, created_at) values (?, ?, ?, ?)"
	stmtin, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmtin.Close()

	uuid := createUUID()
	stmtin.Exec(uuid, topic, user.ID, time.Now())

	stmtout, err := Db.Prepare("select id, uuid, topic, user_id, created_at from threads where uuid = ?")
	if err != nil {
		return
	}
	defer stmtout.Close()

	// use QueryRow to return a row and scan the returned id into the Session struct
	err = stmtout.QueryRow(uuid).Scan(&conv.ID, &conv.UUID, &conv.Topic, &conv.UserID, &conv.CreatedAt)
	return
}

// Create a new post to a thread 在主题下创建新贴子
func (user *User) CreatePost(conv Thread, body string) (post Post, err error) {
	statement := "insert into posts (uuid, body, user_id, thread_id, created_at) values (?, ?, ?, ?, ?)"
	stmtin, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmtin.Close()

	uuid := createUUID()
	stmtin.Exec(uuid, body, user.ID, conv.ID, time.Now())

	stmtout, err := Db.Prepare("select id, uuid, body, user_id, thread_id, created_at from posts where uuid = ?")
	if err != nil {
		return
	}
	defer stmtout.Close()

	// use QueryRow to return a row and scan the returned id into the Session struct
	err = stmtout.QueryRow(uuid).Scan(&post.ID, &post.UUID, &post.Body, &post.UserID, &post.ThreadID, &post.CreatedAt)
	return
}
