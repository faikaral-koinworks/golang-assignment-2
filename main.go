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
	testCreate()
	testDelete(1)
	testGet()
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
	var converted []byte

	converted, _ = json.Marshal(newOrder)

	fmt.Println("New User Data:", string(converted))
}

func testGet() {
	db := database.GetDB()

	var orders []models.Order

	dberr := db.Preload("Items").Find(&orders).Error

	if dberr != nil {
		fmt.Println("Error fetching user data: ", dberr)
		return
	}

	converted, _ := json.Marshal(orders)

	fmt.Println(string(converted))
}

func testDelete(id uint) {
	db := database.GetDB()

	dberr := db.Where("Order_id=?", id).Delete(&models.Item{}).Error

	if dberr != nil {
		fmt.Println("Error fetching user data: ", dberr)
		return
	}

	dberr = db.Delete(&models.Order{}, id).Error

	if dberr != nil {
		fmt.Println("Error fetching user data: ", dberr)
		return
	}

	fmt.Println("Data Deleted")
}
