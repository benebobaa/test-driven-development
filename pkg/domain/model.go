package domain

type User struct {
	ID       int    `common:"id"`
	Username string `common:"username"`
	Password string `common:"password"`
}
