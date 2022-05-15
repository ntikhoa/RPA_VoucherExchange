package services

import (
	"errors"
	"math"

	"github.com/RPA_VoucherExchange/custom_error"
	"github.com/RPA_VoucherExchange/entities"
	"github.com/RPA_VoucherExchange/repositories"
	viewmodel "github.com/RPA_VoucherExchange/view_model"
	"gorm.io/gorm"
)

type ProductService interface {
	Create(product entities.Product) error
	Update(product entities.Product) error
	DeleteByID(productID uint, providerID uint) error
	FindAllWithPage(providerID uint, page int, perPage int) (viewmodel.PagingMetadata, []entities.Product, error)
	FindByID(productID uint, providerID uint) (entities.Product, error)
	GetProductCount(providerID uint) (int64, error)
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
	return s.repo.Create(product)
}

func (s *productService) Update(product entities.Product) error {
	fetchedProduct, err := s.repo.FindByID(product.Model.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &custom_error.NotFoundError{}
		}
		return err
	}

	if fetchedProduct.ProviderID != product.ProviderID {
		return &custom_error.AuthorizedError{}
	}
	return s.repo.Update(product)
}

func (s *productService) DeleteByID(productID uint, providerID uint) error {
	product, err := s.repo.FindByID(productID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		return err
	}

	if product.ProviderID != providerID {
		return &custom_error.AuthorizedError{}
	}
	return s.repo.DeleteByID(productID)
}

func (s *productService) FindAllWithPage(providerID uint, page int, perPage int) (viewmodel.PagingMetadata, []entities.Product, error) {
	pagingMetadata := viewmodel.PagingMetadata{}

	count, err := s.repo.GetProductCount(providerID)
	if err != nil {
		return pagingMetadata, nil, err
	}
	d := float64(count) / float64(perPage)
	totalPages := int(math.Ceil(d))
	if page > totalPages {
		return pagingMetadata, nil, &custom_error.ExhaustedError{}
	}

	pagingMetadata = viewmodel.PagingMetadata{
		Page:         page,
		PerPage:      perPage,
		TotalPages:   totalPages,
		TotalRecords: int(count),
	}
	products, err := s.repo.FindAllWithPage(providerID, page, perPage)
	return pagingMetadata, products, err
}

func (s *productService) FindByID(productID uint, providerID uint) (entities.Product, error) {
	product, err := s.repo.FindByID(productID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return product, &custom_error.NotFoundError{}
		}
		return product, err
	}

	if product.ProviderID != providerID {
		return product, &custom_error.AuthorizedError{}
	}

	return product, nil
}

func (s *productService) GetProductCount(providerID uint) (int64, error) {
	return s.repo.GetProductCount(providerID)
}
