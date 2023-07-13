package pagination

import "gorm.io/gorm"

type (
	Pagination struct {
		Page     int `required:"false" desc:"第几页"`
		PageSize int `required:"false" desc:"每页条数,默认10"`
	}
)

// 自定义类型限制为最大值不能大于20
type PageLimit int

func (p PageLimit) MustLt20() int {
	if p == 0 {
		return 10
	} else if p > 20 {
		return 20
	}
	return int(p)
}

func (p PageLimit) MustLt50() int {
	if p == 0 {
		return 20
	} else if p > 50 {
		return 50
	}
	return int(p)
}

func (p PageLimit) MustRt(page int) int {
	pageSize := p.MustLt50()
	if page == 0 {
		page = 1
	}
	num := pageSize * (page - 1)
	return int(num)
}

func (p PageLimit) BuildPagination(query *gorm.DB) *gorm.DB {
	query = query.Limit(int(p.MustLt50()))
	return query
}

func (p *Pagination) Start() int {
	if p.Page > 0 && p.PageSize > 0 {
		page := p.Page
		perPage := p.PageSize
		return (page - 1) * perPage
	}
	return 0
}

func (p *Pagination) PageLimit(q *gorm.DB) *gorm.DB {
	if p.Page <= 0 {
		p.Page = 1
	}
	if p.PageSize <= 0 {
		p.PageSize = 10
	}
	return q.Offset(p.Start()).Limit(p.PageSize)
}
