package {{.pkgName}}

import (
    "context"
    "net/http"
	{{.imports}}
)

type {{.logic}} struct {
	logx.Logger
	r      *http.Request
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func New{{.logic}}(r *http.Request, svcCtx *svc.ServiceContext) *{{.logic}} {
	return &{{.logic}}{
		Logger: logx.WithContext(r.Context()),
		r:    r,
		ctx:    r.Context(),
		svcCtx: svcCtx,
	}
}

func (l *{{.logic}}) {{.function}}({{.request}}) {{.responseType}} {
	// todo: add your logic here and delete this line
	l.Infow("{{.function}} start", logx.Field("req", req))


	{{.returnString}}
}
