package controllers

import (
	"github.com/U-taro-ogw/gin_todo/src/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

type TodoHandler struct {
	Db *gorm.DB
}

func (h *TodoHandler) GetAll(c *gin.Context) {
	var todos []models.Todo
	h.Db.Find(&todos)
	c.JSON(http.StatusOK, gin.H{
		"todos": todos,
	})
}

func (h *TodoHandler) GetTodo(c *gin.Context) {
	todo := models.Todo{}
	id := c.Param("id")
	h.Db.First(&todo, id)
	c.JSON(http.StatusOK, gin.H{
		"todo": todo,
	})
}

func (h *TodoHandler) CreateTodo(c *gin.Context) {
	var newTodo models.Todo
	c.BindJSON(&newTodo)
	h.Db.Create(&newTodo)
	c.JSON(http.StatusCreated, gin.H{
		"todo": newTodo,
	})
}

func (h *TodoHandler) DeleteTodo(c *gin.Context) {
	todo := models.Todo{}
	id := c.Param("id")
	h.Db.First(&todo, id)
	h.Db.Delete(&todo)
	c.JSON(http.StatusNoContent, "No Content")
}

func (h *TodoHandler) UpdateTodo(c *gin.Context) {
	todo := models.Todo{}
	id := c.Param("id")
	h.Db.First(&todo, id)
	c.BindJSON(&todo)
	h.Db.Save(&todo)
	c.JSON(http.StatusOK, gin.H{
		"todo": todo,
	})
}
