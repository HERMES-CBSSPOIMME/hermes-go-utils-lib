package customhandle

import "net/http"

type (
	Handler func(w http.ResponseWriter, r *http.Request) error
)

// CustomHandle : Custom Handlers Wrapper for API
func CustomHandle(handlers ...Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for _, h := range handlers {
			err := h(w, r)
			if err != nil {
				// w.Write(getResponseOfError(err))
				return
			}
		}
	})
}
