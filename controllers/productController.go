package controllers

import (
	"basic-trade/database"
	"basic-trade/helpers"
	models "basic-trade/models/entity"
	requests "basic-trade/models/request"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
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

	// Extract the filename without extension
	fileName := helpers.RemoveExtension(productReq.Image.Filename)

	uploadResult, err := helpers.UploadFile(productReq.Image, fileName)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	Product := models.Product{
		UUID:     uuid.NewString(),
		Name:     productReq.Name,
		ImageURL: uploadResult,
		// AdminID:  adminID,
	}

	err = db.Debug().Create(&Product).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
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
	err := db.Preload("Variants").Find(&products).Error

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": fmt.Sprintf("Error getting all product: %v", err.Error()),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":    products,
		"message": "succeed get all product",
	})
}

func GetProductByUUID(ctx *gin.Context) {
	db := database.GetDB()

	Product := models.Product{}
	productUUID := ctx.Param("productUUID")
	// err := db.Preload("Items").Find(&products).Error
	err := db.Preload("Variants").Where("uuid = ?", productUUID).First(&Product).Error

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": fmt.Sprintf("Error getting product: %v", err.Error()),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":    Product,
		"message": "succeed get product",
	})
}

func UpdateProduct(ctx *gin.Context) {
	db := database.GetDB()
	productUUID := ctx.Param("productUUID")
	// condition := false

	var productReq requests.ProductRequest
	// updatedItems := models.Item{}

	if err := ctx.ShouldBind(&productReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	Product := models.Product{
		Name: productReq.Name,
		// ImageURL: uploadResult,
		// AdminID:  adminID,
	}

	err := db.Model(&Product).Where("uuid = ?", productUUID).Updates(models.Product{
		Name: productReq.Name,
	}).Error

	if err != nil {
		fmt.Println("Error updating product data:", err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Product with uuid %v has been successfully updated", productUUID),
	})
}

func DeleteProduct(ctx *gin.Context) {
	db := database.GetDB()
	productUUID := ctx.Param("productUUID")

	Variant := models.Variant{}
	Product := models.Product{}

	var err error

	db.Transaction(func(tx *gorm.DB) error {

		if err = tx.Where("product_uuid = ?", productUUID).Delete(&Variant).Error; err != nil {
			return err
		}

		if err = tx.Where("uuid = ?", productUUID).Delete(&Product).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"data":    nil,
			"message": fmt.Sprintf("Error deleting product: %v", err.Error()),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":    nil,
		"message": fmt.Sprintf("Product with uuid %v has been successfully deleted", productUUID),
	})

}
