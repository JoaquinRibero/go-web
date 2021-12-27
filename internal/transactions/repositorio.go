package transactions

import (
	"fmt"

	"github.com/JoaquinRibero/go-web/internal/domain"
	"github.com/JoaquinRibero/go-web/pkg/store"
)

type Repository interface {
	GetAll() []domain.Transaction
	NewUser(id int, codigo string, moneda string, monto int, emisor string, receptor string, fecha string) ([]domain.Transaction, error)
	Update(id int, codigo string, moneda string, monto int, emisor string, receptor string, fecha string) (domain.Transaction, error)
	Delete(id int) error
	UpdateCodigoAndMonto(id int, codigo string, monto int) (domain.Transaction, error)
	LastId() (int, error)
}

type repository struct {
	db store.Store
}

func NewRepository(db store.Store) Repository {
	return &repository{
		db: db,
	}
}

func (repo *repository) GetAll() []domain.Transaction {
	var ts []domain.Transaction
	repo.db.Read(&ts)
	return ts
}

func (repo *repository) LastId() (int, error) {
	var ts []domain.Transaction
	if err := repo.db.Read(&ts); err != nil {
		return 0, err
	}
	if len(ts) == 0 {
		return 0, nil
	}
	return ts[len(ts)-1].Id, nil
}

func (repo *repository) NewUser(id int, codigo string, moneda string, monto int, emisor string, receptor string, fecha string) ([]domain.Transaction, error) {

	var ts []domain.Transaction
	repo.db.Read(&ts)
	t := domain.Transaction{id, codigo, moneda, monto, emisor, receptor, fecha, true}
	ts = append(ts, t)
	if err := repo.db.Write(ts); err != nil {
		return []domain.Transaction{}, err
	}
	return ts, nil
}

func (repo *repository) Update(id int, codigo string, moneda string, monto int, emisor string, receptor string, fecha string) (domain.Transaction, error) {
	t := domain.Transaction{
		Codigo:   codigo,
		Moneda:   moneda,
		Monto:    monto,
		Emisor:   emisor,
		Receptor: receptor,
		Fecha:    fecha,
	}
	updated := false
	var ts []domain.Transaction
	repo.db.Read(&ts)
	for i := range ts {
		if ts[i].Id == id {
			t.Id = id
			ts[i] = t
			updated = true
		}
	}
	if !updated {
		return domain.Transaction{}, fmt.Errorf("transacción %d not found", id)
	}
	if err := repo.db.Write(ts); err != nil {
		return domain.Transaction{}, err
	}
	return t, nil
}

func (repo *repository) Delete(id int) error {
	deleted := false
	var ts []domain.Transaction
	repo.db.Read(&ts)
	for i := range ts {
		if ts[i].Id == id {
			deleted = true
		}
	}
	if !deleted {
		return fmt.Errorf("transaccion %d no encontrado", id)
	}
	ts[id].Estado = false
	if err := repo.db.Write(ts); err != nil {
		return err
	}
	return nil
}

func (repo *repository) UpdateCodigoAndMonto(id int, codigo string, monto int) (domain.Transaction, error) {
	var t domain.Transaction
	updated := false
	var ts []domain.Transaction
	repo.db.Read(&ts)
	for i := range ts {
		if ts[i].Id == id {
			ts[i].Codigo = codigo
			ts[i].Monto = monto
			updated = true
			t = ts[i]
		}
	}
	if !updated {
		return domain.Transaction{}, fmt.Errorf("transaccion %d no encontrada", id)
	}
	if err := repo.db.Write(ts); err != nil {
		return domain.Transaction{}, err
	}
	return t, nil
}
