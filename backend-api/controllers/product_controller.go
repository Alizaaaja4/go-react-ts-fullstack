package controllers

import (
	"net/http"
	"santrikoding/backend-api/database"
	"santrikoding/backend-api/helpers"
	"santrikoding/backend-api/models"
	"santrikoding/backend-api/structs"

	"github.com/gin-gonic/gin"
)

func FindProducts(c *gin.Context) {

	// inisialisasi slice untuk menampung data product
	var products []models.Product

	// ambil data product dari database
	database.DB.Find(&products)

	// kirimkan response sukses dengan data product
	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Lists Data Products",
		Data:    products,
	})
}

func FindProductById(c *gin.Context) {

	// ambil ID product dari parameter URL
	id := c.Param("id")

	// inisialisasi product
	var product models.Product

	// cari product berdasarkan ID
	if err := database.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, structs.ErrorResponse{
			Success: false,
			Message: "Product not found",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	// kirimkan response sukses dengan data product
	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Product Found",
		Data: structs.ProductResponse{
			Id:        product.Id,
			Name:      product.Name,
			Price:     product.Price,
			Stock:     product.Stock,
			CreatedAt: product.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: product.UpdatedAt.Format("2006-01-02 15:04:05"),
		},
	})
}

func CreateProduct(c *gin.Context) {

	//struct product request
	var req = structs.ProductCreateRequest{}

	// Bind JSON request ke struct ProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
			Success: false,
			Message: "Validation Errors",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	// Inisialisasi product baru
	product := models.Product{
		Name:      req.Name,
		Price:     req.Price,
		Stock:     req.Stock,
		CreatedBy: req.CreatedBy,
	}

	// simpan product ke database
	if err := database.DB.Create(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
			Success: false,
			Message: "Failed to create product",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	// kirimkan response sukses
	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Product Created",
		Data: structs.ProductResponse{
			Id:        product.Id,
			Name:      product.Name,
			Price:     product.Price,
			Stock:     product.Stock,
			CreatedAt: product.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: product.UpdatedAt.Format("2006-01-02 15:04:05"),
		},
	})
}

func UpdateProduct(c *gin.Context) {

	// Ambil ID product dari parameter URL
	id := c.Param("id")

	// Inisialisasi product
	var product models.Product

	// Cari product berdasarkan ID
	if err := database.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, structs.ErrorResponse{
			Success: false,
			Message: "Product not found",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	// struct product request
	var req = structs.ProductUpdateRequest{}

	// Bind JSON request ke struct ProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
			Success: false,
			Message: "Validation Errors",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	// Update product di database
	product.Name = req.Name
	product.Price = req.Price
	product.Stock = req.Stock
	if err := database.DB.Save(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
			Success: false,
			Message: "Failed to update product",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	// Kirimkan response sukses
	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Product Updated",
		Data: structs.ProductResponse{
			Id:        product.Id,
			Name:      product.Name,
			Price:     product.Price,
			Stock:     product.Stock,
			CreatedAt: product.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: product.UpdatedAt.Format("2006-01-02 15:04:05"),
		},
	})
}

func DeleteProduct(c *gin.Context) {

	// Ambil ID product dari parameter URL
	id := c.Param("id")

	// Inisialisasi product
	var product models.Product

	// Cari product berdasarkan ID
	if err := database.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, structs.ErrorResponse{
			Success: false,
			Message: "Product not found",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	// Hapus product dari database
	if err := database.DB.Delete(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
			Success: false,
			Message: "Failed to delete product",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	// Kirimkan response sukses
	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Product Deleted",
	})
}
