
func (m *default{{.upperStartCamelObject}}Model) FindOneBy{{.upperField}}(ctx context.Context, {{.in}}) (resp *{{.upperStartCamelObject}},ok bool,err error) {
{{if .withCache}}{{.cacheKey}}
	resp = &{{.upperStartCamelObject}}{}
	err = m.QueryRowIndexCtx(ctx, &resp, {{.cacheKeyVariable}}, m.formatPrimary, func(conn *gorm.DB, v interface{}) (interface{}, error) {
		if err := conn.WithContext(ctx).Model(&{{.upperStartCamelObject}}{}).Where("{{.originalField}}", {{.lowerStartCamelField}}).First(&resp).Error; err != nil {

		}
		return
	}, m.queryPrimary)
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
{{else}}
    resp = &{{.upperStartCamelObject}}{}
	err = m.conn.WithContext(ctx).Model(&{{.upperStartCamelObject}}{}).Where("{{.originalField}}", {{.lowerStartCamelField}}).First(&resp).Error
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
{{end}}
