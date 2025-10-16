package main

import (
	"log"
	"test_go_gin/controllers"
	"test_go_gin/database"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1. Conectar a la base de datos
	database.ConnectDB()
	defer database.DB.Close() //Asegura que la conexión se cierre al finalizar

	// 2. Crear el router de gin

	router := gin.Default()

	//3. Definir las rutas de la api
	//agrupamos las rutas bajo un prefijo /api/v1

	v1 := router.Group("/api/v1")
	{
		// Ruta opst para crear un producto
		v1.POST("/products", controllers.CreateProduct)


		// NUEVA RUTA: DELETE  para eliminar por id
		//la ruta /products/;id indica que gin capturará el valo despés de /products/
        v1.DELETE("/products/:id", controllers.DeleteProduct)


		//Aquií se pueden agregar las demas rutas como get, put, delete, etc.
	}

	// 4. Inicar el servidor
	log.Fatal(router.Run(":8080"))

}
