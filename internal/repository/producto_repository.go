package repository

import(
	"catalogo-api/internal/models"
	"gorm.io/gorm"
)

type ProductoRepository interface{
	ListarProductos() ([]models.Producto, error)
}

type productoRepository struct{
	db *gorm.DB
}

func NewProductoRepository(db *gorm.DB) ProductoRepository{
	return &productoRepository{db: db}
}

func(r *productoRepository) ListarProductos() ([]models.Producto, error){
	var productos []models.Producto
	err := r.db.Preload("Categoria").Preload("UnidadMedida").Find(&productos).Error
	return productos,err
}