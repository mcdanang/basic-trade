package controllers

import (
	"basic-trade/database"
	models "basic-trade/models/entity"
	"fmt"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

	paramPairs := ctx.Request.URL.Query()
	searchWord := ""

	if paramPairs["name"] != nil {
		searchWord = paramPairs["name"][0]
	}

	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(ctx.DefaultQuery("size", "5"))

	offset := (page - 1) * size

	variants := []models.Variant{}

	var totalItems int64
	err := db.Where("variant_name LIKE ?", "%"+searchWord+"%").Find(&variants).Count(&totalItems).Error

	err = db.Offset(offset).Limit(size).Where("variant_name LIKE ?", "%"+searchWord+"%").Find(&variants).Error

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": fmt.Sprintf("Error getting all variant: %v", err.Error()),
		})
		return
	}

	totalPages := int(math.Ceil(float64(totalItems) / float64(size)))

	ctx.JSON(http.StatusOK, gin.H{
		"data":       variants,
		"totalItems": totalItems,
		"totalPages": totalPages,
		"message":    "succeed get all variant",
	})
}

func GetVariantByUUID(ctx *gin.Context) {
	db := database.GetDB()

	Variant := models.Variant{}
	variantUUID := ctx.Param("variantUUID")
	// err := db.Preload("Items").Find(&products).Error
	err := db.Where("uuid = ?", variantUUID).First(&Variant).Error

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": fmt.Sprintf("Error getting variant: %v", err.Error()),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":    Variant,
		"message": "succeed get variant",
	})
}

func UpdateVariant(ctx *gin.Context) {
	db := database.GetDB()
	variantUUID := ctx.Param("variantUUID")

	var Variant models.Variant

	if err := ctx.ShouldBind(&Variant); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := db.Model(&Variant).Where("uuid = ?", variantUUID).Updates(models.Variant{
		VariantName: Variant.VariantName,
		Quantity:    Variant.Quantity,
	}).Error

	if err != nil {
		fmt.Println("Error updating variant data:", err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Variant with uuid %v has been successfully updated", variantUUID),
	})
}

func DeleteVariant(ctx *gin.Context) {
	db := database.GetDB()
	variantUUID := ctx.Param("variantUUID")

	Variant := models.Variant{}

	var err error

	db.Transaction(func(tx *gorm.DB) error {

		if err = tx.Where("uuid = ?", variantUUID).Delete(&Variant).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"data":    nil,
			"message": fmt.Sprintf("Error deleting variant: %v", err.Error()),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":    nil,
		"message": fmt.Sprintf("Variant with uuid %v has been successfully deleted", variantUUID),
	})

}
