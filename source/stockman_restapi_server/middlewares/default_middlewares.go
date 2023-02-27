package middlewares

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

/* middleware performers */
func LogInputRequest(w http.ResponseWriter, r *http.Request, p httprouter.Params) {}

/* middleware instances */
var LogInputRequestMiddleware = NewMiddleware(LogInputRequest)

/* middleware sets */
var DefaultMiddelwareSet = NewMiddlewareSet(LogInputRequestMiddleware)
