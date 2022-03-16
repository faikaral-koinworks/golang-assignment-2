package main

import (
	"assignment-2/database"
	"assignment-2/models"
	"encoding/json"
	"fmt"
)

func main() {
	database.StartDB()
	testCreate()
}

func testCreate() {
	db := database.GetDB()

	mockOrder := `
	{
		"orderedAt":"2019-11-09T21:21:46+00:00",
		"customerName":"Test",
		"items":[
			{
				"itemCode":"123",
				"description":"test",
				"quantity":1
			},
			{
				"itemCode":"123",
				"description":"test",
				"quantity":1
			}
		]
	}
	`
	var newOrder models.Order

	err := json.Unmarshal([]byte(mockOrder), &newOrder)
	if err != nil {
		panic(err)
	}

	dberr := db.Debug().Create(&newOrder).Error

	if dberr != nil {
		fmt.Println("Error creating user data: ", dberr)
		return
	}

	fmt.Println("New User Data:", newOrder)
}
