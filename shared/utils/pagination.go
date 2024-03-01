package utils

import (
	"gorm.io/gorm"
)

func Paginate(pagination *Pagination) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if pagination.GetPageSize() == 0 {
			return db.Order(pagination.GetSort())
		}

		return db.Offset(pagination.GetOffset()).Limit(pagination.GetPageSize()).Order(pagination.GetSort())
	}
}

type Pagination struct {
	PageSize   int         `json:"page_size,omitempty;query:limit"`
	Page       int         `json:"page,omitempty;query:page"`
	Sort       string      `json:"sort,omitempty;query:sort"`
	TotalData  int64       `json:"total_data"`
	TotalPages int         `json:"total_pages"`
	Data       interface{} `json:"data"`
}

func (p *Pagination) GetOffset() int {
	return (p.GetPage() - 1) * p.GetPageSize()
}

func (p *Pagination) GetPageSize() int {
	/* if p.PageSize == 0 {
		p.PageSize = 10
	} */
	return p.PageSize
}

func (p *Pagination) GetPage() int {
	if p.Page == 0 {
		p.Page = 1
	}
	return p.Page
}

func (p *Pagination) GetSort() string {
	if p.Sort == "" {
		p.Sort = "ID desc"
	}
	return p.Sort
}
