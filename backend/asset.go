package backend

import "net/http"

const LAssetVideoRoute = "/LAssetVideoRead.mp4"

func LAssetMiddlewareCreate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		if request.URL.Path == LAssetVideoRoute {
			LAssetVideoRead(response, request)
			return
		}

		next.ServeHTTP(response, request)
	})
}

func LAssetVideoRead(response http.ResponseWriter, request *http.Request) {
	if request.URL.Path != LAssetVideoRoute {
		http.NotFound(response, request)
		return
	}

	id := request.URL.Query().Get("asset")
	if id == "" {
		http.Error(response, "missing video asset", http.StatusBadRequest)
		return
	}

	file, info, err := LAssetFileOpen(request.Context(), id, request.URL.Query().Get("preview") == "compatibility", request.URL.Query().Get("session"))
	if err != nil {
		http.NotFound(response, request)
		return
	}
	defer file.Close()

	response.Header().Set("Cache-Control", "no-store")
	response.Header().Set("Accept-Ranges", "bytes")
	http.ServeContent(response, request, info.Name(), info.ModTime(), file)
}
