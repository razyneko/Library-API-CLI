package routes

import (
	"github.com/Library-API-CLI/handlers"
	"github.com/gin-gonic/gin"
)

func SetUpRoutes(router *gin.Engine) {
	router.GET("/books", handlers.GetAllBooks)
	router.GET("/book/:id", handlers.GetBookById)
	router.POST("/book/add", handlers.AddBook)
	router.PATCH("/book/return/:id", handlers.ReturnBook)
	router.PATCH("/book/checkout/:id", handlers.CheckoutBook)
	// PATCH because partial updation just stock updation
}