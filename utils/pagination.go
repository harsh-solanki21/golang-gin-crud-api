package utils

import (
	"math"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PaginationData struct {
	Limit      int   `json:"limit,omitempty"`
	Page       int   `json:"page,omitempty"`
	TotalRows  int64 `json:"total_rows"`
	TotalPages int   `json:"total_pages"`
}

type PaginatedResponse struct {
	Data       interface{}    `json:"data"`
	Pagination PaginationData `json:"pagination"`
}

type Pagination struct {
	Limit int
	Page  int
	Sort  string
}

func GeneratePaginationFromRequest(c *gin.Context) Pagination {
	// Initializing default
	limit := 10
	page := 1
	sort := "created_at desc"

	var err error
	if c.Query("limit") != "" {
		limit, err = strconv.Atoi(c.Query("limit"))
		if err != nil || limit <= 0 {
			limit = 10
		}
	}
	if c.Query("page") != "" {
		page, err = strconv.Atoi(c.Query("page"))
		if err != nil || page <= 0 {
			page = 1
		}
	}
	if c.Query("sort") != "" {
		sort = c.Query("sort")
	}

	return Pagination{
		Limit: limit,
		Page:  page,
		Sort:  sort,
	}
}

func (p *Pagination) GetOffset() int {
	return (p.GetPage() - 1) * p.GetLimit()
}

func (p *Pagination) GetLimit() int {
	if p.Limit <= 0 {
		p.Limit = 10
	}
	return p.Limit
}

func (p *Pagination) GetPage() int {
	if p.Page <= 0 {
		p.Page = 1
	}
	return p.Page
}

func (p *Pagination) GetSort() string {
	if p.Sort == "" {
		p.Sort = "created_at desc"
	}
	return p.Sort
}

func (p *Pagination) GenerateResponse(data interface{}, totalRows int64) PaginatedResponse {
	totalPages := int(math.Ceil(float64(totalRows) / float64(p.GetLimit())))

	return PaginatedResponse{
		Data: data,
		Pagination: PaginationData{
			Limit:      p.GetLimit(),
			Page:       p.GetPage(),
			TotalRows:  totalRows,
			TotalPages: totalPages,
		},
	}
}
