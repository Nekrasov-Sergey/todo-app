package todo_app

import "errors"

type TodoList struct {
	Id          int    `json:"id" db:"id" example:"0"`
	Title       string `json:"title" db:"title" binding:"required" example:"ДЗ"`
	Description string `json:"description" db:"description" example:"На числитель"`
}

type UsersList struct {
	Id     int
	UserId int
	ListId int
}

type TodoItem struct {
	Id          int    `json:"id" db:"id" example:"0"`
	Title       string `json:"title" db:"title" binding:"required" example:"Этика"`
	Description string `json:"description" db:"description" example:"Реферат"`
	Done        bool   `json:"done" db:"done" example:"false"`
}

type ListsItem struct {
	Id     int
	ListId int
	ItemId int
}

type UpdateListInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
}

func (i UpdateListInput) Validate() error {
	if i.Title == nil && i.Description == nil {
		return errors.New("update structure has no values")
	}

	return nil
}

type UpdateItemInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Done        *bool   `json:"done"`
}

func (i UpdateItemInput) Validate() error {
	if i.Title == nil && i.Description == nil && i.Done == nil {
		return errors.New("update structure has no values")
	}

	return nil
}
