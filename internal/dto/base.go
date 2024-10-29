package dto

import (
	"math"

	"gorm.io/gorm"
)

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

func (p *PagingRequestDTO) Paginate() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		p.Skip = (p.Page - 1) * p.PageSize
		return db.Offset(p.Skip).Limit(p.PageSize)
	}
}

func (p *PagingResponse) Paginate(page int, pageSize int, totalItems int) {
	p.CurrentPage = page
	p.TotalPages = int(math.Ceil(float64(float64(totalItems) / float64(pageSize))))
	p.TotalItems = totalItems
}
