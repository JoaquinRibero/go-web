package main

import (
	"log"
	"os"

	"github.com/JoaquinRibero/go-web/cmd/server/controller"
	"github.com/JoaquinRibero/go-web/docs"
	"github.com/JoaquinRibero/go-web/internal/transactions"
	"github.com/JoaquinRibero/go-web/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// @title MELI Bootcamp API
// @version 1.0
// @description Transactions API ML Bootcamp GO
// @contact.name API Support
// @contact.email joaquin.ribero@mercadolibre.com
// @host localhost:8080
// @BasePath /transactions/
// @securityDefinitions.basic BasicAuth
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
	docs.SwaggerInfo.Host = os.Getenv("HOST")
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	tr := r.Group("/transactions", t.ValidateToken())
	tr.GET("/", t.GetAll())
	tr.POST("/new", t.NewUser())
	tr.PUT("/:id", t.Update())
	tr.PATCH("/:id", t.UpdatePartial())
	r.Run()
}
