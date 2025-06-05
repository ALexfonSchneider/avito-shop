//go:build integrational
// +build integrational

package integration

import (
	"bytes"
	"encoding/json"
	"fmt"
	authhandler "github.com/ALexfonSchneider/avito-shop/internal/adapter/http/handler/auth/authenticate"
	"github.com/labstack/echo/v4"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_BuyMerch(t *testing.T) {
	authRequest, err := json.Marshal(authhandler.Request{
		Username: "alex",
		Password: "password",
	})
	if err != nil {
		panic(err)
	}

	req := httptest.NewRequest(http.MethodPost, "/api/auth", bytes.NewBuffer(authRequest))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	if webApp.ServeHTTP(rec, req); rec.Code != http.StatusOK {
		panic(fmt.Sprintf("http status code: %d", rec.Code))
	}

	var authResp authhandler.Response
	err = json.Unmarshal(rec.Body.Bytes(), &authResp)
	if err != nil {
		panic(err)
	}

	req = httptest.NewRequest(http.MethodGet, "/api/buy/t-shirt", nil)
	req.Header.Set(echo.HeaderAuthorization, authResp.Token)

	if webApp.ServeHTTP(rec, req); rec.Code != http.StatusOK {
		panic(rec.Body.String())
	}
}
