package model

type User struct {
	Id        string `db:"id"`
	NickName  string `db:"nick_name"`
	LoginName string `db:"login_name"`
	LoginPass string `db:"login_pass"`
	IsLocking bool   `db:"is_locking"`
}
