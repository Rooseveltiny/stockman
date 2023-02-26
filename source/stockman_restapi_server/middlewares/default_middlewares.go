package middlewares

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

/* middleware performers */
func EmptyMiddlewareFunc(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Println("some prints")
}

/* middleware instances */
var EmptyMiddleware = NewMiddleware(EmptyMiddlewareFunc)

/* middleware sets */
var DefaultMiddelwareSet = NewMiddlewareSet(EmptyMiddleware)
