package main

import (
	"github.com/JoaquinRibero/go-web/cmd/server/controller"
	"github.com/JoaquinRibero/go-web/internal/transactions"
	"github.com/gin-gonic/gin"
)

func main() {
	repo := transactions.NewRepository()
	service := transactions.NewService(repo)
	t := controller.NewTransaction(service)

	r := gin.Default()
	tr := r.Group("/transactions", t.ValidateToken())
	tr.GET("/", t.GetAll())
	tr.POST("/new", t.NewUser())
	tr.PUT("/:id", t.Update())
	tr.PATCH("/:id", t.UpdatePartial())
	r.Run()
}
