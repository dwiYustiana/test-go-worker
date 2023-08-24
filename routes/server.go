package routes

import (
	"log"
	"net/http"
	"test-go-worker/config"
	"test-go-worker/handler"
	"test-go-worker/repository"

	"test-go-worker/db"

	"github.com/gin-gonic/gin"
)

var configJson = config.NewViperConfig()

func Init() {
	r := SetRouter()

	err := r.Run(configJson.GetString("server.port"))
	if err != nil {
		log.Fatal(err)
	}
}

func SetRouter() *gin.Engine {
	//connection database
	db := db.ConnectionMysql()

	//Repository
	newTransactionRepo := repository.NewTransactionRepository(db)

	router := gin.New()
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "API WORKING")
	})

	transaction := router.Group("transaction")
	{
		transactionController := &handler.TransactionHandler{
			TransactionModel: newTransactionRepo,
		}
		TransactionRoutes(transaction, transactionController)
	}

	return router

}
