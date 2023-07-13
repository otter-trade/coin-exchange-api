
func (m *default{{.upperStartCamelObject}}Model) Delete(ctx context.Context, {{.lowerStartCamelPrimaryKey}} {{.dataType}}) (rowsAffected int64,err error) {
{{if .withCache}}{{if .containsIndexCache}}
    data, err = m.FindOne(ctx, {{.lowerStartCamelPrimaryKey}})
	if err!=nil{
		return
	}
{{end}}	{{.keys}}
	 err {{if .containsIndexCache}}={{else}}={{end}} m.ExecCtx(ctx, func(conn *gorm.DB) error {
		return conn.WithContext(ctx).Delete(&{{.upperStartCamelObject}}{}, {{.lowerStartCamelPrimaryKey}})
	}, {{.keyValues}}){{else}}
	    _,_, err = m.FindOne(ctx, {{.lowerStartCamelPrimaryKey}})
        if err != nil {
            return
        }
        result := m.conn.WithContext(ctx).Table(m.tableName()).Delete(&{{.upperStartCamelObject}}{},{{.lowerStartCamelPrimaryKey}})
	{{end}}
	return result.RowsAffected, result.Error
}