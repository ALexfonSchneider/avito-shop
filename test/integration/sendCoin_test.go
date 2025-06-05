//go:build integrational
// +build integrational

package integration

import (
	"bytes"
	"encoding/json"
	authhandler "github.com/ALexfonSchneider/avito-shop/internal/adapter/http/handler/auth/authenticate"
	sendcoinsnhandler "github.com/ALexfonSchneider/avito-shop/internal/adapter/http/handler/transaction/sendCoins"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_SendCoin(t *testing.T) {
	reqAuthUser1 := authhandler.Request{
		Username: "TestUser1",
		Password: "password",
	}
	body, err := json.Marshal(reqAuthUser1)
	if err != nil {
		panic(err)
	}

	req := httptest.NewRequest(http.MethodPost, "/api/auth", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	webApp.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Errorf("Test_SendCoin: expected status 200 got %d", rec.Code)
	}

	var user1Token authhandler.Response
	err = json.Unmarshal(rec.Body.Bytes(), &user1Token)
	if err != nil {
		panic(err)
	}

	reqAuthUser2 := authhandler.Request{
		Username: "TestUser2",
		Password: "password",
	}
	body, err = json.Marshal(reqAuthUser2)
	if err != nil {
		panic(err)
	}

	req = httptest.NewRequest(http.MethodPost, "/api/auth", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	rec = httptest.NewRecorder()

	webApp.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("Test_SendCoin: expected status 200 got %d", rec.Code)
	}

	reqSendCoins := sendcoinsnhandler.Request{
		ToUser: reqAuthUser2.Username,
		Amount: 500,
	}
	body, err = json.Marshal(reqSendCoins)
	if err != nil {
		panic(err)
	}

	req = httptest.NewRequest(http.MethodPost, "/api/sendCoin", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", user1Token.Token)

	rec = httptest.NewRecorder()

	webApp.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Errorf("Test_SendCoin: expected status 200 got %d", rec.Code)
	}
}
