package xresp

import (
	"context"
	"github.com/go-playground/validator/v10"
	en_translation "github.com/go-playground/validator/v10/translations/en"
	zh_translation "github.com/go-playground/validator/v10/translations/zh"
	"github.com/otter-trade/coin-exchange-api/common/i18n"
	"go.opentelemetry.io/otel/trace"
	"net/http"

	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/ko"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/rest/httpx"
	"google.golang.org/grpc/status"
)

var (
	Validate *validator.Validate
	uni      *ut.UniversalTranslator
	trans    ut.Translator
)

func init() {
	Validate = validator.New()
}

type Body struct {
	Code int64       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

// mis返回
type PaginationInfo struct {
	PageNo   int         `json:"page_no,omitempty"`
	PageSize int         `json:"page_size,omitempty"`
	Total    int         `json:"total"`
	List     interface{} `json:"list"`
}

func Fail(r *http.Request, w http.ResponseWriter, err error) {
	Success(r, w, nil, err)
}

// handler模板会调用该函数
func Success(r *http.Request, w http.ResponseWriter, resp interface{}, err error) {
	var body Body
	language := r.Header.Get("language")
	translator := i18n.NewTranslator(i18n.LocaleFS)
	if err == nil {
		// 若没有错误产生
		body.Code = i18n.Success
		body.Msg = translator.Trans(r.Context(), i18n.CodeToMsg(i18n.Success), language)
		body.Data = resp
	} else {
		// 若有错误
		causeErr := errors.Cause(err)
		switch v := causeErr.(type) {
		case *i18n.CodeError:
			body.Code = v.GetCode()
			body.Msg = translator.Trans(r.Context(), i18n.CodeToMsg(v.GetCode()), language)
		case *i18n.ApiError:
			body.Code = v.GetCode()
			body.Msg = translator.Trans(r.Context(), i18n.CodeToMsg(v.GetCode()), language)
		case validator.ValidationErrors:
			// 参数校验错误
			zhT := zh.New() //chinese
			enT := en.New() //english
			koT := ko.New() //Korean
			uni := ut.New(enT, zhT, koT)
			trans, _ = uni.GetTranslator("zh")
			switch language {
			case "en":
				err = en_translation.RegisterDefaultTranslations(Validate, trans)
			case "zh":
				err = zh_translation.RegisterDefaultTranslations(Validate, trans)
			case "ko":
				//err = ko_translation.RegisterDefaultTranslations(Validate, trans)
			default:
				err = zh_translation.RegisterDefaultTranslations(Validate, trans)
			}

			body.Code = i18n.ParamError
			body.Msg = v[0].Translate(trans)
		default:
			gstatus, ok := status.FromError(v)

			message := gstatus.Message()
			if gstatus.Code() != 0 && gstatus.Message() == "" {
				message = translator.Trans(r.Context(), i18n.CodeToMsg(int64(gstatus.Code())), language)
			}

			if ok { // grpc错误
				body.Code = int64(gstatus.Code())
				body.Msg = message
			} else { // 未知错误
				body.Code = i18n.ServerError
				body.Msg = "unknown"
			}
		}
	}

	w.Header().Set("trace-id", TraceIdFromContext(r.Context()))

	httpx.OkJson(w, body)
}

func TraceIdFromContext(ctx context.Context) (traceId string) {
	spanCtx := trace.SpanContextFromContext(ctx)
	if spanCtx.HasTraceID() {
		return spanCtx.TraceID().String()
	}
	return
}
