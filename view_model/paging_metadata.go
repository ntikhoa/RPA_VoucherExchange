package viewmodel

type PagingMetadata struct {
	Page         int `json:"page"`
	PerPage      int `json:"per_page"`
	TotalPages   int `json:"total_pages"`
	TotalRecords int `json:"total_records"`
}
