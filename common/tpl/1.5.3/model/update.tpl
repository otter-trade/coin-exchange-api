
func (m *default{{.upperStartCamelObject}}Model) Update(ctx context.Context, data *{{.upperStartCamelObject}}) (rowsAffected int64,err error) {
	{{if .withCache}}{{.keys}}
    err = m.ExecCtx(ctx, func(conn *gorm.DB) error {
        result := conn.WithContext(ctx).Save(data)
		return result.RowsAffected, result.Error
	}, {{.keyValues}})
	return
	{{else}}
	conn := m.conn.WithContext(ctx).Begin()
    defer func() {
        if r := recover(); r != nil {
            conn.Rollback()
        }
    }()
    info := &{{.upperStartCamelObject}}{}
    if err = conn.Clauses(clause.Locking{Strength: "UPDATE"}).First(&info, data.Id).Error; err != nil {
        conn.Rollback()
        return
    }

    result :=conn.Table(m.tableName()).Updates(data)
    err = result.Error

    if err != nil {
        conn.Rollback()
        return
    }

    if err = conn.Commit().Error; err != nil {
        conn.Rollback()
        return
    }

    return result.RowsAffected, result.Error
	{{end}}
}
