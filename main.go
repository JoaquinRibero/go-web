package main

import (
	"log"

	"github.com/JoaquinRibero/go-web/cmd/server/controller"
	"github.com/JoaquinRibero/go-web/internal/transactions"
	"github.com/JoaquinRibero/go-web/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error al intentar obtener las variables de entorno")
	}
	db := store.New(store.FileType, "transactions.json")
	repo := transactions.NewRepository(db)
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
