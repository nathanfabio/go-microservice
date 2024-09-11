package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nathanfabio/go-microservice/cmd/api/controllers"
	"github.com/nathanfabio/go-microservice/internal/repositories"
)

func CategoryRoutes(router *gin.Engine) {
	routes := router.Group("/categories")

	inMemoryCategoryRepository := repositories.NewInMemoryCategoryRepository()
	
	routes.POST("/", func(ctx *gin.Context) {
		controllers.CreateCategory(ctx, inMemoryCategoryRepository)

	})

	routes.GET("/", func(ctx *gin.Context) {
		controllers.ListCategory(ctx, inMemoryCategoryRepository)
		
	})
}