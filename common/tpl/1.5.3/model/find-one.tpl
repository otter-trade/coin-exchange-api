
func (m *default{{.upperStartCamelObject}}Model) FindOne(ctx context.Context, {{.lowerStartCamelPrimaryKey}} {{.dataType}}) (resp *{{.upperStartCamelObject}},ok bool,err error) {
	resp = &{{.upperStartCamelObject}}{}
{{if .withCache}}{{.cacheKey}}
	err = m.QueryCtx(ctx, &resp, {{.cacheKeyVariable}}, func(conn *gorm.DB, v interface{}) error {
    		return conn.WithContext(ctx).Model(&{{.upperStartCamelObject}}{}).Where("{{.originalPrimaryKey}} = ?", {{.lowerStartCamelPrimaryKey}}).First(&resp).Error
    	})
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
	err = m.conn.WithContext(ctx).Model(&{{.upperStartCamelObject}}{}).Where("{{.originalPrimaryKey}} = ?", {{.lowerStartCamelPrimaryKey}}).First(&resp).Error
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

