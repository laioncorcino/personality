package middleware

import "net/http"

func ContentTypeMid(next http.Handler) http.Handler {
	return http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		resp.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(resp, req)
	})
}
