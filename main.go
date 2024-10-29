package main

import (
	"fmt"
	"pizarria-golang/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/pizzas", getPizzas)
	router.POST("/pizzas", postPizzas)
	router.GET("/pizzas/:id", getPizzasByID)
	router.Run()
}

var pizzas = []models.Pizza{
	{ID: 1, Nome: "Toscana", Preco: 50.2},
	{ID: 2, Nome: "Marguerita", Preco: 90.9},
	{ID: 3, Nome: "Atum com queijo", Preco: 100},
}

func getPizzas(c *gin.Context) {
	c.JSON(200, gin.H{
		"pizzas": pizzas,
	})
}

func postPizzas(c *gin.Context) {
	var newPizza models.Pizza
	if err := c.ShouldBindJSON(&newPizza); err != nil {
		c.JSON(400, gin.H{
			"erro": err.Error()})
		return
	}
	pizzas = append(pizzas, newPizza)
}

func getPizzasByID(c *gin.Context) {
	idParam := c.Param("id")
	fmt.Println(idParam)
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{
			"erro": err.Error()})
		return
	}

	for _, p := range pizzas {
		if p.ID == id {
			c.JSON(200, p)
			return
		}
	}

	c.JSON(404, gin.H{
		"message": "Pizza not found"})
}
