package todo_app

type User struct {
	Id       int    `json:"-" db:"id"`
	Name     string `json:"name"     binding:"required" example:"Sergey"`
	Username string `json:"username" binding:"required" example:"nekrasov"`
	Password string `json:"password" binding:"required" example:"qwerty"`
}
