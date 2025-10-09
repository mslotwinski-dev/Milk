package web

import (
	"fmt"
	"milk/core"
	"milk/core/utility"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
)

func detectCallerDir() string {
	_, file, _, ok := runtime.Caller(2)
	if !ok {
		utility.Warn("Caller directory not found, defaulting to current working directory")
		wd, _ := os.Getwd()
		return wd
	}
	return filepath.Dir(file)
}

func Run(page core.Renderable, port int) {
	rootFolder := detectCallerDir()

	assetsPath := filepath.Join(rootFolder, "assets")
	publicPath := filepath.Join(rootFolder, "public")

	if _, err := os.Stat(assetsPath); os.IsNotExist(err) {
		utility.Warn("Assets folder not found in %s", assetsPath)
	}
	if _, err := os.Stat(publicPath); os.IsNotExist(err) {
		utility.Warn("Public folder not found in %s", publicPath)
	}

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir(assetsPath))))
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir(publicPath))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprint(w, page.Render())
	})

	utility.Info("Server running at http://localhost:%d/\n", port)
	if err := http.ListenAndServe(fmt.Sprintf("localhost:%d", port), nil); err != nil {
		utility.Error("Failed to start server: %v", err)
	}
}
