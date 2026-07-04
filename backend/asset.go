package backend

import (
	"net/http"
	"os"
)

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

	path := request.URL.Query().Get("path")
	if path == "" {
		http.Error(response, "missing video path", http.StatusBadRequest)
		return
	}

	info, err := os.Stat(path)
	if err != nil || info.IsDir() {
		http.NotFound(response, request)
		return
	}

	servedPath, err := LAssetPreviewResolve(path, info)
	if err != nil {
		servedPath = path
	}

	response.Header().Set("Cache-Control", "no-store")
	response.Header().Set("Accept-Ranges", "bytes")
	http.ServeFile(response, request, servedPath)
}
