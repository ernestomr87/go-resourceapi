package main

import (
	"fmt"
	"github.com/ernestomr87/go-resourceapi/controllers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln("Error loading .env file")
	}
	r := gin.Default()
	todoRepo := controllers.New()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome todo app!!!",
		})
	})

	r1 := r.Group("/api")
	{
		r1.POST("/todo", todoRepo.CreateTodo)
		r1.GET("/todo", todoRepo.GetTodos)
		r1.GET("/todo/:id", todoRepo.GetTodo)
		r1.PUT("/todo/:id", todoRepo.UpdateTodo)
		r1.DELETE("/todo/:id", todoRepo.DeleteTodoById)
	}

	r.Run("localhost:8080")
	fmt.Println("Server is running!!!")
}
