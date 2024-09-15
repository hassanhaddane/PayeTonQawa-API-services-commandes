package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Order struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

var orders []Order

func main() {
	router := gin.Default()

	// Create a new order
	router.POST("/orders", func(c *gin.Context) {
		var newOrder Order
		if err := c.ShouldBindJSON(&newOrder); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		orders = append(orders, newOrder)
		c.JSON(http.StatusCreated, newOrder)
	})

	// Get all orders
	router.GET("/orders", func(c *gin.Context) {
		c.JSON(http.StatusOK, orders)
	})

	// Get a specific order by ID
	router.GET("/orders/:id", func(c *gin.Context) {
		id := c.Param("id")

		for _, order := range orders {
			if order.ID == id {
				c.JSON(http.StatusOK, order)
				return
			}
		}

		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
	})

	// Update a order by ID
	router.PUT("/orders/:id", func(c *gin.Context) {
		id := c.Param("id")

		var updatedOrder Order
		if err := c.ShouldBindJSON(&updatedOrder); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		for i, user := range orders {
			if user.ID == id {
				orders[i] = updatedOrder
				c.JSON(http.StatusOK, updatedOrder)
				return
			}
		}

		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
	})

	// Delete a order by ID
	router.DELETE("/orders/:id", func(c *gin.Context) {
		id := c.Param("id")

		for i, order := range orders {
			if order.ID == id {
				orders = append(orders[:i], orders[i+1:]...)
				c.JSON(http.StatusOK, gin.H{"message": "Order deleted"})
				return
			}
		}

		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
	})

	router.Run(":8080")
}
