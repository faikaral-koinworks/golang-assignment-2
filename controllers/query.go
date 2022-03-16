package controllers

import (
	"assignment-2/database"
	"assignment-2/models"
	"encoding/json"
	"fmt"
)

func QueryCreate(orderInput models.Order) models.Order {
	db := database.GetDB()

	newOrder := orderInput

	dberr := db.Debug().Create(&newOrder).Error

	if dberr != nil {
		panic(dberr)

	}

	return newOrder
}

func QueryGetAll() []models.Order {
	db := database.GetDB()

	var orders []models.Order

	dberr := db.Preload("Items").Find(&orders).Error

	if dberr != nil {
		panic(dberr)
	}

	return orders
}

func QueryDeleteByID(id uint) {
	db := database.GetDB()

	dberr := db.Where("Order_id=?", id).Delete(&models.Item{}).Error

	if dberr != nil {
		panic(dberr)
	}

	dberr = db.Delete(&models.Order{}, id).Error

	if dberr != nil {
		panic(dberr)
	}

	fmt.Println("Data Deleted")
}

func testUpdate(id uint) {
	db := database.GetDB()
	mockOrder := `
	{
		"orderedAt":"2022-11-09T21:21:46+00:00",
		"customerName":"Test Updated",
		"items":[
			{
				"lineItemID":1,
				"itemCode":"112312323",
				"description":"Updatedtest2",
				"quantity":1
			},
			{
				"lineItemID":2,
				"itemCode":"121231233",
				"description":"Updatedtest2",
				"quantity":1
			}
		]
	}
	`

	var updatedOrder models.Order

	err := json.Unmarshal([]byte(mockOrder), &updatedOrder)
	if err != nil {
		panic(err)
	}

	for i := range updatedOrder.Items {
		err = db.Model(&updatedOrder.Items[i]).Where("Item_id=?", updatedOrder.Items[i].Item_id).Updates(updatedOrder.Items[i]).Error
		if err != nil {
			panic(err)
		}
	}

	var updatedOnlyOrder models.Order
	updatedOnlyOrder.Customer_name = updatedOrder.Customer_name
	updatedOnlyOrder.Ordered_at = updatedOrder.Ordered_at

	dberr := db.Model(&updatedOnlyOrder).Where("Order_id=?", id).Updates(updatedOnlyOrder).Error

	if dberr != nil {
		panic(err)
	}

	fmt.Println(updatedOrder)
}
