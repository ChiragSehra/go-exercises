package urlshort

import (
	"net/http"

	"gopkg.in/yaml.v2"
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

// YAMLHandler function
func YAMLHandler(yamlBytes []byte, fallback http.Handler) (http.HandlerFunc, error) {
	//1. Parse the yaml somehow
	var pathUrls []pathUrl
	err := yaml.Unmarshal(yamlBytes, &pathUrls)
	if err != nil {
		return nil, err
	}
	//2. Convert YAML to map
	pathToUrls := make(map[string]string)
	for _, pu := range pathUrls {
		pathToUrls[pu.Path] = pu.URL
	}
	//3. Return a mapHandler using the map
	return MapHandler(pathToUrls, fallback), nil

}

type pathUrl struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}
