package incoming_service

import (
	"controle-de-gastos/src/model"
	"controle-de-gastos/src/repository/incoming_repo"
)

type incomingService struct {
	repo incoming_repo.IncomingRepo
}

type IncomingService interface {
	GetAll() ([]model.Incoming, error)
	GetById(id int) (*model.Incoming, error)
	Create(incomingCreate model.IncomingCreateRequest) (int, error)
	Update(id int, incoming model.Incoming) error
	Delete(id int) error
}

func NewIncomingService(repo incoming_repo.IncomingRepo) IncomingService {
	return &incomingService{repo: repo}
}

func (s *incomingService) GetAll() ([]model.Incoming, error) {
	return s.repo.GetAll()
}

func (s *incomingService) GetById(id int) (*model.Incoming, error) {
	return s.repo.GetById(id)
}

func (s *incomingService) Create(incomingCreate model.IncomingCreateRequest) (int, error) {
	var incoming model.Incoming
	incoming.Value = incomingCreate.Value
	incoming.Description = incomingCreate.Description
	incoming.IncomeDate = incomingCreate.IncomeDate
	incoming.UserId = incomingCreate.UserId
	return s.repo.Create(incoming)
}

func (s *incomingService) Update(id int, incoming model.Incoming) error {
	incoming.Id = id
	return s.repo.Update(incoming)
}

func (s *incomingService) Delete(id int) error {
	return s.repo.Delete(id)
}