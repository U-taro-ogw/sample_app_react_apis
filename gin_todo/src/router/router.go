package router

import (
	"github.com/U-taro-ogw/gin_todo/src/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func Router(dbConm *gorm.DB) {
	todoHandler := controllers.TodoHandler{
		Db: dbConm,
	}

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH"},
		AllowHeaders:     []string{"Content-Type"},
		AllowCredentials: false,
	}))

	r.GET("/todos", todoHandler.GetAll)
	r.GET("/todos/:id", todoHandler.GetTodo)
	r.POST("/todos", todoHandler.CreateTodo)
	r.DELETE("/todos/:id", todoHandler.DeleteTodo)
	r.PUT("todos/:id", todoHandler.UpdateTodo)

	r.Run(":8083")
}
