package controller

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/JoaquinRibero/go-web/internal/transactions"
	"github.com/JoaquinRibero/go-web/pkg/web"
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

func Mensaje(verr validator.ValidationErrors) []string {
	//errs := make(map[string]string)
	var errors []string
	for _, f := range verr {
		//errs[f.Field()] = fmt.Sprintf("El campo %s es requerido", f.Field())
		errors = append(errors, fmt.Sprintf("El campo %s es requerido", f.Field()))
	}
	return errors
}

func (t *Transaction) ValidateToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		tokenEnv := os.Getenv("TOKEN")
		if token == "" {
			ctx.AbortWithStatusJSON(401, gin.H{"error": "API token is required"})
			return
		}
		if token != tokenEnv {
			ctx.AbortWithStatusJSON(401, gin.H{"errors": "Unauthorized. Please add a valid API token"})
			return
		}
		ctx.Next()
	}
}

func (c *Transaction) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		t := c.service.GetAll()
		ctx.JSON(200, web.NewResponse(200, t, []string{""}))
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
				ctx.JSON(400, web.NewResponse(400, nil, Mensaje(verr)))
				return
			}
		} else {
			t, err := c.service.NewUser(req.Codigo, req.Moneda, req.Monto, req.Emisor, req.Receptor, req.Fecha)
			if err != nil {
				ctx.JSON(401, web.NewResponse(401, nil, []string{err.Error()}))
			}
			ctx.JSON(200, web.NewResponse(200, t, []string{}))
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
				ctx.JSON(400, web.NewResponse(400, nil, Mensaje(verr)))
				return
			}
		}
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, []string{"invalid ID"}))
			return
		}

		t, err := c.service.Update(int(id), req.Codigo, req.Moneda, req.Monto, req.Emisor, req.Receptor, req.Fecha)
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, []string{err.Error()}))
		}
		ctx.JSON(200, web.NewResponse(200, t, []string{}))
	}
}

func (t *Transaction) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, []string{"invalid ID"}))
			return
		}
		err = t.service.Delete(int(id))
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, []string{err.Error()}))
		}
		ctx.JSON(200, web.NewResponse(200, nil, []string{}))
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
				ctx.JSON(400, web.NewResponse(400, nil, Mensaje(verr)))
				return
			}
		}
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, []string{"invalid ID"}))
			return
		}
		t, err := t.service.UpdateCodigoAndMonto(int(id), req.Codigo, req.Monto)
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, []string{err.Error()}))
		}
		ctx.JSON(200, web.NewResponse(200, t, []string{}))
	}
}
