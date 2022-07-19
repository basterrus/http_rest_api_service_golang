package service

import (
	"github.com/basterrus/http_rest_api_service_golang"
	"github.com/basterrus/http_rest_api_service_golang/pkg/repository"
)

type TodoListService struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{repo: repo}
}

func (s *TodoListService) Create(userId int, list http_rest_api_service_golang.TodoList) (int, error) {
	return s.repo.Create(userId, list)
}

func (s *TodoListService) GetAll(userId int) ([]http_rest_api_service_golang.TodoList, error) {
	return s.repo.GetAll(userId)
}

func (s *TodoListService) GetById(userId, listId int) (http_rest_api_service_golang.TodoList, error) {
	return s.repo.GetById(userId, listId)
}

func (s *TodoListService) Update(userId, listId int, input http_rest_api_service_golang.UpdateListInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(userId, listId, input)
}

func (s *TodoListService) Delete(userId, listId int) error {
	return s.repo.Delete(userId, listId)
}
