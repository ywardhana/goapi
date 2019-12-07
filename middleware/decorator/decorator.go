package decorator

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/ywardhana/goapi/middleware/authenticator"
)

type HandleWithError func(http.ResponseWriter, *http.Request, httprouter.Params) error

type Decorator struct {
	afterSuccessHandler http.Handler
	failedAuthHandler   http.Handler
	failedHandler       http.Handler
}

func NewDecorator(
	afterSuccessHandler http.Handler,
	failedAuthHandler http.Handler,
	failedHandler http.Handler,
) *Decorator {
	return &Decorator{
		afterSuccessHandler: afterSuccessHandler,
		failedAuthHandler:   failedAuthHandler,
		failedHandler:       failedHandler,
	}
}

func (d *Decorator) ApplyDecorator(handler HandleWithError,
	auth authenticator.Auth) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		if !d.evaluate(auth, r) {
			d.failedAuthHandler.ServeHTTP(w, r)
			return
		}
		if err := handler(w, r, params); err != nil {
			d.failedHandler.ServeHTTP(w, r)
			return
		}
		d.afterSuccessHandler.ServeHTTP(w, r)
	}
}

func (d *Decorator) evaluate(auth authenticator.Auth, r *http.Request) bool {
	return auth.Authenticate(r)
}
