package controllers

import (
	"errors"
	"github.com/ernestomr87/go-resourceapi/database"
	"github.com/ernestomr87/go-resourceapi/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

//	Dependency Injection
type TodoRepo struct {
	Db *gorm.DB
}

func New() *TodoRepo {
	db := database.InitDB()
	db.AutoMigrate(&models.Todo{})

	return &TodoRepo{Db: db}
}

func (repository *TodoRepo) CreateTodo(c *gin.Context) {
	var todo models.Todo
	if c.BindJSON(&todo) == nil {
		err := models.CreateTodo(repository.Db, &todo)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
		c.JSON(http.StatusOK, todo)
	} else {
		c.JSON(http.StatusBadRequest, todo)
	}
}

func (repository *TodoRepo) GetTodos(c *gin.Context) {
	var todos []models.Todo
	err := models.GetTodos(repository.Db, &todos)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, todos)
}

//	http://<server>/todo/2
func (repository *TodoRepo) GetTodo(c *gin.Context) {
	id, _ := c.Params.Get("id")
	idn, _ := strconv.Atoi(id)
	var todo models.Todo

	err := models.GetTodoById(repository.Db, &todo, idn)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, todo)
}

//	http://<server>/todo/2
func (repository *TodoRepo) UpdateTodo(c *gin.Context) {
	var todo models.Todo
	var updatedTodo models.Todo

	id, _ := c.Params.Get("id")
	idn, _ := strconv.Atoi(id)
	err := models.GetTodoById(repository.Db, &updatedTodo, idn)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}
	if c.BindJSON(&todo) == nil {
		updatedTodo.Task = todo.Task
		updatedTodo.Completed = todo.Completed
		updatedTodo.StartDate = todo.StartDate
		updatedTodo.EndDate = todo.EndDate

		err = models.UpdateTodo(repository.Db, &updatedTodo)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
			return
		}
		c.JSON(http.StatusOK, updatedTodo)
	} else {
		c.JSON(http.StatusBadRequest, todo)
	}
}

//	http://<server>/todo/2
func (repository *TodoRepo) DeleteTodoById(c *gin.Context) {
	var todo models.Todo
	id, _ := c.Params.Get("id")
	idn, _ := strconv.Atoi(id)

	err := models.DeleteTodoByID(repository.Db, &todo, idn)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Todo was deleted succesfully!!!",
	})
}
