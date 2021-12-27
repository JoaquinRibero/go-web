package transactions

import (
	"fmt"
)

type Transaction struct {
	Id       int    `json:"_id"`
	Codigo   string `json:"codigo"`
	Moneda   string `json:"moneda"`
	Monto    int    `json:"monto"`
	Emisor   string `json:"emisor"`
	Receptor string `json:"receptor"`
	Fecha    string `json:"fecha"`
	Estado   bool   `json:"estado"`
}

var transactions []Transaction

type Repository interface {
	GetAll() []Transaction
	NewUser(codigo string, moneda string, monto int, emisor string, receptor string, fecha string) []Transaction
	Update(id int, codigo string, moneda string, monto int, emisor string, receptor string, fecha string) (Transaction, error)
	Delete(id int) error
	UpdateCodigoAndMonto(id int, codigo string, monto int) (Transaction, error)
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (repo *repository) GetAll() []Transaction {
	t1 := Transaction{
		Id:       1,
		Codigo:   "afasfa",
		Moneda:   "dolar",
		Monto:    1500,
		Emisor:   "Pepito",
		Receptor: "Juancito",
		Fecha:    "2021-12-23",
		Estado:   true,
	}
	transactions = append(transactions, t1)
	return transactions
}

func (repo *repository) NewUser(codigo string, moneda string, monto int, emisor string, receptor string, fecha string) []Transaction {
	lastId := len(transactions)
	id := lastId + 1
	t := Transaction{id, codigo, moneda, monto, emisor, receptor, fecha, true}
	transactions = append(transactions, t)
	return transactions
}

func (repo *repository) Update(id int, codigo string, moneda string, monto int, emisor string, receptor string, fecha string) (Transaction, error) {
	t := Transaction{
		Codigo:   codigo,
		Moneda:   moneda,
		Monto:    monto,
		Emisor:   emisor,
		Receptor: receptor,
		Fecha:    fecha,
	}
	updated := false
	for i := range transactions {
		if transactions[i].Id == id {
			t.Id = id
			transactions[i] = t
			updated = true
		}
	}
	if !updated {
		return Transaction{}, fmt.Errorf("transacción %d not found", id)
	}
	return t, nil
}

func (repo *repository) Delete(id int) error {
	deleted := false
	for i := range transactions {
		if transactions[i].Id == id {
			deleted = true
		}
	}
	if !deleted {
		return fmt.Errorf("transaccion %d no encontrado", id)
	}
	transactions[id].Estado = false
	return nil
}

func (repo *repository) UpdateCodigoAndMonto(id int, codigo string, monto int) (Transaction, error) {
	var t Transaction
	updated := false
	for i := range transactions {
		if transactions[i].Id == id {
			transactions[i].Codigo = codigo
			transactions[i].Monto = monto
			updated = true
			t = transactions[i]
		}
	}
	if !updated {
		return Transaction{}, fmt.Errorf("transaccion %d no encontrada", id)
	}
	return t, nil
}
