package controller

import (
	"errors"
	"fmt"

	"github.com/JoaquinRibero/go-web/internal/transactions"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Transaction struct {
	service transactions.Service
}

func NewTransaction(t transactions.Service) *Transaction {
	return &Transaction{
		service: t,
	}
}

func Mensaje(verr validator.ValidationErrors) map[string]string {
	errs := make(map[string]string)
	for _, f := range verr {
		errs[f.Field()] = fmt.Sprintf("El campo %s es requerido", f.Field())
	}
	return errs
}

func (t *Transaction) ValidateToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != "joaquin" {
			ctx.JSON(401, gin.H{"errors": "no tiene permisos para realizar la peticion solicitada"})
		}
	}
}

func (c *Transaction) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		t := c.service.GetAll()
		ctx.JSON(200, t)
	}
}

func (c *Transaction) NewUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		type request struct {
			Codigo   string `json:"codigo"`
			Moneda   string `json:"moneda"`
			Monto    int    `json:"monto"`
			Emisor   string `json:"emisor"`
			Receptor string `json:"receptor"`
			Fecha    string `json:"fecha"`
		}
		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			var verr validator.ValidationErrors
			if errors.As(err, &verr) {
				ctx.JSON(400, gin.H{"errors": Mensaje(verr)})
				return
			}
		} else {
			t := c.service.NewUser(req.Codigo, req.Moneda, req.Monto, req.Emisor, req.Receptor, req.Fecha)
			ctx.JSON(200, t)
		}

	}
}
