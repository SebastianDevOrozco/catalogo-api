package main

import (
	"fmt"

	"catalogo-api/internal/handler"
	"catalogo-api/internal/repository"
	"catalogo-api/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("=== Iniciando servicio Catalogo API ===")

	// 1. Inicializar conexión a MySQL y ejecutar AutoMigrate
	// Guardamos la conexión en la variable 'db' para pasársela al repositorio
	db := repository.InitDB()

	// 2. Instanciar Capas (Inyección de Dependencias)
	// Cosechas
	cosechaRepo := repository.NewCosechaRepository(db)
	cosechaService := service.NewCosechaService(cosechaRepo)
	cosechaHandler := handler.NewCosechaHandler(cosechaService)

	// Productos
	productoRepo := repository.NewProductoRepository(db)
	productoService := service.NewProductoService(productoRepo)
	productoHandler := handler.NewProductoHandler(productoService) 

	// 3. Crear servidor Gin
	r := gin.Default()

	// Endpoint de prueba de estado
	r.GET("/api/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "UP",
			"service": "Catalogo API",
		})
	})

	// 4. Registrar Rutas
	api := r.Group("/api/v1")
	{
		// Rutas de Cosechas
		cosechas := api.Group("/cosechas")
		{
			cosechas.GET("", cosechaHandler.GetAll)
			cosechas.GET("/:id", cosechaHandler.GetByID)
			cosechas.GET("/buscar", cosechaHandler.Buscar)
			cosechas.POST("", cosechaHandler.Create)
			cosechas.PUT("/:id", cosechaHandler.Update)
			cosechas.DELETE("/:id", cosechaHandler.Delete)
		}

		// Rutas de Productos
		productos := api.Group("/productos")
		{
			productos.GET("", productoHandler.ListarProductos) 
		}
	}

	// 5. Levantar servidor HTTP
	r.Run(":8080")
}