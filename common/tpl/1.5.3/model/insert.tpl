
func (m *default{{.upperStartCamelObject}}Model) Insert(ctx context.Context, data *{{.upperStartCamelObject}}) (err error) {
{{if .withCache}}{{.keys}}
    err = m.ExecCtx(ctx, func(conn *gorm.DB) error {
		return conn.WithContext(ctx).Table(m.tableName()).Save(&data).Error
	}, {{.keyValues}})
{{else}}
	err=m.conn.WithContext(ctx).Table(m.tableName()).Save(&data).Error
{{end}}
	return
}