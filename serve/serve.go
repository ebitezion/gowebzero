package serve

import (
	"fmt"
	"net/http"
	"time"
)

//using fmt
func serveDynamic(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Time is", time.Now().String())
}

//serving static files
func serveStatic(w http.ResponseWriter, r *http.Request) {

	http.ServeFile(w, r, "C:/Users/P. Zion/go/src/gowebzero/serve/site/qrc.png")
}

//Run func to be executed in main
func Run() {
	http.HandleFunc("/dynamic", serveDynamic)
	http.HandleFunc("/static", serveStatic)

	http.ListenAndServe(":8080", nil)
}
