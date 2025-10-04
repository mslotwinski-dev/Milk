package web

import (
	"fmt"
	"milk/core"
	"net/http"
)

func Run(page core.Renderable, port int) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprint(w, page.Render())
	})

	fmt.Printf("Server running at http://localhost:%d/\n", port)
	if err := http.ListenAndServe(fmt.Sprintf("localhost:%d", port), nil); err != nil {
		fmt.Println("Error:", err)
	}
}
