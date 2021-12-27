package transactions

type Service interface {
	GetAll() []Transaction
	NewUser(codigo string, moneda string, monto int, emisor string, receptor string, fecha string) []Transaction
}

type service struct {
	repo Repository
}

func NewService(s Repository) Service {
	return &service{
		repo: s,
	}
}

func (s *service) GetAll() []Transaction {
	ts := s.repo.GetAll()
	return ts
}

func (s *service) NewUser(codigo string, moneda string, monto int, emisor string, receptor string, fecha string) []Transaction {
	ts := s.repo.NewUser(codigo, moneda, monto, emisor, receptor, fecha)
	return ts
}
