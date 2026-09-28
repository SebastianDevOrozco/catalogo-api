package service

import (
	"catalogo-api/internal/models"
	"catalogo-api/internal/repository"
)

// 1. Interfaz pública que consumirá el Handler
type ProductoService interface {
	ListarProductos() ([]models.Producto, error)
}

// 2. Estructura privada que guarda la dependencia del repositorio
type productoService struct {
	repo repository.ProductoRepository // Con 'P' mayúscula
}

// 3. Función constructora (Inyecta el repositorio)
func NewProductoService(repo repository.ProductoRepository) ProductoService {
	return &productoService{repo: repo}
}

// 4. Método que implementa la interfaz (usa 's' como equivalente a 'this')
func (s *productoService) ListarProductos() ([]models.Producto, error) {
	return s.repo.ListarProductos()
}