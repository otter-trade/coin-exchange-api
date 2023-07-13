// Copyright 2023 The Ryan SU Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package i18n

import (
	"context"
	"embed"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/text/language"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//go:embed locale/*.json
var LocaleFS embed.FS

// Translator is a struct storing translating data.
type Translator struct {
	bundle       *i18n.Bundle
	localizer    map[language.Tag]*i18n.Localizer
	supportLangs []language.Tag
}

// NewBundle returns a bundle from FS.
func (l *Translator) NewBundle(file embed.FS) {
	bundle := i18n.NewBundle(language.Chinese)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	_, err := bundle.LoadMessageFileFS(file, "locale/zh.json")
	logx.Must(err)
	_, err = bundle.LoadMessageFileFS(file, "locale/en.json")
	logx.Must(err)

	l.bundle = bundle
}

// NewTranslator sets localize for translator.
func (l *Translator) NewTranslator() {
	l.supportLangs = append(l.supportLangs, language.Chinese)
	l.supportLangs = append(l.supportLangs, language.English)
	l.localizer = make(map[language.Tag]*i18n.Localizer)
	l.localizer[language.Chinese] = i18n.NewLocalizer(l.bundle, language.Chinese.String())
	l.localizer[language.English] = i18n.NewLocalizer(l.bundle, language.English.String())
}

// Trans used to translate any i18n string.
func (l *Translator) Trans(ctx context.Context, msgId string, lang string) string {
	message, err := l.MatchLocalizer(lang).LocalizeMessage(&i18n.Message{ID: msgId})
	if err != nil || message == "" {
		return ""
	}

	return message
}

// TransError translates the error message
func (l *Translator) TransError(ctx context.Context, err error) error {
	lang := ctx.Value("lang").(string)
	if IsGrpcError(err) {
		message, e := l.MatchLocalizer(lang).LocalizeMessage(&i18n.Message{ID: strings.Split(err.Error(), "desc = ")[1]})
		if e != nil || message == "" {
			message = err.Error()
		}
		return status.Error(status.Code(err), message)
	} else if codeErr, ok := err.(*CodeError); ok {
		message, e := l.MatchLocalizer(lang).LocalizeMessage(&i18n.Message{ID: codeErr.Error()})
		if e != nil || message == "" {
			message = codeErr.Error()
		}
		return NewCodeErrorWithMsg(codeErr.Code, message)
	} else if apiErr, ok := err.(*ApiError); ok {
		message, e := l.MatchLocalizer(lang).LocalizeMessage(&i18n.Message{ID: apiErr.Error()})
		if e != nil {
			message = apiErr.Error()
		}
		return NewApiErrorWithMsg(apiErr.Code, message)
	} else {
		return NewApiErrorWithMsg(http.StatusInternalServerError, "failed to translate error message")
	}
}

// MatchLocalizer used to matcher the localizer in map
func (l *Translator) MatchLocalizer(lang string) *i18n.Localizer {
	tags := ParseTags(lang)
	for _, v := range tags {
		if val, ok := l.localizer[v]; ok {
			return val
		}
	}

	return l.localizer[language.Chinese]
}

// NewTranslator returns a translator by FS.
func NewTranslator(file embed.FS) *Translator {
	trans := &Translator{}
	trans.NewBundle(file)
	trans.NewTranslator()
	return trans
}

func CodeFromGrpcError(err error) int {
	code := status.Code(err)
	switch code {
	case codes.OK:
		return http.StatusOK
	case codes.InvalidArgument, codes.FailedPrecondition, codes.OutOfRange:
		return http.StatusBadRequest
	case codes.Unauthenticated:
		return http.StatusUnauthorized
	case codes.PermissionDenied:
		return http.StatusForbidden
	case codes.NotFound:
		return http.StatusNotFound
	case codes.Canceled:
		return http.StatusRequestTimeout
	case codes.AlreadyExists, codes.Aborted:
		return http.StatusConflict
	case codes.ResourceExhausted:
		return http.StatusTooManyRequests
	case codes.Internal, codes.DataLoss, codes.Unknown:
		return http.StatusInternalServerError
	case codes.Unimplemented:
		return http.StatusNotImplemented
	case codes.Unavailable:
		return http.StatusServiceUnavailable
	case codes.DeadlineExceeded:
		return http.StatusGatewayTimeout
	}

	return http.StatusInternalServerError
}

// IsGrpcError checks if the error is a gRPC error.
func IsGrpcError(err error) bool {
	if err == nil {
		return false
	}

	_, ok := err.(interface {
		GRPCStatus() *status.Status
	})

	return ok
}
