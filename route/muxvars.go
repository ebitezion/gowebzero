package route

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func page(w http.ResponseWriter, r *http.Request) {
	pageID := mux.Vars(r)["id"]
	filename := "C:/Users/P. Zion/go/src/gowebzero/route/site/" + pageID + ".html"
	if _, err := os.Stat(filename); err != nil {
		filename = "C:/Users/P. Zion/go/src/gowebzero/route/site/" + "404.html"
	}
	http.ServeFile(w, r, filename)

}

func RunMuxVars() {
	route := mux.NewRouter()

	route.HandleFunc("/page/{id:[0-9]+}", page)
	route.HandleFunc("/", page)
	http.Handle("/", route)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
