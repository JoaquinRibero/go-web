package transactions

import (
	"encoding/json"
	"errors"
	"fmt"
	"testing"

	"github.com/JoaquinRibero/go-web/internal/domain"
	"github.com/JoaquinRibero/go-web/pkg/store"
	"github.com/stretchr/testify/assert"
)

func TestServiceGetAll(t *testing.T) {
	input := []domain.Transaction{
		{
			Id:       1,
			Codigo:   "transacción 3",
			Moneda:   "peso",
			Monto:    1500,
			Emisor:   "joaquin",
			Receptor: "bootcamp",
			Fecha:    "2021-12-21T00:00:00",
			Estado:   true,
		}, {
			Id:       2,
			Codigo:   "transacción 2",
			Moneda:   "peso",
			Monto:    13456,
			Emisor:   "joaquin",
			Receptor: "tony stark",
			Fecha:    "2021-12-21T00:00:00",
			Estado:   true,
		},
	}
	dataJson, _ := json.Marshal(input)
	dbMock := store.Mock{
		Data: dataJson,
	}
	storeStub := store.FileStore{
		FileName: "",
		Mock:     &dbMock,
	}
	myRepo := NewRepository(&storeStub)
	myService := NewService(myRepo)

	result, err := myService.GetAll()

	assert.Equal(t, input, result)
	assert.Nil(t, err)
}

func TestServiceGetAllError(t *testing.T) {
	// Initializing Input/output
	expectedError := errors.New("error for GetAll")
	dbMock := store.Mock{
		Err: expectedError,
	}

	storeStub := store.FileStore{
		FileName: "",
		Mock:     &dbMock,
	}
	myRepo := NewRepository(&storeStub)
	myService := NewService(myRepo)

	result, err := myService.GetAll()

	assert.Equal(t, expectedError, err)
	assert.Nil(t, result)
}

func TestNewTransaction(t *testing.T) {
	testTransaction := domain.Transaction{
		Id:       1,
		Codigo:   "transacción 3",
		Moneda:   "peso",
		Monto:    1500,
		Emisor:   "joaquin",
		Receptor: "bootcamp",
		Fecha:    "2021-12-21",
		Estado:   true,
	}
	var input = []domain.Transaction{}
	input = append(input, testTransaction)

	dataJson, _ := json.Marshal(input)
	dbMock := store.Mock{
		Data: dataJson,
	}

	storeStub := store.FileStore{
		FileName: "",
		Mock:     &dbMock,
	}
	myRepo := NewRepository(&storeStub)
	myService := NewService(myRepo)

	result, _ := myService.NewUser(testTransaction.Codigo, testTransaction.Moneda, testTransaction.Monto, testTransaction.Emisor, testTransaction.Receptor, testTransaction.Fecha)
	fmt.Println(result)
	assert.Equal(t, testTransaction.Codigo, result.Codigo)
	assert.Equal(t, testTransaction.Moneda, result.Moneda)
	assert.Equal(t, testTransaction.Monto, result.Monto)
	assert.Equal(t, testTransaction.Emisor, result.Emisor)
	assert.Equal(t, testTransaction.Receptor, result.Receptor)
	assert.Equal(t, testTransaction.Fecha, result.Fecha)
	assert.Equal(t, 2, result.Id)
}

func TestNewTransactionError(t *testing.T) {
	testTransaction := domain.Transaction{
		Codigo:   "transacción 3",
		Moneda:   "peso",
		Monto:    1500,
		Emisor:   "joaquin",
		Receptor: "bootcamp",
		Fecha:    "2021-12-21T00:00:00",
		Estado:   true,
	}
	expectedError := errors.New("error for Storage")
	dbMock := store.Mock{
		Err: expectedError,
	}
	storeStub := store.FileStore{
		FileName: "",
		Mock:     &dbMock,
	}
	myRepo := NewRepository(&storeStub)
	myService := NewService(myRepo)

	result, err := myService.NewUser(testTransaction.Codigo, testTransaction.Moneda, testTransaction.Monto, testTransaction.Emisor, testTransaction.Receptor, testTransaction.Fecha)

	assert.Equal(t, expectedError, err)
	assert.Equal(t, domain.Transaction{}, result)
}
