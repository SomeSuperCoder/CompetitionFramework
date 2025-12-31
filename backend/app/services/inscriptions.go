package services

import (
	"net/http"

	"github.com/SomeSuperCoder/CompetitionFramework/backend/repository"
)

type InscriptionService struct {
	Repo *repository.Queries
}

func (s *InscriptionService) Insert(r *http.Request, args *repository.InsertInscriptionParams, reply *repository.Inscription) error {
	inscription, err := s.Repo.InsertInscription(r.Context(), *args)
	if err != nil {
		return err
	}
	*reply = inscription
	return nil
}
