package models

// CosechaFilter representa los parámetros opcionales de búsqueda en la URL
type CosechaFilter struct {
	ProductoID  uint    `form:"producto_id"`  // /cosechas/buscar?producto_id=1
	CategoriaID uint    `form:"categoria_id"` // /cosechas/buscar?categoria_id=2
	PrecioMin   float64 `form:"precio_min"`   // /cosechas/buscar?precio_min=1000
	PrecioMax   float64 `form:"precio_max"`   // /cosechas/buscar?precio_max=5000
}