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
	CreateProduct(productDTO dto.ProductDTO, providerID uint) error
	UpdateProduct(productDTO dto.ProductDTO, providerID uint, productID uint) error
	DeleteProductByID(productID uint, providerID uint) error
	FindAllProductWithPage(providerID uint, page int, perPage int) (viewmodel.PagingMetadata, []entities.Product, error)
	FindProductByID(productID uint, providerID uint) (entities.Product, error)
	GetProductCount(providerID uint) (int64, error)
	CheckProductsExist(productIDs []uint) error
}

type productService struct {
	repo repositories.ProductRepo
}

func NewProductService(repo repositories.ProductRepo) ProductService {
	return &productService{
		repo: repo,
	}
}

func (s *productService) CreateProduct(productDTO dto.ProductDTO, providerID uint) error {
	product := productDTO.ToEntity(providerID)
	return s.repo.CreateProduct(product)
}

func (s *productService) UpdateProduct(productDTO dto.ProductDTO, providerID uint, productID uint) error {
	product := productDTO.ToEntity(providerID)
	product.Model.ID = productID
	fetchedProduct, err := s.repo.FindProductByID(product.Model.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return custom_error.NewNotFoundError(constants.NOT_FOUND_ERROR)
		}
		return err
	}

	if fetchedProduct.ProviderID != product.ProviderID {
		return custom_error.NewForbiddenError(constants.AUTHORIZE_ERROR)
	}
	return s.repo.UpdateProduct(product)
}

func (s *productService) DeleteProductByID(productID uint, providerID uint) error {
	product, err := s.repo.FindProductByID(productID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		return err
	}

	if product.ProviderID != providerID {
		return custom_error.NewForbiddenError(constants.AUTHORIZE_ERROR)
	}
	return s.repo.DeleteProductByID(productID)
}

func (s *productService) FindAllProductWithPage(providerID uint, page int, perPage int) (viewmodel.PagingMetadata, []entities.Product, error) {
	pagingMetadata := viewmodel.PagingMetadata{}

	count, err := s.repo.GetProductCount(providerID)
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
	products, err := s.repo.FindAllProductWithPage(providerID, page, perPage)
	return pagingMetadata, products, err
}

func (s *productService) FindProductByID(productID uint, providerID uint) (entities.Product, error) {
	product, err := s.repo.FindProductByID(productID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return product, &custom_error.NotFoundError{}
		}
		return product, err
	}

	if product.ProviderID != providerID {
		return product, custom_error.NewForbiddenError(constants.AUTHORIZE_ERROR)
	}

	return product, nil
}

func (s *productService) GetProductCount(providerID uint) (int64, error) {
	return s.repo.GetProductCount(providerID)
}

func (s *productService) CheckProductsExist(productIDs []uint) error {
	fetchedID, err := s.repo.CheckProductsExist(productIDs)
	if err != nil {
		return errors.New("cannot fetch products")
	}

	invalidProductIDs := extractInvalidIDs(fetchedID, productIDs)

	if len(invalidProductIDs) > 0 {
		return errors.New("invalid product ids: " + convertToStringError(invalidProductIDs))
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
