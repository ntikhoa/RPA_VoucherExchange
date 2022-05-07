package services

import (
	"errors"
	"log"

	"github.com/RPA_VoucherExchange/entities"
	"github.com/RPA_VoucherExchange/repositories"
)

type ProductService interface {
	Create(product entities.Product) error
	Update(product entities.Product) error
	DeleteById(productID uint) error
	FindAll(providerID uint) ([]entities.Product, error)
}

type productService struct {
	repo repositories.ProductRepo
}

func NewProductService(repo repositories.ProductRepo) ProductService {
	return &productService{
		repo: repo,
	}
}

func (s *productService) Create(product entities.Product) error {
	if err := s.repo.Create(product); err != nil {
		log.Fatalln("GORM ERROR:", err)
		return errors.New("cannot create product")
	}
	return nil
}

func (s *productService) Update(product entities.Product) error {
	if err := s.repo.Update(product); err != nil {
		log.Fatalln("GORM ERROR:", err)
		return errors.New("cannot update product")
	}
	return nil
}

func (s *productService) DeleteById(productID uint) error {
	if err := s.repo.DeleteById(productID); err != nil {
		log.Fatalln("GORM ERROR:", err)
		return errors.New("cannot update product")
	}
	return nil
}

func (s *productService) FindAll(providerID uint) ([]entities.Product, error) {
	products, err := s.repo.FindAll(providerID)
	if err != nil {
		log.Fatalln("GORM ERROR:", err)
		err = errors.New("cannot get products")
	}
	return products, err
}
