package services

import (
	"net/http"

	"github.com/SomeSuperCoder/CompetitionFramework/backend/repository"
)

type TaskService struct {
	Repo *repository.Queries
}

func (s *TaskService) Insert(r *http.Request, args *repository.InsertTaskParams, reply *repository.Task) error {
	task, err := s.Repo.InsertTask(r.Context(), *args)
	if err != nil {
		return err
	}
	*reply = task
	return nil
}

func (s *TaskService) FindAll(r *http.Request, args *any, reply *[]repository.Task) error {
	tasks, err := s.Repo.FindAllTasks(r.Context())
	if err != nil {
		return err
	}
	*reply = tasks
	return nil
}

func (s *TaskService) Update(r *http.Request, args *repository.UpdateTaskParams, reply *repository.Task) error {
	task, err := s.Repo.UpdateTask(r.Context(), *args)
	if err != nil {
		return nil
	}
	*reply = task
	return nil
}

func (s *TaskService) Delete(r *http.Request, args *repository.DeleteTaskParams, reply *repository.Task) error {
	task, err := s.Repo.DeleteTask(r.Context(), *args)
	if err != nil {
		return err
	}
	*reply = task
	return nil
}
