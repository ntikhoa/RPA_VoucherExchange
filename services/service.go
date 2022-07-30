package services

import (
	"math"

	"github.com/RPA_VoucherExchange/constants"
	"github.com/RPA_VoucherExchange/custom_error"
	viewmodel "github.com/RPA_VoucherExchange/view_model"
)

type Count func(uint) (int64, error)

func paging(countFunc Count, providerID uint, page int, perPage int) (viewmodel.PagingMetadata, error) {
	var pagingMetadata viewmodel.PagingMetadata

	count, err := countFunc(providerID)
	if err != nil {
		return pagingMetadata, err
	}
	d := float64(count) / float64(perPage)
	totalPages := int(math.Ceil(d))
	if page > totalPages {
		return pagingMetadata, custom_error.NewNotFoundError(constants.EXHAUSTED_ERROR)
	}

	pagingMetadata = viewmodel.PagingMetadata{
		Page:         page,
		PerPage:      perPage,
		TotalPages:   totalPages,
		TotalRecords: int(count),
	}
	return pagingMetadata, nil
}

func paging2(count int64, page int, perPage int) (viewmodel.PagingMetadata, error) {
	var pagingMetadata viewmodel.PagingMetadata

	d := float64(count) / float64(perPage)
	totalPages := int(math.Ceil(d))
	if page > totalPages {
		return pagingMetadata, custom_error.NewNotFoundError(constants.EXHAUSTED_ERROR)
	}

	pagingMetadata = viewmodel.PagingMetadata{
		Page:         page,
		PerPage:      perPage,
		TotalPages:   totalPages,
		TotalRecords: int(count),
	}
	return pagingMetadata, nil
}
