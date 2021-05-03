package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)
type article struct {
	ID      int
	Title   string
	Content string
}
var data = []article{
	article{1, "lorem", "lorem"},
	article{2, "ipsum", "ipsum"},
}
func articles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		var result, err = json.Marshal(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(result)
		return
	}
	http.Error(w, "", http.StatusBadRequest)
}

func main() {
  http.HandleFunc("/articles", articles)
  fmt.Println("starting web server at http://localhost:8080/")
  http.ListenAndServe(":8080", nil)
}


