package models

import (

)

type Producto struct {
	ID             uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Nombre         string `gorm:"type:varchar(100);not null" json:"nombre"`
	Descripcion    string `gorm:"type:text" json:"descripcion"`
	ImagenUrl      string `gorm:"type:varchar(255)" json:"imagen_url,omitempty"`

	// Llave foránea para Categoría
	IdCategoria    uint      `gorm:"not null" json:"id_categoria"`
	Categoria      Categoria `gorm:"foreignKey:IdCategoria;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;" json:"categoria,omitempty"`

	// Llave foránea para UnidadMedida
	IdUnidadMedida uint         `gorm:"not null" json:"id_unidad_medida"`
	UnidadMedida   UnidadMedida `gorm:"foreignKey:IdUnidadMedida;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;" json:"unidad_medida,omitempty"`
	
	Cosechas []Cosecha `gorm:"foreignKey:IdProducto" json:"cosechas,omitempty"` // 1:N
}