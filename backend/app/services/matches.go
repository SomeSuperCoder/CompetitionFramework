package services

import (
	"net/http"

	"github.com/SomeSuperCoder/CompetitionFramework/backend/repository"
)

type MatchService struct {
	Repo *repository.Queries
}

func (s *MatchService) Insert(r *http.Request, args *repository.InsertMatchParams, reply *repository.Match) error {
	match, err := s.Repo.InsertMatch(r.Context(), *args)
	if err != nil {
		return err
	}
	*reply = match
	return nil
}

func (s *MatchService) FindAll(r *http.Request, args *any, reply *[]repository.Match) error {
	matches, err := s.Repo.FindAllMatches(r.Context())
	if err != nil {
		return err
	}
	*reply = matches
	return nil
}
