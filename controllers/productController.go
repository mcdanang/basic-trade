package controllers

import (
	"basic-trade/database"
	models "basic-trade/models/entity"
	requests "basic-trade/models/request"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateProduct(ctx *gin.Context) {
	db := database.GetDB()
	if db == nil {
		fmt.Println("Error: Database connection is nil.")
		return
	}

	var productReq requests.ProductRequest

	if err := ctx.ShouldBind(&productReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":    productReq,
		"message": "succeed create product",
	})
}

func GetAllProduct(ctx *gin.Context) {
	db := database.GetDB()

	products := []models.Product{}
	// err := db.Preload("Items").Find(&products).Error
	err := db.Find(&products).Error

	if err != nil {
		fmt.Println("Error getting product data:", err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":    products,
		"message": "succeed get all product",
	})
}

// func UpdateOrder(ctx *gin.Context) {
// 	db := database.GetDB()
// 	orderID := ctx.Param("orderID")
// 	// condition := false

// 	updatedOrder := models.Order{}
// 	updatedItems := models.Item{}

// 	if err := ctx.ShouldBindJSON(&updatedOrder); err != nil {
// 		ctx.AbortWithError(http.StatusBadRequest, err)
// 		return
// 	}

// 	orderIDUint, e := strconv.ParseUint(orderID, 10, 64)
// 	if e != nil {
// 		fmt.Println("Error parsing orderID from URL params", e)
// 		return
// 	}

// 	err := db.Model(&updatedOrder).Where("id = ?", orderID).Updates(models.Order{
// 		ID:           orderIDUint,
// 		CustomerName: updatedOrder.CustomerName,
// 		OrderedAt:    updatedOrder.OrderedAt,
// 	}).Error

// 	if err != nil {
// 		fmt.Println("Error updating order data:", err)
// 		return
// 	}

// 	for _, item := range updatedOrder.Items {
// 		err := db.Model(&updatedItems).Where("order_id = ?", orderID).Updates(models.Item{
// 			Name:        item.Name,
// 			Description: item.Description,
// 			Quantity:    item.Quantity,
// 		}).Error

// 		if err != nil {
// 			fmt.Println("Error updating item data:", err)
// 			return
// 		}
// 	}

// 	fmt.Printf("Updated order 2: %+v \n", updatedOrder)

// 	// bookID := ctx.Param("bookID")
// 	// condition := false
// 	// var updatedBook Book

// 	// if err := ctx.ShouldBindJSON(&updatedBook); err != nil {
// 	// 	ctx.AbortWithError(http.StatusBadRequest, err)
// 	// 	return
// 	// }

// 	// for i, book := range BookDatas {
// 	// 	if bookID == book.BookID {
// 	// 		condition = true
// 	// 		BookDatas[i] = updatedBook
// 	// 		BookDatas[i].BookID = bookID
// 	// 		break
// 	// 	}
// 	// }

// 	// if !condition {
// 	// 	ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
// 	// 		"data":    nil,
// 	// 		"message": fmt.Sprintf("book with id %v not found", bookID),
// 	// 	})
// 	// 	return
// 	// }

// 	// convertBookId, _ := strconv.Atoi(bookID)
// 	ctx.JSON(http.StatusOK, gin.H{
// 		"message": fmt.Sprintf("Order with id %v has been successfully updated", orderID),
// 	})
// }

// func DeleteOrder(ctx *gin.Context) {
// 	db := database.GetDB()
// 	orderID := ctx.Param("orderID")

// 	order := models.Order{}
// 	item := models.Item{}

// 	var err error

// 	db.Transaction(func(tx *gorm.DB) error {

// 		if err = tx.Where("order_id = ?", orderID).Delete(&item).Error; err != nil {
// 			return err
// 		}

// 		if err = tx.Where("id = ?", orderID).Delete(&order).Error; err != nil {
// 			return err
// 		}

// 		return nil
// 	})

// 	if err != nil {
// 		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
// 			"data":    nil,
// 			"message": fmt.Sprintf("Error deleting order: %v", err.Error()),
// 		})
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, gin.H{
// 		"data":    nil,
// 		"message": fmt.Sprintf("Order with id %v has been successfully deleted", orderID),
// 	})

// }
