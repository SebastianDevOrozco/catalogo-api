package repository

import (
	"fmt"
	"log"

	"catalogo-api/internal/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	// Cadena de conexión DSN (Data Source Name)

	dsn := "root:servicioCatalogo@tcp(127.0.0.1:3308)/AgronetCatalogoService?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error al conectar a la base de datos: %v", err)
	}

	fmt.Println("Conexión exitosa a MySQL")

	// Equivale a 'dotnet ef database update'
	err = db.AutoMigrate(
		&models.Categoria{},
		&models.UnidadMedida{},
		&models.EstadoCosecha{},
		&models.Producto{},
		&models.Cosecha{},
	)

	if err != nil {
		log.Fatalf("Error al migrar las tablas: %v", err)
	}

	fmt.Println("Tablas migradas correctamente")

	DB = db
	return db
}