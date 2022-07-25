package service

import (
	"github.com/basterrus/http_rest_api_service_golang"
	"github.com/basterrus/http_rest_api_service_golang/pkg/repository"
)

type TodoItemService struct {
	repo     repository.TodoItem
	listRepo repository.TodoList
}

func NewTodoItemService(repo repository.TodoItem, listRepo repository.TodoList) *TodoItemService {
	return &TodoItemService{repo: repo, listRepo: listRepo}
}

func (s *TodoItemService) Create(userId, listId int, item http_rest_api_service_golang.TodoItem) (int, error) {
	_, err := s.listRepo.GetById(userId, listId)
	if err != nil {
		return 0, nil
	}
	return s.repo.Create(listId, item)
}

func (s *TodoItemService) Update(userId, itemId int, input http_rest_api_service_golang.UpdateItemInput) error {
	return s.repo.Update(userId, itemId, input)
}

func (s *TodoItemService) GetAll(userId, listId int) ([]http_rest_api_service_golang.TodoItem, error) {
	return s.repo.GetAll(userId, listId)
}

func (s *TodoItemService) GetById(userId, itemId int) (http_rest_api_service_golang.TodoItem, error) {
	return s.repo.GetById(userId, itemId)
}

func (s *TodoItemService) Delete(userId, itemId int) error {
	return s.repo.Delete(userId, itemId)
}
