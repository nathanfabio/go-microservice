package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nathanfabio/go-microservice/cmd/api/controllers"
)

func CategoryRoutes(router *gin.Engine) {
	routes := router.Group("/categories")

	routes.POST("/", controllers.CreateCategory)
}