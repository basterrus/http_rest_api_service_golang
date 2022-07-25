package repository

import (
	"github.com/basterrus/http_rest_api_service_golang"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user http_rest_api_service_golang.User) (int, error)
	GetUser(username, password string) (http_rest_api_service_golang.User, error)
}

type TodoList interface {
	Create(userId int, list http_rest_api_service_golang.TodoList) (int, error)
	GetAll(userId int) ([]http_rest_api_service_golang.TodoList, error)
	GetById(userId int, listId int) (http_rest_api_service_golang.TodoList, error)
	Update(userId, listId int, input http_rest_api_service_golang.UpdateListInput) error
	Delete(userId, listId int) error
}

type TodoItem interface {
	Create(listId int, item http_rest_api_service_golang.TodoItem) (int, error)
	GetAll(userId, listId int) ([]http_rest_api_service_golang.TodoItem, error)
	Update(userId, itemId int, input http_rest_api_service_golang.UpdateItemInput) error
	GetById(userId, itemId int) (http_rest_api_service_golang.TodoItem, error)
	Delete(userId, itemId int) error
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList:      NewTodoListPostgres(db),
		TodoItem:      NewTodoItemPostgres(db),
	}
}
