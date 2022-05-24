package main

import (
	"test/cmd/server/handler"
	"test/internal/products"
	"test/pkg/store"

	"github.com/gin-gonic/gin"
)

func main() {
	db := store.New(store.FileType, "./products.json")
	repo := products.NewRepository(db)
	service := products.NewService(repo)
	handler := handler.NewProduct(service)

	r := gin.Default()
	pr := r.Group("/products")
	pr.POST("/", handler.Store)
	pr.GET("/", handler.GetAll)
	pr.PUT("/:id", handler.Update)
	pr.PATCH("/:id", handler.UpdateName)
	pr.DELETE("/:id", handler.Delete)
	r.Run()

}
