package dto

type Order string

const (
	ORDER_ASC  Order = "asc"
	ORDER_DESC Order = "desc"
)

type BaseRequestDTO struct {
	TraceId string `form:"traceId" json:"traceId"`
}

type (
	PagingRequestDTO struct {
		BaseRequestDTO
		Page     int `form:"page"`
		PageSize int `form:"pageSize" validate:"required"`
		Skip     int
		Order    Order  `form:"order"`
		OrderBy  string `form:"orderBy"`
	}

	PagingResponse struct {
		List        []*any `json:"list"`
		CurrentPage int    `json:"currentPage"`
		TotalPages  int    `json:"totalPages"`
		TotalItems  int    `json:"totalItems"`
	}
)
