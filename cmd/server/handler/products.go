package handler

import (
	"fmt"
	"strconv"
	"test/internal/products"

	"github.com/gin-gonic/gin"
)

type request struct {
	Name  string  `json:"nombre"`
	Type  string  `json:"tipo"`
	Count int     `json:"cantidad"`
	Price float64 `json:"precio"`
}

type Product struct {
	service products.Service
}

func NewProduct(p products.Service) *Product {
	return &Product{
		service: p,
	}
}

func (c *Product) GetAll(ctx *gin.Context) {
	token := ctx.Request.Header.Get("token")
	if token != "123456" {
		ctx.JSON(401, gin.H{
			"error": "token inválido",
		})
		return
	}

	p, err := c.service.GetAll()
	if err != nil {
		ctx.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(200, p)
}

func (c *Product) Store(ctx *gin.Context) {
	token := ctx.Request.Header.Get("token")
	if token != "123456" {
		ctx.JSON(401, gin.H{"error": "token inválido"})
		return
	}
	var req request
	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}
	p, err := c.service.Create(req.Name, req.Type, req.Count, req.Price)
	if err != nil {
		ctx.JSON(404, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, p)
}

func (c *Product) Update(ctx *gin.Context) {
	token := ctx.Request.Header.Get("token")
	if token != "123456" {
		ctx.JSON(401, gin.H{"error": "token inválido"})
		return
	}
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "invalid ID"})
		return
	}
	var req request
	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(404, gin.H{"error": err.Error()})
		return
	}
	if req.Name == "" {
		ctx.JSON(400, gin.H{"error": "El nombre del producto es requerido"})
		return
	}
	if req.Type == "" {
		ctx.JSON(400, gin.H{"error": "El tipo del producto es requerido"})
		return
	}
	if req.Count == 0 {
		ctx.JSON(400, gin.H{"error": "La cantidad es requerida"})
		return
	}
	if req.Price == 0 {
		ctx.JSON(400, gin.H{"error": "El precio es requerido"})
		return
	}
	p, err := c.service.Update(int(id), req.Name, req.Type, req.Count, req.Price)
	if err != nil {
		ctx.JSON(404, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, p)
}

func (c *Product) UpdateName(ctx *gin.Context) {
	token := ctx.Request.Header.Get("token")
	if token != "123456" {
		ctx.JSON(401, gin.H{"error": "token inválido"})
		return
	}
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "invalid ID"})
		return
	}
	var req request
	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(404, gin.H{"error": err.Error()})
		return
	}
	if req.Name == "" {
		ctx.JSON(400, gin.H{"error": "El nombre del producto es requerido"})
		return
	}
	p, err := c.service.UpdateName(int(id), req.Name)
	if err != nil {
		ctx.JSON(404, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, p)
}

func (c *Product) Delete(ctx *gin.Context) {
	token := ctx.Request.Header.Get("token")
	if token != "123456" {
		ctx.JSON(401, gin.H{"error": "token inválido"})
		return
	}
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "invalid ID"})
		return
	}
	err = c.service.Delete(int(id))
	if err != nil {
		ctx.JSON(404, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"data": fmt.Sprintf("El producto %d ha sido eliminado", id)})
}
