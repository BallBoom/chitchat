package moudels

import "time"

type Thread struct {
	ID        int
	UUID      string
	Topic     string
	UserID    int
	CreatedAt time.Time
}

func (t *Thread) CreatedAtDate() string {
	return t.CreatedAt.Format("Jan 2, 2006 at 3:04pm")
}

// get the number of posts in a thread 返回主题下的所有文章的数量
func (thread *Thread) NumReplies() (count int) {
	rows, err := Db.Query("SELECT count(*) FROM posts where thread_id = ?", thread.ID)
	if err != nil {
		return
	}
	for rows.Next() {
		if err = rows.Scan(&count); err != nil {
			return
		}
	}
	rows.Close()
	return
}

// get posts to a thread 获取主题的所有文章
func (thread *Thread) Posts() (posts []Post, err error) {
	rows, err := Db.Query("SELECT id, uuid, body, user_id, thread_id, created_at FROM posts where thread_id = ?", thread.ID)
	if err != nil {
		return
	}
	for rows.Next() {
		post := Post{}
		if err = rows.Scan(&post.ID, &post.UUID, &post.Body, &post.UserID, &post.ThreadID, &post.CreatedAt); err != nil {
			return
		}
		posts = append(posts, post)
	}
	rows.Close()
	return
}

// Get all threads in the database and returns it 获取所有的主题
func Threads() (threads []Thread, err error) {
	rows, err := Db.Query("SELECT id, uuid, topic, user_id, created_at FROM threads ORDER BY created_at DESC")
	if err != nil {
		return
	}
	for rows.Next() {
		conv := Thread{}
		if err = rows.Scan(&conv.ID, &conv.UUID, &conv.Topic, &conv.UserID, &conv.CreatedAt); err != nil {
			return
		}
		threads = append(threads, conv)
	}
	rows.Close()
	return
}

// Get a thread by the UUID 通过UUID获取主题
func ThreadByUUID(uuid string) (conv Thread, err error) {
	conv = Thread{}
	err = Db.QueryRow("SELECT id, uuid, topic, user_id, created_at FROM threads WHERE uuid = ?", uuid).
		Scan(&conv.ID, &conv.UUID, &conv.Topic, &conv.UserID, &conv.CreatedAt)
	return
}

// Get the user who started this thread  获取该主题的发起人
func (thread *Thread) User() (user User) {
	user = User{}
	Db.QueryRow("SELECT id, uuid, name, email, created_at FROM users WHERE id = ?", thread.UserID).
		Scan(&user.ID, &user.UUID, &user.Name, &user.Email, &user.CreatedAt)
	return
}
