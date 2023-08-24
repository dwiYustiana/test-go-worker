package routes

import (
	"test-go-worker/handler"

	"github.com/gin-gonic/gin"
)

func TransactionRoutes(route *gin.RouterGroup, TransactionHandler *handler.TransactionHandler) {

	route.POST("/inserts", TransactionHandler.InsertTransaction)
}
