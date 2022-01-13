package transactions

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/JoaquinRibero/go-web/internal/domain"
	"github.com/JoaquinRibero/go-web/pkg/store"
	"github.com/stretchr/testify/assert"
)

const errorGetAll = "error for GetAll"

func TestGetAll(t *testing.T) {
	var ts []domain.Transaction
	t1 := domain.Transaction{
		Id:       1,
		Codigo:   "transacción 3",
		Moneda:   "peso",
		Monto:    1500,
		Emisor:   "joaquin",
		Receptor: "bootcamp",
		Fecha:    "2021-12-21T00:00:00",
		Estado:   true,
	}
	t2 := domain.Transaction{
		Id:       2,
		Codigo:   "transacción 2",
		Moneda:   "peso",
		Monto:    13456,
		Emisor:   "joaquin",
		Receptor: "tony stark",
		Fecha:    "2021-12-21T00:00:00",
		Estado:   true,
	}
	ts = append(ts, t1, t2)
	dataJson, _ := json.Marshal(ts)
	dbStub := store.Mock{
		Data: dataJson,
		Err:  nil,
	}
	storeMocked := store.FileStore{
		FileName: "",
		Mock:     &dbStub,
	}
	myRepo := NewRepository(&storeMocked)
	resp, _ := myRepo.GetAll()
	assert.Equal(t, resp, ts)
}

func TestGetAllError(t *testing.T) {
	// Initializing Input/output
	expectedError := errors.New(errorGetAll)
	dbMock := store.Mock{
		Data: nil,
		Err:  expectedError,
	}
	storeMocked := store.FileStore{
		FileName: "",
		Mock:     &dbMock,
	}
	myRepo := NewRepository(&storeMocked)

	_, err := myRepo.GetAll()

	assert.Equal(t, err, expectedError)
}
