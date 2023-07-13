package {{.pkg}}
{{if .withCache}}
import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
    "context"
    "github.com/otter-trade/coin-exchange-api/common/pagination"
	"gorm.io/gorm"
)
{{else}}
import (
    "context"
    "fmt"
    "github.com/otter-trade/coin-exchange-api/common/pagination"
    "github.com/otter-trade/coin-exchange-api/common/util"
    "gorm.io/gorm"
)
{{end}}
var _ {{.upperStartCamelObject}}Model = (*custom{{.upperStartCamelObject}}Model)(nil)

type (
	// {{.upperStartCamelObject}}Model is an interface to be customized, add more methods here,
	// and implement the added methods in custom{{.upperStartCamelObject}}Model.
	{{.upperStartCamelObject}}Model interface {
		{{.lowerStartCamelObject}}Model
		{{.upperStartCamelObject}}Search(ctx context.Context, params *{{.upperStartCamelObject}}SearchReq) (int64, []*{{.upperStartCamelObject}}, error)
        {{.upperStartCamelObject}}List(ctx context.Context, params *{{.upperStartCamelObject}}ListReq) ([]*{{.upperStartCamelObject}}, error)
        {{.upperStartCamelObject}}Details(ctx context.Context, params *{{.upperStartCamelObject}}DetailsReq) (*{{.upperStartCamelObject}},bool, error)
        {{.upperStartCamelObject}}Total(ctx context.Context, params *{{.upperStartCamelObject}}TotalReq) (int64, error)
        {{.upperStartCamelObject}}Updates(ctx context.Context, params *{{.upperStartCamelObject}}) (int64, error)
        {{.upperStartCamelObject}}Upsert(ctx context.Context, params *{{.upperStartCamelObject}}) (int64, error)
	}

	custom{{.upperStartCamelObject}}Model struct {
		*default{{.upperStartCamelObject}}Model
	}
)

// New{{.upperStartCamelObject}}Model returns a model for the database table.
func New{{.upperStartCamelObject}}Model(conn *gorm.DB{{if .withCache}}, c cache.CacheConf{{end}}) {{.upperStartCamelObject}}Model {
	return &custom{{.upperStartCamelObject}}Model{
		default{{.upperStartCamelObject}}Model: new{{.upperStartCamelObject}}Model(conn{{if .withCache}}, c{{end}}),
	}
}

func (m *default{{.upperStartCamelObject}}Model) {{.upperStartCamelObject}}Search(ctx context.Context, params *{{.upperStartCamelObject}}SearchReq) (count int64,resp []*{{.upperStartCamelObject}},err error) {
	conn := m.conn.WithContext(ctx).Table(m.tableName())
	resp = make([]*{{.upperStartCamelObject}}, 0)

    if params.OrderField != "" {
		params.OrderField = util.Camel2Case(params.OrderField)
	}

	count = int64(0)
	err = conn.Count(&count).Error
	if err != nil {
		return
	}

	page := pagination.Pagination{
		Page:     int(params.PageNo),
		PageSize: int(params.PageSize),
	}
	conn = page.PageLimit(conn)

	if params.OrderField != "" && params.OrderParam != "" {
		conn = conn.Order(fmt.Sprintf("%s %s", params.OrderField, params.OrderParam))
	} else {
		conn = conn.Order("id desc")
	}

	err = conn.Select({{.lowerStartCamelObject}}Rows).Find(&resp).Error
	if err != nil {
		return
	}

	return
}

func (m *default{{.upperStartCamelObject}}Model) {{.upperStartCamelObject}}List(ctx context.Context, params *{{.upperStartCamelObject}}ListReq) (resp []*{{.upperStartCamelObject}},err error) {
	conn := m.conn.WithContext(ctx).Table(m.tableName())
	resp = make([]*{{.upperStartCamelObject}}, 0)

	err = conn.Select({{.lowerStartCamelObject}}Rows).Find(&resp).Error
	if err != nil {
		return
	}

	return
}

func (m *default{{.upperStartCamelObject}}Model) {{.upperStartCamelObject}}Details(ctx context.Context, params *{{.upperStartCamelObject}}DetailsReq) (resp *{{.upperStartCamelObject}},ok bool,err error) {
	conn := m.conn.WithContext(ctx).Table(m.tableName())
	resp = &{{.upperStartCamelObject}}{}

	err = conn.Select({{.lowerStartCamelObject}}Rows).First(&resp).Error
	if err != nil {
        if err == gorm.ErrRecordNotFound {
            err = nil
            return
        }
        return
    }
    ok = true
    return
}

func (m *default{{.upperStartCamelObject}}Model) {{.upperStartCamelObject}}Updates(ctx context.Context, params *{{.upperStartCamelObject}}) (rowsAffected int64,err error) {
	conn := m.conn.WithContext(ctx).Begin()

	if params.Id != 0 {
		conn = conn.Where("id = ?", params.Id)
	}

	result := conn.Table(m.tableName()).Updates(params)
	if result.Error != nil {
		conn.Rollback()
	} else {
		conn.Commit()
	}
	return result.RowsAffected, result.Error
}

func (m *default{{.upperStartCamelObject}}Model) {{.upperStartCamelObject}}Total(ctx context.Context, params *{{.upperStartCamelObject}}TotalReq) (count int64,err error) {
	conn := m.conn.WithContext(ctx).Table(m.tableName())

	if params.StartTime != 0 && params.EndTime != 0 {
        conn = conn.Where("create_time between ? and ?", params.StartTime, params.EndTime)
    } else {
        if params.StartTime != 0 {
            conn = conn.Where("create_time > ? ", params.StartTime)
        }

        if params.EndTime != 0 {
            conn = conn.Where("create_time < ?", params.EndTime)
        }
    }

	err = conn.Count(&count).Error

	return
}

func (m *default{{.upperStartCamelObject}}Model) {{.upperStartCamelObject}}Upsert(ctx context.Context, params *{{.upperStartCamelObject}}) (int64, error) {
	result := m.conn.WithContext(ctx).Table(m.table).Select({{.lowerStartCamelObject}}Rows).Where("id = ?", params.Id).Updates(&params)
	if result.RowsAffected == 0 {
		result = m.conn.WithContext(ctx).Table(m.table).Save(&params)
	}
	return result.RowsAffected, result.Error
}