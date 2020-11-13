package urlshort

import (
	"net/http"
)

// MapHandler function
func MapHandler(pathToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if dest, ok := pathToUrls[path]; ok {
			http.Redirect(w, r, dest, http.StatusFound)
		}
		//if we can match a path, redirect to it

		//else
		fallback.ServeHTTP(w, r)
	}
}
