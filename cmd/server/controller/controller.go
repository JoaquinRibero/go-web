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

//Validacion de errores en el request usando validator
func Mensaje(verr validator.ValidationErrors) []string {

	var errors []string
	for _, f := range verr {
		err := f.ActualTag()
		// Controlar los params
		if f.Param() != "" {
			err = fmt.Sprintf("%s=%s", err, f.Param())
		}
		// Dependiendo de cual sea el posible error, voy agregando mensajes personalizados.
		// Si no lo conozco, simplemente devuelvo "nombre del campo = motivo de la falla" (default)
		switch err {
		case "required":
			errors = append(errors, fmt.Sprintf("El campo %s es requerido", f.Field()))
		default:
			errors = append(errors, fmt.Sprintf("%s=%s", f.Field(), err))
		}
	}
	return errors
}

func (t *Transaction) ValidateToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		tokenEnv := os.Getenv("TOKEN")
		if token == "" {
			ctx.AbortWithStatusJSON(401, web.NewResponse(401, nil, "API token is required"))
			return
		}
		if token != tokenEnv {
			ctx.AbortWithStatusJSON(401, web.NewResponse(401, nil, "Unauthorized. Please add a valid API token"))
			return
		}
		ctx.Next()
	}
}

// ListTransactions godoc
// @Summary List Transactions
// @Tags Transactions
// @Description get all transactions
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @Router / [get]
func (c *Transaction) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		t := c.service.GetAll()
		ctx.JSON(200, web.NewResponse(200, t, nil))
	}
}

// @Summary Add Transaction
// @Tags Transactions
// @Description add new transactions
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param codigo body string true "codigo"
// @Param moneda body string true "moneda"
// @Param monto body int true "monto"
// @Param emisor body string true "emisor"
// @Param receptor body string true "receptor"
// @Param fecha body string true "fecha"
// @Success 200 {object} web.Response
// @Router /new [post]
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
				ctx.JSON(401, web.NewResponse(401, nil, err.Error()))
			}
			ctx.JSON(200, web.NewResponse(200, t, nil))
		}

	}
}

// @Summary Update Transaction
// @Tags Transactions
// @Description update full transaction
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param id path int true "id"
// @Param codigo body string true "codigo"
// @Param moneda body string true "moneda"
// @Param monto body int true "monto"
// @Param emisor body string true "emisor"
// @Param receptor body string true "receptor"
// @Param fecha body string true "fecha"
// @Success 200 {object} web.Response
// @Router /{id} [put]
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
			ctx.JSON(400, web.NewResponse(400, nil, "invalid ID"))
			return
		}

		t, err := c.service.Update(int(id), req.Codigo, req.Moneda, req.Monto, req.Emisor, req.Receptor, req.Fecha)
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
		}
		ctx.JSON(200, web.NewResponse(200, t, nil))
	}
}

func (t *Transaction) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, "invalid ID"))
			return
		}
		err = t.service.Delete(int(id))
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
		}
		ctx.JSON(200, web.NewResponse(200, nil, nil))
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
			ctx.JSON(400, web.NewResponse(400, nil, "invalid ID"))
			return
		}
		t, err := t.service.UpdateCodigoAndMonto(int(id), req.Codigo, req.Monto)
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
		}
		ctx.JSON(200, web.NewResponse(200, t, nil))
	}
}
