package models

import (
	"time"
)

type Cosecha struct {
	ID            uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Cantidad      float64   `gorm:"type:decimal(10,2);not null" json:"cantidad"`
	PrecioUnidad  float64   `gorm:"type:decimal(10,2)" json:"precio_unidad"`
	FechaEstimada time.Time `gorm:"type:date;not null" json:"fecha_estimada"`

	// Relación con Producto
	IdProducto    uint      `gorm:"not null" json:"id_producto"`
	Producto      Producto  `gorm:"foreignKey:IdProducto;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;" json:"producto,omitempty"`

	// Relación con EstadoCosecha
	IdEstado      uint          `gorm:"not null" json:"id_estado"`
	EstadoCosecha EstadoCosecha `gorm:"foreignKey:IdEstado;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;" json:"estado_cosecha,omitempty"`

	CreadoEn      time.Time `gorm:"autoCreateTime" json:"creado_en"`
}