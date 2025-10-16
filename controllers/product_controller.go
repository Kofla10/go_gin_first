package controllers

import (
	"context"
	"net/http"
	"strconv"

	"test_go_gin/database"
	"test_go_gin/models"

	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {
	var req models.CreateProducRequest

	//1. validar la entrada del json
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//2. insertar en la db
	var newID int

	query := `INSERT INTO products (name, price, description) VALUES ($1, $2, $3) RETURNING id`

	err := database.DB.QueryRow(
		context.Background(),
		query,
		req.Name,
		req.Price,
		req.Description).Scan(&newID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Product created successfully",
		"id":      newID,
	})
}

func DeleteProduct(c *gin.Context) {
	// 1. Oberner el ID de los parámetos de la url
	idParam := c.Param("id")

	// 2. Convertir el ID a entero (int)
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	// 3. Ejecutar la sentencia DELETE en postgresql
	query := `DELETE FROM products WHERE id = $1`

	// Exec ejecuta la consulta y devuelve el resultado
	result, err := database.DB.Exec(context.Background(), query, id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete product"})
		return
	}

	// 4. Verificar si se eliminó alguna fila
	if result.RowsAffected() == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.Status(http.StatusNoContent)
}
