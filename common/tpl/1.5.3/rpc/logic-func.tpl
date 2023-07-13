{{if .hasComment}}{{.comment}}{{end}}
func (l *{{.logicName}}) {{.method}} ({{if .hasReq}}in {{.request}}{{if .stream}},stream {{.streamBody}}{{end}}{{else}}stream {{.streamBody}}{{end}}) ({{if .hasReply}} resp {{.response}},{{end}} err error) {
	// todo: add your logic here and delete this line
	l.Infow("{{.method}} start", logx.Field("in", in))
	{{if .hasReply}}resp = &{{.responseType}}{}{{end}}

	return {{if .hasReply}}resp,{{end}} nil
}
