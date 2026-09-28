package models

import (

)

type UnidadMedida struct {
	ID        uint    `gorm:"primaryKey;autoIncrement" json:"id"`
	Prefijo   string  `gorm:"type:varchar(5);not null" json:"prefijo"`
	Productos []Producto `gorm:"foreignKey:IdUnidadMedida" json:"productos,omitempty"` 
}