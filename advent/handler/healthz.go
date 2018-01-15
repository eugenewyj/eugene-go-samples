package handler

import "net/http"

// healthz 是一个程序运行的探针
func healthz(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}
