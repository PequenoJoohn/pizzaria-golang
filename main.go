package main

import (
	"pizarria-golang/models"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/pizzas", getPizzas)
	router.Run()
}

func getPizzas(c *gin.Context) {
	var pizzas = []models.Pizza{
		{ID: 1, Nome: "Toscana", Preco: 50.2},
		{ID: 2, Nome: "Marguerita", Preco: 90.9},
		{ID: 3, Nome: "Atum com queijo", Preco: 100},
	}
	c.JSON(200, gin.H{
		"pizzas": pizzas,
	})
}
