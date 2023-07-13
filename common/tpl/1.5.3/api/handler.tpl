package {{.PkgName}}

import (
    "github.com/otter-trade/coin-exchange-api/common/xresp"
	"net/http"
	"github.com/otter-trade/coin-exchange-api/common/i18n"
	"github.com/zeromicro/go-zero/rest/httpx"
	{{.ImportPackages}}
)

func {{.HandlerName}}(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		{{if .HasRequest}}var req types.{{.RequestType}}

		if err := httpx.Parse(r, &req); err != nil {
            xresp.Fail(r, w, i18n.NewCodeError(i18n.ParseParamsError))
            return
        }

		if err := xresp.Validate.StructCtx(r.Context(), req); err != nil {
            xresp.Fail(r, w, err)
            return
        }



		{{end}}l := {{.LogicName}}.New{{.LogicType}}(r, svcCtx)
		{{if .HasResp}}resp, {{end}}err := l.{{.Call}}({{if .HasRequest}}&req{{end}})
		{{if .HasResp}}xresp.Success(r, w, resp, err){{else}}xresp.Success(r, w, nil, err){{end}}
	}
}
