package middleware

import "net/http"

type Auth interface {
	Authenticate(r *http.Request) bool
}
