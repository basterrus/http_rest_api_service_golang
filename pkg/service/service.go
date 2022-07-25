package service

import (
	"github.com/basterrus/http_rest_api_service_golang"
	"github.com/basterrus/http_rest_api_service_golang/pkg/repository"
)

type Authorization interface {
	CreateUser(user http_rest_api_service_golang.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	Create(userId int, list http_rest_api_service_golang.TodoList) (int, error)
	GetAll(userId int) ([]http_rest_api_service_golang.TodoList, error)
	GetById(userId, listId int) (http_rest_api_service_golang.TodoList, error)
	Update(userId, listId int, input http_rest_api_service_golang.UpdateListInput) error
	Delete(userId, listId int) error
}

type TodoItem interface {
	Create(userId, listId int, item http_rest_api_service_golang.TodoItem) (int, error)
	GetAll(userId, listId int) ([]http_rest_api_service_golang.TodoItem, error)
	Update(userId, itemId int, input http_rest_api_service_golang.UpdateItemInput) error
	GetById(userId, itemId int) (http_rest_api_service_golang.TodoItem, error)
	Delete(userId, itemId int) error
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TodoList:      NewTodoListService(repos.TodoList),
		TodoItem:      NewTodoItemService(repos.TodoItem, repos.TodoList),
	}
}
