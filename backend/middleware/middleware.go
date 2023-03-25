// This package is responsible for any kind of middleware.
// That is, everything not related to the actual service that needs to be done before using the actual service.
package middleware

import "net/http"

type middlewareRoute (func(http.ResponseWriter, *http.Request) (ok bool))
type MiddlewareRoutes []middlewareRoute

// Calls all the midlewares
// If the request passed all the middleware functions the `originalRoute` will be called
// Be aware that this function is a Closure.
func AllMiddleware(originalRoute http.HandlerFunc, middlewares MiddlewareRoutes) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		for _, middleware := range middlewares {
			// If `ok` is false it means the http.ResponseWriter `w` WAS modified, therefore should return
			// If `ok` is true, it means the http.ResponseWriter was NOT modified, therefore proceed the execution
			ok := middleware(w, r)

			if ok == false {
				return
			}
		}
		originalRoute(w, r)
	}
}
