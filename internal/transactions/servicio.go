package transactions

import "github.com/JoaquinRibero/go-web/internal/domain"

type Service interface {
	GetAll() ([]domain.Transaction, error)
	NewUser(codigo string, moneda string, monto int, emisor string, receptor string, fecha string) (domain.Transaction, error)
	Update(id int, codigo string, moneda string, monto int, emisor string, receptor string, fecha string) (domain.Transaction, error)
	Delete(id int) error
	UpdateCodigoAndMonto(id int, codigo string, monto int) (domain.Transaction, error)
}

type service struct {
	repo Repository
}

func NewService(s Repository) Service {
	return &service{
		repo: s,
	}
}

func (s *service) GetAll() ([]domain.Transaction, error) {
	ts, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return ts, nil
}

func (s *service) NewUser(codigo string, moneda string, monto int, emisor string, receptor string, fecha string) (domain.Transaction, error) {
	lastId, err := s.repo.LastId()
	if err != nil {
		return domain.Transaction{}, err
	}
	lastId++
	ts, err := s.repo.NewUser(lastId, codigo, moneda, monto, emisor, receptor, fecha)
	return ts, err
}

func (s *service) Update(id int, codigo string, moneda string, monto int, emisor string, receptor string, fecha string) (domain.Transaction, error) {
	t, err := s.repo.Update(id, codigo, moneda, monto, emisor, receptor, fecha)
	return t, err
}

func (s *service) Delete(id int) error {
	err := s.repo.Delete(id)
	return err
}

func (s *service) UpdateCodigoAndMonto(id int, codigo string, monto int) (domain.Transaction, error) {
	t, err := s.repo.UpdateCodigoAndMonto(id, codigo, monto)
	return t, err
}
