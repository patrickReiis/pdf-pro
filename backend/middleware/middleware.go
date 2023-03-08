// This package is responsible for any kind of middleware.
// That is, everything not related to the actual service that needs to be done before using the actual service.
package middleware

import "net/http"

// Calls all the midlewares
// If the request passed all the middleware functions the `originalRoute` will be called
// Be aware that this function is a Closure.
func AllMiddleware(originalRoute http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		// If the request doesn't pass through any middleware, the http.ResponseWriter `w` will be modified by writing to it a response.
		// If `ok` is false it means the http.ResponseWriter `w` WAS modified, therefore should return
		// If `ok` is true, it means the http.ResponseWriter was NOT modified, therefore proceed the execution
		ok := RouteWithAuth(w, r)
		if ok == false {
			return
		}

		originalRoute(w, r)
	}

}
