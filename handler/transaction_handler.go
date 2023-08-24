package handler

import (
	"fmt"
	"net/http"
	"sync"
	"test-go-worker/config"
	"test-go-worker/db/entity"
	"test-go-worker/db/request"
	"test-go-worker/repository"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TransactionHandler struct {
	TransactionModel repository.TransactionInterface
}

func (p TransactionHandler) Worker(id int, taskQueue <-chan entity.Transaction, wg *sync.WaitGroup) error {
	defer wg.Done()

	for task := range taskQueue {
		fmt.Printf("Worker %d: Processing task %s\n", id, task.Customer)
		time.Sleep(10 * time.Millisecond) // Simulate task processing
		p.TransactionModel.InsertTransaction(task)
		fmt.Printf("Worker %d: Task %s completed\n", id, task.Customer)
	}
	return nil
}

func (p TransactionHandler) InsertTransaction(c *gin.Context) {

	var (
		err  error
		rows request.TransactionRequest
		wg   sync.WaitGroup
	)

	if err = c.ShouldBindJSON(&rows); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}

	config := config.NewViperConfig()
	numWorkers := config.GetInt("worker")
	numTasks := len(rows.Transaction)
	taskQueue := make(chan entity.Transaction, numTasks)

	// Create worker goroutines
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go p.Worker(i, taskQueue, &wg)
	}

	// Enqueue tasks
	for _, result := range rows.Transaction {
		wg.Add(1)
		trimString := result.Timestamp
		if len(result.Timestamp) >= 20 {
			trimString = result.Timestamp[0:19]
		}
		timestamp, err := time.Parse("2006-01-02 15:04:05", trimString) // Customize the layout
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		id := uuid.New()
		transaction := entity.Transaction{
			ID:            id,
			TransactionId: result.ID,
			RequestId:     rows.RequestId,
			Customer:      result.Customer,
			Price:         result.Price,
			Quantity:      result.Quantity,
			Timestamp:     timestamp,
		}
		taskQueue <- transaction

	}

	go func() {
		wg.Wait()
		close(taskQueue)
	}()

	c.JSON(http.StatusOK, gin.H{
		"msg": "transaction complate",
	})
}
