package main 

import (
	"net/http"
	"template/html"
	"log"	
)

func render(page string) string {
	temp, err = template.ParseFile(page + ".html")
	if err != nil {
		log.Fatal("The template doesn't existing...")
	}
	temp.ExecuteTemplate(w, page, data)
}

func main() {
	router := httprouter.New()
	router.GET("/", func(w http.ResponseWriter, req *http.Request) {
		/* checking the current paths */
		path := req.URL.path

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
			render("index")
			break;			
		}
	})
	log.Fatal(http.ListenAndServe(":8000", router))
}
