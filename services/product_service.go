package services

import (
	"errors"
	"math"
	"strconv"

	"github.com/RPA_VoucherExchange/constants"
	"github.com/RPA_VoucherExchange/custom_error"
	"github.com/RPA_VoucherExchange/dto"
	"github.com/RPA_VoucherExchange/entities"
	"github.com/RPA_VoucherExchange/repositories"
	viewmodel "github.com/RPA_VoucherExchange/view_model"
	"gorm.io/gorm"
)

type ProductService interface {
	Create(productDTO dto.ProductDTO, providerID uint) error
	Update(productDTO dto.ProductDTO, providerID uint, productID uint) error
	DeleteByID(productID uint, providerID uint) error
	DeleteByIDs(productIDs []uint, providerID uint) error
	FindAllWithPage(
		providerID uint,
		page int,
		perPage int) (viewmodel.PagingMetadata, []entities.Product, error)
	FindByID(productID uint, providerID uint) (entities.Product, error)
	GetCount(providerID uint) (int64, error)
	CheckExistence(productIDs []uint) error
	GetAll(providerID uint) ([]entities.Product, error)
}

type productService struct {
	repo repositories.ProductRepo
}

func NewProductService(repo repositories.ProductRepo) ProductService {
	return &productService{
		repo: repo,
	}
}

func (s *productService) Create(productDTO dto.ProductDTO, providerID uint) error {
	product := entities.NewProduct(productDTO, providerID)
	return s.repo.Create(product)
}

func (s *productService) Update(productDTO dto.ProductDTO, providerID uint, productID uint) error {
	product := entities.NewProduct(productDTO, providerID)
	product.Model.ID = productID
	fetchedProduct, err := s.repo.FindByID(product.Model.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return custom_error.NewNotFoundError(constants.NOT_FOUND_ERROR)
		}
		return err
	}

	if fetchedProduct.ProviderID != product.ProviderID {
		return custom_error.NewForbiddenError(constants.AUTHORIZE_ERROR)
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
		return custom_error.NewForbiddenError(constants.AUTHORIZE_ERROR)
	}
	return s.repo.DeleteByID(productID)
}

func (s *productService) FindAllWithPage(providerID uint, page int, perPage int) (viewmodel.PagingMetadata, []entities.Product, error) {
	pagingMetadata := viewmodel.PagingMetadata{}

	count, err := s.repo.Count(providerID)
	if err != nil {
		return pagingMetadata, nil, err
	}
	d := float64(count) / float64(perPage)
	totalPages := int(math.Ceil(d))
	if page > totalPages {
		return pagingMetadata, nil, custom_error.NewNotFoundError(constants.EXHAUSTED_ERROR)
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
			return product, custom_error.NewNotFoundError(constants.NOT_FOUND_ERROR)
		}
		return product, err
	}

	if product.ProviderID != providerID {
		return product, custom_error.NewForbiddenError(constants.AUTHORIZE_ERROR)
	}

	return product, nil
}

func (s *productService) GetCount(providerID uint) (int64, error) {
	return s.repo.Count(providerID)
}

func (s *productService) CheckExistence(productIDs []uint) error {
	fetchedID, err := s.repo.CheckExistence(productIDs)
	if err != nil {
		return err
	}

	invalidProductIDs := extractInvalidIDs(fetchedID, productIDs)

	if len(invalidProductIDs) > 0 {
		invalidIDstr := convertToStringError(invalidProductIDs)
		return custom_error.NewConflictError("invalid product ids: " + invalidIDstr)
	}
	return nil
}

func (s *productService) GetAll(providerID uint) ([]entities.Product, error) {
	products, err := s.repo.GetAll(providerID)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (s *productService) DeleteByIDs(productIDs []uint, providerID uint) error {
	products, err := s.repo.FindByIDs(productIDs)
	if err != nil {
		return err
	}
	var filteredProductIDs []uint
	if len(products) > 0 {
		for _, product := range products {
			if product.ProviderID == providerID {
				filteredProductIDs = append(filteredProductIDs, product.Model.ID)
			}
		}
	}
	if len(filteredProductIDs) > 0 {
		return s.repo.DeleteByIDs(filteredProductIDs)
	}
	return nil
}

func extractInvalidIDs(fetchedIDs []uint, requestIDs []uint) []uint {
	invalidProductIDs := []uint{}
	mapFetchedID := make(map[uint]interface{}, len(fetchedIDs))
	for _, x := range fetchedIDs {
		mapFetchedID[x] = nil
	}

	for _, id := range requestIDs {
		if _, ok := mapFetchedID[id]; !ok {
			invalidProductIDs = append(invalidProductIDs, id)
		}
	}
	return invalidProductIDs
}

func convertToStringError(invalidIDs []uint) string {
	var invalidStr string
	if len(invalidIDs) > 0 {
		invalidStr = strconv.FormatUint(uint64(invalidIDs[0]), 10)
		for i, id := range invalidIDs {
			if i == 0 {
				continue
			}
			invalidStr += ", " + strconv.FormatUint(uint64(id), 10)
		}
	}
	return invalidStr
}
