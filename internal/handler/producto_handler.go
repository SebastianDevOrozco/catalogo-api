package handler

import (
	"net/http"

	"catalogo-api/internal/service"
	"github.com/gin-gonic/gin"
)

// 1. ESTRUCTURA DEL HANDLER
// Guarda internamente la dependencia del servicio.
// Es la puerta de entrada para atender peticiones HTTP relacionadas con Productos.
type ProductoHandler struct {
	service service.ProductoService // Dependemos de la INTERFAZ del servicio (Inversión de Dependencias)
}

// 2. FUNCIÓN CONSTRUCTORA
// Recibe la instancia del servicio y retorna un puntero a nuestro Handler.
// Se llamará desde cmd/api/main.go.
func NewProductoHandler(service service.ProductoService) *ProductoHandler {
	return &ProductoHandler{
		service: service, // Guardamos la variable que entra por parámetro en el campo 'service' de la estructura
	}
}

// 3. MÉTODO PARA EL ENDPOINT (GET /api/v1/productos)
// (h *ProductoHandler) -> Le asigna este método al struct. 'h' funciona como 'this'.
// (c *gin.Context) -> Objeto especial de Gin que encapsula la petición HTTP (Request)
//                     y las herramientas para enviar la respuesta al cliente (Response).
func (h *ProductoHandler) ListarProductos(c *gin.Context) {
	// A. Llamamos al servicio a través de 'h.service' para ejecutar la consulta
	productos, err := h.service.ListarProductos()

	// B. Manejo de Errores
	// Si ocurrió un error (ej. falla en la DB), detenemos la ejecución y respondemos 500.
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Ocurrió un error al consultar la lista de productos",
		})
		return // ¡IMPORTANTE! El 'return' evita que el código siga ejecutándose y responda dos veces
	}

	// C. Respuesta Exitosa
	// Si no hubo error, enviamos un estado 200 OK y convertimos la lista 'productos' a JSON automáticamente
	c.JSON(http.StatusOK, productos)
}