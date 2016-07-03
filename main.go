package main

import (
	"net/http"
	"fmt"
	"time"
	"os"
	"io"
	"os/exec"
	"strings"
)

func main(){
	//http.HandleFunc("/", handler)
	http.HandleFunc("/home", uploadRoot)
	http.HandleFunc("/evaluate",evaluatehandler)
	http.ListenAndServe(":8080", nil)
}
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func uploadRoot(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, `<html><body>
	<form action="/evaluate" method="POST" enctype='multipart/form-data'>
	primer <input type="file" name="f1"><br>
	segundo <input type="file" name="f2"><br>
	<input type="submit" value="Carga">
	</form></body></html>
	`)
}
func evaluatehandler(w http.ResponseWriter, r *http.Request){

	r.ParseMultipartForm(32 << 20)

	f1,err := parseFile(r,"f1")
	if err != nil{
		fmt.Fprintf(w,"ERROR: %v",err)
		return
	}
	f2,err := parseFile(r,"f2")
	if err != nil{
		fmt.Fprintf(w,"ERROR: %v",err)
		return
	}

	fmt.Printf("Archivos '%s', '%s'",f1,f2)

	out, err := exec.Command("br","-algorithm","FaceRecognition","-compare",f1,f2).Output()
	if err != nil {
		fmt.Fprintf(w,"ERROR: %v",err)
		return
	}
	out = strings.Trim(out, " \n\t")
	fmt.Fprintf(w,"OUTPUT '%s'\n", out)
	os.Remove(f1)
	os.Remove(f2)
}

func parseFile(r *http.Request, filevar string)(string,error){
	file, handler, err := r.FormFile(filevar)
	if err != nil {
		return "",err
	}
	defer file.Close()
	fmt.Printf("%v", handler.Header)
	crutime := time.Now().Unix()
	filename := fmt.Sprintf("/tmp/downloads/%d_%s",crutime,handler.Filename)
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return "",err
	}
	defer f.Close()
	io.Copy(f, file)

	return filename,nil
}