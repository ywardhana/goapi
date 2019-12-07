package authenticator

import "net/http"

type Auth interface {
	Authenticate(r *http.Request) bool
}
