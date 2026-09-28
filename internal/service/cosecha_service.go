package service

import (
	"catalogo-api/internal/models"
	"catalogo-api/internal/repository"
	"errors"
)

type CosechaService interface {
	GetAll() ([]models.Cosecha, error)
	GetByID(id uint) (*models.Cosecha, error)
	Create(cosecha *models.Cosecha) error
	Update(id uint, cosecha *models.Cosecha) (*models.Cosecha, error)
	Delete(id uint) error
	Buscar(filter models.CosechaFilter) ([]models.Cosecha, error)
}

type cosechaService struct {
	repo repository.CosechaRepository
}

func NewCosechaService(repo repository.CosechaRepository) CosechaService {
	return &cosechaService{repo: repo}
}

func (s *cosechaService) GetAll() ([]models.Cosecha, error) {
	return s.repo.GetAll()
}

func (s *cosechaService) GetByID(id uint) (*models.Cosecha, error) {
	return s.repo.GetByID(id)
}

func (s *cosechaService) Create(cosecha *models.Cosecha) error {
	if cosecha.Cantidad <= 0 {
		return errors.New("la cantidad debe ser mayor a 0")
	}
	return s.repo.Create(cosecha)
}

func (s *cosechaService) Update(id uint, cosechaData *models.Cosecha) (*models.Cosecha, error) {
	cosechaExistente, err := s.repo.GetByID(id)
	if err != nil {
		return nil, errors.New("cosecha no encontrada")
	}

	// Actualizamos los campos necesarios
	cosechaExistente.Cantidad = cosechaData.Cantidad
	cosechaExistente.PrecioUnidad = cosechaData.PrecioUnidad
	cosechaExistente.FechaEstimada = cosechaData.FechaEstimada
	cosechaExistente.IdProducto = cosechaData.IdProducto
	cosechaExistente.IdEstado = cosechaData.IdEstado

	err = s.repo.Update(cosechaExistente)
	if err != nil {
		return nil, err
	}

	// Volvemos a consultar para retornar los datos con sus relaciones cargadas
	return s.repo.GetByID(id)
}

func (s *cosechaService) Delete(id uint) error {
	_, err := s.repo.GetByID(id)
	if err != nil {
		return errors.New("cosecha no encontrada")
	}
	return s.repo.Delete(id)
}

func (s *cosechaService) 	Buscar(filter models.CosechaFilter) ([]models.Cosecha, error){
	return s.repo.Buscar(filter)
}