package transactions

type Transaction struct {
	Id       int    `json:"_id"`
	Codigo   string `json:"codigo"`
	Moneda   string `json:"moneda"`
	Monto    int    `json:"monto"`
	Emisor   string `json:"emisor"`
	Receptor string `json:"receptor"`
	Fecha    string `json:"fecha"`
}

var transactions []Transaction

type Repository interface {
	GetAll() []Transaction
	NewUser(codigo string, moneda string, monto int, emisor string, receptor string, fecha string) []Transaction
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
	}
	transactions = append(transactions, t1)
	return transactions
}

func (repo *repository) NewUser(codigo string, moneda string, monto int, emisor string, receptor string, fecha string) []Transaction {
	lastId := len(transactions)
	id := lastId + 1
	t := Transaction{id, codigo, moneda, monto, emisor, receptor, fecha}
	transactions = append(transactions, t)
	return transactions
}
