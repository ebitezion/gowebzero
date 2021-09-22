package route

import "net/http"

func RunTrad() {
	http.ListenAndServe(":8080", http.FileServer(http.Dir("C:/Users/P. Zion/go/src/gowebzero/route/site")))
}
