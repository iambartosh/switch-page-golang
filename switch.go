package main 

import (
	"net/http"
	"template/html"
	"log"	
)

func render(page) // e.g func render("index") {
	temp, err = template.ParseFile(page + ".html")
	
	if err {
		log.Fatal("The template doesn't existing...")
	}
	temp.ExecuteTemplate(w, page, data)
}

func main() {
	router := httprouter.New()
	router.GET("/", func(w http.ResponseWriter, req *http.Request) {
		/* checking the currently paths */
		path := req.URL.path
		/* switching path */

		switch path {
		case "/hello":
			render("hello")
			break;
		case "/bye":
			render("bye")
			break;
		case "/":
			render("index")
			break;
		default:
			// if path is another than other of paths rendering a index page
			render("index")
			break;			
		}
	})
	log.Fatal(http.ListenAndServe(":8000", router))
}
