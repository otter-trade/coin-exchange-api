package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/otter-trade/coin-exchange-api/common/i18n"
	"github.com/otter-trade/coin-exchange-api/common/xresp"
	go_sign "github.com/parkingwang/go-sign"
	"github.com/zeromicro/go-zero/core/logx"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const (
	// 默认签名密钥
	secret = "shoope@2023"
)

type Sign struct {
	secret  string
	timeout int64
	mode    string
}

type SignOption struct {
	Secret  string
	Timeout int64
	Mode    string
}

// 检查签名中间件，可自定义密钥和超时时间
func NewSign(option *SignOption) *Sign {
	mw := &Sign{secret: secret, timeout: 10}
	if option != nil {
		if option.Secret != "" {
			mw.secret = option.Secret
		}
		if option.Timeout > 0 {
			mw.timeout = option.Timeout
		}
		if option.Mode != "" {
			mw.mode = option.Mode
		}
	}
	return mw
}

func (m *Sign) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if m.mode == "dev" {
			next(w, r)
			return
		}
		sign := r.Header.Get("Sign")
		// 处理json格式传参
		if (r.Method == http.MethodPost || r.Method == http.MethodDelete || r.Method == http.MethodPut) && strings.HasPrefix(r.Header.Get("Content-Type"), "application/json") {
			params := make(map[string]interface{})
			body, err := io.ReadAll(r.Body)
			r.Body = io.NopCloser(bytes.NewBuffer(body))
			if err != nil {
				xresp.Fail(r, w, i18n.NewApiError(i18n.SignError))
				return
			}
			decoder := json.NewDecoder(strings.NewReader(string(body)))
			decoder.UseNumber()
			err = decoder.Decode(&params)
			if err != nil {
				xresp.Fail(r, w, i18n.NewApiError(i18n.SignError))
				return
			}
			logx.Infof("check sign json get params: %v", params)

			if sign == "" {
				xresp.Fail(r, w, i18n.NewApiError(i18n.SignError))
				return
			}
			delete(params, "sign")

			signer := RecursionHandle(params)
			signer.SetSignBodySuffix(m.secret)

			tempStr := signer.GetSignBodyString()
			logx.Infof("sign string: %v", tempStr)

			logx.Infof("sign: %v", signer.GetSignature())

			if sign != signer.GetSignature() {
				xresp.Fail(r, w, i18n.NewApiError(i18n.SignError))
				return
			}

			next(w, r)
			return
		} else {
			signer := go_sign.NewGoSignerMd5()
			signer.SetSignBodySuffix(m.secret)
			// 处理表单和url参数
			var params url.Values
			r.ParseForm()
			if r.Method == http.MethodPost {
				params = r.PostForm
			} else if r.Method == http.MethodGet {
				params = r.Form
			}

			if sign == "" {
				xresp.Fail(r, w, i18n.NewApiError(i18n.SignError))
				return
			}
			params.Del("sign")

			for k, v := range params {
				signer.AddBodies(k, v)
			}

			if sign != signer.GetSignature() {
				xresp.Fail(r, w, i18n.NewApiError(i18n.SignError))
				return
			}

			next(w, r)
		}
	}
}

func RecursionHandle(params map[string]interface{}) (signer *go_sign.GoSigner) {
	signer = go_sign.NewGoSignerMd5()
	for k, v := range params {
		switch v.(type) {
		case []interface{}:
			op := RecursionMapHandle(v.([]interface{}))
			signer.AddBodies(k, []string{fmt.Sprintf("%v", strings.Join(op, ","))})
		case map[string]interface{}:
			temp := RecursionHandle(v.(map[string]interface{})).GetSignBodyString()
			signer.AddBodies(k, []string{fmt.Sprintf("%v", temp)})
		default:
			signer.AddBodies(k, []string{fmt.Sprintf("%v", v)})
		}
	}
	return
}

func RecursionMapHandle(params []interface{}) []string {
	op := make([]string, 0)
	for _, item := range params {
		switch item.(type) {
		case map[string]interface{}:
			temp := RecursionHandle(item.(map[string]interface{})).GetSignBodyString()
			op = append(op, fmt.Sprintf("%v", temp))
		case []interface{}:
			temp := RecursionMapHandle(item.([]interface{}))
			op = append(op, temp...)
		default:
			op = append(op, fmt.Sprintf("%v", item))
		}
	}
	return op
}
