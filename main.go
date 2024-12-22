package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/aniketxpawar/go-task-manager/manager"
	"github.com/gin-gonic/gin"
)

func main() {
	dispatcher := manager.NewDispatcher(3)
	dispatcher.Start()

	router := gin.Default()

	router.POST("/submit", func (c *gin.Context){
		var payload struct {
			Message string `json:"message"`
		}

		if err := c.ShouldBindJSON(&payload); err != nil{
			c.JSON(http.StatusBadRequest, gin.H{"error":"Invalid Payload"})
			return
		}

		task := manager.Task{
			ID: int(time.Now().UnixNano()),
			Payload: payload.Message,
			SubmittedAt: time.Now(),
		}
		dispatcher.TaskQueue.Enqueue(task)

		c.JSON(http.StatusCreated, gin.H{"message":fmt.Sprintf("Task %d submitted", task.ID)})
	})

	fmt.Println("Server is running...")
	router.Run(":8080")
}