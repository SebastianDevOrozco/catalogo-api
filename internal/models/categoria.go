package models

import (

)

type Categoria struct {
	ID        uint    `gorm:"primaryKey;autoIncrement" json:"id"`
	Nombre    string  `gorm:"type:varchar(100);not null" json:"nombre"`
	Productos []Producto `gorm:"foreignKey:IdCategoria" json:"productos,omitempty"` // 1:N
}