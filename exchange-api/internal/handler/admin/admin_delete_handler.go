package admin

import (
	"github.com/otter-trade/coin-exchange-api/common/i18n"
	"github.com/otter-trade/coin-exchange-api/common/xresp"
	"github.com/otter-trade/coin-exchange-api/exchange-api/internal/logic/admin"
	"github.com/otter-trade/coin-exchange-api/exchange-api/internal/svc"
	"github.com/otter-trade/coin-exchange-api/exchange-api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func AdminDeleteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AdminDeleteReq

		if err := httpx.Parse(r, &req); err != nil {
			xresp.Fail(r, w, i18n.NewCodeError(i18n.ParseParamsError))
			return
		}

		if err := xresp.Validate.StructCtx(r.Context(), req); err != nil {
			xresp.Fail(r, w, err)
			return
		}

		l := admin.NewAdminDeleteLogic(r, svcCtx)
		resp, err := l.AdminDelete(&req)
		xresp.Success(r, w, resp, err)
	}
}
