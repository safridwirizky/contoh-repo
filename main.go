package main

import (
	"github.com/gin-gonic/gin"
)

type Item struct {
	ItemID      string `json:"itemId"`
	ItemCode    string `json:"itemCode"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
	OrderID     string `json:"orderId"`
}

type Order struct {
	OrderID      string `json:"orderId"`
	CustomerName string `json:"customerName"`
	OrderedAt    string `json:"orderedAt"`
	Items        []Item `json:"items"`
}

var orders []Order

func main() {
	r := gin.Default()

	r.POST("/orders", createOrder)
	r.GET("/orders", getOrders)
	r.PUT("/orders/:orderId", updateOrder)
	r.DELETE("/orders/:orderId", deleteOrder)

	r.Run()
}

func createOrder(c *gin.Context) {
	var newOrder Order
	if err := c.ShouldBindJSON(&newOrder); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	orders = append(orders, newOrder)
	c.JSON(200, newOrder)
}

func getOrders(c *gin.Context) {
	c.JSON(200, orders)
}

func updateOrder(c *gin.Context) {
	var updatedOrder Order
	if err := c.ShouldBindJSON(&updatedOrder); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	orderId := c.Param("orderId")
	for i, o := range orders {
		if o.OrderID == orderId {
			orders[i] = updatedOrder
			c.JSON(200, updatedOrder)
			return
		}
	}

	c.JSON(404, gin.H{"error": "Order not found"})
}

func deleteOrder(c *gin.Context) {
	orderId := c.Param("orderId")
	for i, o := range orders {
		if o.OrderID == orderId {
			orders = append(orders[:i], orders[i+1:]...)
			c.JSON(200, gin.H{"message": "Order deleted"})
			return
		}
	}

	c.JSON(404, gin.H{"error": "Order not found"})
}
