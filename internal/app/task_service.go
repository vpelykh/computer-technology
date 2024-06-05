package app

import (
	"log"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/database"
)

type TaskService interface {
	Save(t domain.Task) (domain.Task, error)
	GetForUser(uId uint64) ([]domain.Task, error)
}

type taskService struct {
	taskRepo database.TaskRepository
}

func NewTaskService(tr database.TaskRepository) TaskService {
	return taskService{
		taskRepo: tr,
	}
}

func (s taskService) Save(t domain.Task) (domain.Task, error) {
	task, err := s.taskRepo.Save(t)
	if err != nil {
		log.Printf("TaskService -> Save: %s", err)
		return domain.Task{}, err
	}
	return task, nil
}

func (s taskService) GetForUser(uId uint64) ([]domain.Task, error) {
	tasks, err := s.taskRepo.GetByUserId(uId)
	if err != nil {
		log.Printf("TaskService -> GetForUser: %s", err)
		return nil, err
	}
	return tasks, nil
}
func (s taskService) UpdateByTaskId(t domain.Task) (domain.Task, error) {
	updatedTask, err := s.taskRepo.UpdateByTaskId(t)
	if err != nil {
	  log.Printf("TaskService -> UpdateByTaskId %s", err)
	  return domain.Task{}, err
	}
	return updatedTask, nil
}