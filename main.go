package main

import (
	"net/http"
	"fmt"
)

func main(){
	http.HandleFunc("/", handler)
	http.HandleFunc("/home", uploadRoot)
	http.HandleFunc("/evaluate",evaluatehandler)
	http.ListenAndServe(":8080", nil)
}
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func uploadRoot(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, `
	<form action="/evaluate" method="POST" enctype='multipart/form-data'>
	primer <input type="file" name="f1"><br>
	segundo <input type="file" name="f2><br>
	<input type="submit" value="Carga">
	</form>
	`)
}
func evaluatehandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "uploaded =D")
}