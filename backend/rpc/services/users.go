package services

import (
	"net/http"

	"github.com/SomeSuperCoder/CompetitionFramework/backend/repository"
)

type UserService struct {
	Repo *repository.Queries
}

func (s *UserService) Insert(r *http.Request, args *repository.InsertUserParams, reply *repository.InsertUserRow) error {
	user, err := s.Repo.InsertUser(r.Context(), *args)
	if err != nil {
		return err
	}
	*reply = user
	return nil
}
