package handler

import (
	"catalogo-api/internal/models"
	"catalogo-api/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CosechaHandler struct {
	service service.CosechaService
}

func NewCosechaHandler(service service.CosechaService) *CosechaHandler {
	return &CosechaHandler{service: service}
}

func (h *CosechaHandler) GetAll(c *gin.Context) {
	cosechas, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, cosechas)
}

func (h *CosechaHandler) GetByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	cosecha, err := h.service.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, cosecha)
}

func (h *CosechaHandler) Create(c *gin.Context) {
	var cosecha models.Cosecha
	if err := c.ShouldBindJSON(&cosecha); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.Create(&cosecha); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, cosecha)
}

func (h *CosechaHandler) Update(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var cosecha models.Cosecha
	if err := c.ShouldBindJSON(&cosecha); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	actualizada, err := h.service.Update(uint(id), &cosecha)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, actualizada)
}

func (h *CosechaHandler) Delete(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	if err := h.service.Delete(uint(id)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"mensaje": "Cosecha eliminada correctamente"})
}

func (h *CosechaHandler) Buscar(c *gin.Context) {
	var filter models.CosechaFilter

	// ShouldBindQuery lee los parámetros de la URL y los mapea automáticamente a CosechaFilter
	if err := c.ShouldBindQuery(&filter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Parámetros de búsqueda inválidos",
		})
		return
	}

	cosechas, err := h.service.Buscar(filter)
	
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error al procesar la búsqueda de cosechas",
		})
		return
	}

	c.JSON(http.StatusOK, cosechas)
}