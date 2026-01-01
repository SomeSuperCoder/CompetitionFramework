package services

import (
	"net/http"

	"github.com/SomeSuperCoder/CompetitionFramework/backend/repository"
)

type CompetitionService struct {
	Repo *repository.Queries
}

func (s *CompetitionService) Insert(r *http.Request, args *repository.InsertCompetitionParams, reply *repository.Competition) error {
	competition, err := s.Repo.InsertCompetition(r.Context(), *args)
	if err != nil {
		return err
	}
	*reply = competition
	return nil
}

func (s *CompetitionService) FindAll(r *http.Request, args *any, reply *[]repository.Competition) error {
	competitions, err := s.Repo.FindAllCompetitions(r.Context())
	if err != nil {
		return err
	}
	*reply = competitions
	return nil
}

func (s *CompetitionService) Rename(r *http.Request, args *repository.RenameCompetitionParams, reply *repository.Competition) error {
	competition, err := s.Repo.RenameCompetition(r.Context(), *args)
	if err != nil {
		return err
	}
	*reply = competition
	return nil
}
