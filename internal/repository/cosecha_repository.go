package repository

import (
	"catalogo-api/internal/models"
	"gorm.io/gorm"
)

type CosechaRepository interface {
	GetAll() ([]models.Cosecha, error)
	GetByID(id uint) (*models.Cosecha, error)
	Create(cosecha *models.Cosecha) error
	Update(cosecha *models.Cosecha) error
	Delete(id uint) error
	Buscar(filter models.CosechaFilter) ([]models.Cosecha, error)
}

type cosechaRepository struct {
	db *gorm.DB
}

func NewCosechaRepository(db *gorm.DB) CosechaRepository {
	return &cosechaRepository{db: db}
}

func (r *cosechaRepository) GetAll() ([]models.Cosecha, error) {
	var cosechas []models.Cosecha
	// Preload carga las relaciones de Producto (con su categoría y unidad) y EstadoCosecha
	err := r.db.
		Preload("Producto.Categoria").
		Preload("Producto.UnidadMedida").
		Preload("EstadoCosecha").
		Find(&cosechas).Error
	return cosechas, err
}

func (r *cosechaRepository) GetByID(id uint) (*models.Cosecha, error) {
	var cosecha models.Cosecha
	err := r.db.
		Preload("Producto.Categoria").
		Preload("Producto.UnidadMedida").
		Preload("EstadoCosecha").
		First(&cosecha, id).Error
	if err != nil {
		return nil, err
	}
	return &cosecha, nil
}

func (r *cosechaRepository) Create(cosecha *models.Cosecha) error {
	return r.db.Create(cosecha).Error
}

func (r *cosechaRepository) Update(cosecha *models.Cosecha) error {
	return r.db.Save(cosecha).Error
}

func (r *cosechaRepository) Delete(id uint) error {
	return r.db.Delete(&models.Cosecha{}, id).Error
}

func (r *cosechaRepository) Buscar(filter models.CosechaFilter) ([]models.Cosecha, error) {

	var cosechas []models.Cosecha

	// A. Iniciamos la consulta base precargando las relaciones necesarias
	query := r.db.Model(&models.Cosecha{}).
		Preload("Producto").
		Preload("Producto.Categoria").
		Preload("Producto.UnidadMedida").
		Preload("EstadoCosecha")

	// B. Filtro opcional por Producto
	if (filter.ProductoID > 0) {
		query = query.Where("cosechas.id_producto = ?", filter.ProductoID)
	}

	// C. Filtro opcional por Categoría (Hacemos un JOIN con la tabla de productos)
	if (filter.CategoriaID > 0) {
		query = query.Joins("JOIN productos ON productos.id = cosechas.producto_id").
			Where("productos.id_categoria = ?", filter.CategoriaID)
	}

	// D. Filtros opcionales por Rango de Precios
	if (filter.PrecioMin > 0) {
		query = query.Where("cosechas.precio_unidad >= ?", filter.PrecioMin)
	}

	if (filter.PrecioMax > 0) {
		query = query.Where("cosechas.precio_unidad <= ?", filter.PrecioMax)
	}

	// E. Ejecutamos la consulta con todas las cláusulas agregadas
	err := query.Find(&cosechas).Error
	return cosechas, err
}
