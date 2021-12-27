package transactions

type Service interface {
	GetAll() []Transaction
	NewUser(codigo string, moneda string, monto int, emisor string, receptor string, fecha string) []Transaction
	Update(id int, codigo string, moneda string, monto int, emisor string, receptor string, fecha string) (Transaction, error)
	Delete(id int) error
	UpdateCodigoAndMonto(id int, codigo string, monto int) (Transaction, error)
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

func (s *service) Update(id int, codigo string, moneda string, monto int, emisor string, receptor string, fecha string) (Transaction, error) {
	t, err := s.repo.Update(id, codigo, moneda, monto, emisor, receptor, fecha)
	return t, err
}

func (s *service) Delete(id int) error {
	err := s.repo.Delete(id)
	return err
}

func (s *service) UpdateCodigoAndMonto(id int, codigo string, monto int) (Transaction, error) {
	t, err := s.repo.UpdateCodigoAndMonto(id, codigo, monto)
	return t, err
}
