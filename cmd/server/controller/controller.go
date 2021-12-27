package controller

import (
	"errors"
	"fmt"
	"os"
	"strconv"

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
		tokenEnv := os.Getenv("TOKEN")
		if token != tokenEnv {
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
			Codigo   string `json:"codigo" binding:"required"`
			Moneda   string `json:"moneda" binding:"required"`
			Monto    int    `json:"monto" binding:"required"`
			Emisor   string `json:"emisor" binding:"required"`
			Receptor string `json:"receptor" binding:"required"`
			Fecha    string `json:"fecha" binding:"required"`
		}
		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			var verr validator.ValidationErrors
			if errors.As(err, &verr) {
				ctx.JSON(400, gin.H{"errors": Mensaje(verr)})
				return
			}
		} else {
			t, err := c.service.NewUser(req.Codigo, req.Moneda, req.Monto, req.Emisor, req.Receptor, req.Fecha)
			if err != nil {
				ctx.JSON(401, gin.H{"error": err.Error()})
			}
			ctx.JSON(200, t)
		}

	}
}

func (c *Transaction) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		type request struct {
			Codigo   string `json:"codigo" binding:"required"`
			Moneda   string `json:"moneda" binding:"required"`
			Monto    int    `json:"monto" binding:"required"`
			Emisor   string `json:"emisor" binding:"required"`
			Receptor string `json:"receptor" binding:"required"`
			Fecha    string `json:"fecha" binding:"required"`
		}
		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			var verr validator.ValidationErrors
			if errors.As(err, &verr) {
				ctx.JSON(400, gin.H{"errors": Mensaje(verr)})
				return
			}
		}
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "invalid ID"})
			return
		}

		t, err := c.service.Update(int(id), req.Codigo, req.Moneda, req.Monto, req.Emisor, req.Receptor, req.Fecha)
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
		}
		ctx.JSON(200, t)
	}
}

func (t *Transaction) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "invalid ID"})
			return
		}
		err = t.service.Delete(int(id))
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
		}
		ctx.JSON(200, gin.H{"message": "success"})
	}
}

func (t *Transaction) UpdatePartial() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		type request struct {
			Codigo string `json:"codigo" binding:"required"`
			Monto  int    `json:"monto" binding:"required"`
		}
		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			var verr validator.ValidationErrors
			if errors.As(err, &verr) {
				ctx.JSON(400, gin.H{"errors": Mensaje(verr)})
				return
			}
		}
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "invalid ID"})
			return
		}
		t, err := t.service.UpdateCodigoAndMonto(int(id), req.Codigo, req.Monto)
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
		}
		ctx.JSON(200, t)
	}
}
