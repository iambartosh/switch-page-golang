package main 

import (
	"net/http"
	"template/html"
	"log"
)

/* or 
	* func() HTMLRender(page, w http.ResponseWriter, r *http.Request)
*/
func (page string) HTMLRender(w http.ResponseWriter, r *http.Request) string {
	temp, err := template.ParseFile(page + ".html")
	if err != nil {
		http.NotFound(w, r)
	}
	temp.ExecuteTemplate(w, page)
}

func main() {
	route := mux.NewRouter()

	route.Handle("/", HTMLRender("index"))
	route.Handle("/login", HTMLRender("login"))

	log.Fatal(http.ListenAndServe(":8000", route))
}
