package controllers

import (
	"assignment-2/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateOrders(c *gin.Context) {
	var newOrder models.Order

	if err := c.ShouldBindJSON(&newOrder); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	newOrder = QueryCreate(newOrder)

	c.JSON(http.StatusCreated, gin.H{
		"data":    newOrder,
		"message": "Data sucessfully created",
		"status":  http.StatusCreated,
	})
}

func GetAllOrders(c *gin.Context) {
	orders := QueryGetAll()

	c.JSON(http.StatusOK, gin.H{
		"data":    orders,
		"message": "Orders fetched sucessfully",
		"status":  fmt.Sprintf("%d", http.StatusOK),
	})

}
