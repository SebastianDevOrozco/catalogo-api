package models

import (

)

type EstadoCosecha struct {
	ID        uint    `gorm:"primaryKey;autoIncrement" json:"id"`
	Nombre    string  `gorm:"type:varchar(100);not null" json:"nombre"`
	Cosechas []Cosecha `gorm:"foreignKey:IdEstado" json:"cosechas,omitempty"` // 1:N
}