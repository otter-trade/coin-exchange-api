
func new{{.upperStartCamelObject}}Model(conn *gorm.DB{{if .withCache}}, c cache.CacheConf{{end}}) *default{{.upperStartCamelObject}}Model {
	return &default{{.upperStartCamelObject}}Model{
		{{if .withCache}}CachedConn: gormc.NewConn(conn, c){{else}}conn:conn{{end}},
		table: {{.table}},
	}
}

func (m *default{{.upperStartCamelObject}}Model) BeforeCreate(tx *gorm.DB) error {
	tx.Statement.SetColumn("create_time", time.Now().Unix())
	tx.Statement.SetColumn("update_time", time.Now().Unix())
	return nil
}

func (m *default{{.upperStartCamelObject}}Model) BeforeUpdate(tx *gorm.DB) error {
	tx.Statement.SetColumn("update_time", time.Now().Unix())
	return nil
}

func (u *default{{.upperStartCamelObject}}Model) BeforeDelete(tx *gorm.DB) error {
    tx.Statement.SetColumn("delete_time", time.Now().Unix())
    return nil
}