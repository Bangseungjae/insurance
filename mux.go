package main

import (
	"Bangseungjae/insurance/handler"
	"Bangseungjae/insurance/service"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func NewMux() (http.Handler, error) {
	mux := chi.NewRouter()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		// 정적 분석 오류를 회피하기 위해 명시적으로 반환값을 버린다.
		_, _ = w.Write([]byte(`{"status": "ok"}`))
	})
	li := &handler.GetUserInsurance{Service: &service.ListUserInsurance{}}
	mux.Get("/insurance", li.ServeHTTP)
	return mux, nil
}
