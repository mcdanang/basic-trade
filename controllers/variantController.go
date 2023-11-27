package controllers

import (
	"basic-trade/database"
	models "basic-trade/models/entity"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateVariant(ctx *gin.Context) {
	db := database.GetDB()
	if db == nil {
		fmt.Println("Error: Database connection is nil.")
		return
	}

	var Variant models.Variant

	if err := ctx.ShouldBind(&Variant); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	Variant.UUID = uuid.New()

	err := db.Debug().Create(&Variant).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":    Variant,
		"message": "succeed create variant",
	})
}

func GetAllVariant(ctx *gin.Context) {
	db := database.GetDB()

	variants := []models.Variant{}
	// err := db.Preload("Items").Find(&products).Error
	err := db.Find(&variants).Error

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": fmt.Sprintf("Error getting all variant: %v", err.Error()),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":    variants,
		"message": "succeed get all variant",
	})
}

// func GetProductByUUID(ctx *gin.Context) {
// 	db := database.GetDB()

// 	Product := models.Product{}
// 	productUUID := ctx.Param("productUUID")
// 	// err := db.Preload("Items").Find(&products).Error
// 	err := db.Where("uuid = ?", productUUID).First(&Product).Error

// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{
// 			"error":   "Bad request",
// 			"message": fmt.Sprintf("Error getting product: %v", err.Error()),
// 		})
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, gin.H{
// 		"data":    Product,
// 		"message": "succeed get product",
// 	})
// }

// func UpdateProduct(ctx *gin.Context) {
// 	db := database.GetDB()
// 	productUUID := ctx.Param("productUUID")
// 	// condition := false

// 	var productReq requests.ProductRequest
// 	// updatedItems := models.Item{}

// 	if err := ctx.ShouldBind(&productReq); err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	Product := models.Product{
// 		Name: productReq.Name,
// 		// ImageURL: uploadResult,
// 		// AdminID:  adminID,
// 	}

// 	err := db.Model(&Product).Where("uuid = ?", productUUID).Updates(models.Product{
// 		Name: productReq.Name,
// 	}).Error

// 	if err != nil {
// 		fmt.Println("Error updating product data:", err)
// 		return
// 	}

// 	// for _, item := range updatedOrder.Items {
// 	// 	err := db.Model(&updatedItems).Where("order_id = ?", orderID).Updates(models.Item{
// 	// 		Name:        item.Name,
// 	// 		Description: item.Description,
// 	// 		Quantity:    item.Quantity,
// 	// 	}).Error

// 	// 	if err != nil {
// 	// 		fmt.Println("Error updating item data:", err)
// 	// 		return
// 	// 	}
// 	// }

// 	// fmt.Printf("Updated order 2: %+v \n", updatedOrder)

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
// 		"message": fmt.Sprintf("Product with uuid %v has been successfully updated", productUUID),
// 	})
// }

// func DeleteProduct(ctx *gin.Context) {
// 	db := database.GetDB()
// 	productUUID := ctx.Param("productUUID")

// 	// item := models.Item{}
// 	Product := models.Product{}

// 	var err error

// 	db.Transaction(func(tx *gorm.DB) error {

// 		// if err = tx.Where("order_id = ?", orderID).Delete(&item).Error; err != nil {
// 		// 	return err
// 		// }

// 		if err = tx.Where("uuid = ?", productUUID).Delete(&Product).Error; err != nil {
// 			return err
// 		}

// 		return nil
// 	})

// 	if err != nil {
// 		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
// 			"data":    nil,
// 			"message": fmt.Sprintf("Error deleting product: %v", err.Error()),
// 		})
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, gin.H{
// 		"data":    nil,
// 		"message": fmt.Sprintf("Product with uuid %v has been successfully deleted", productUUID),
// 	})

// }
