package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/mux"
		//"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
	"mime/multipart"
	"github.houston.softwaregrp.net/CSB/fileserver/pkg/fileserver"
	"io"
)

type Init struct {
	Dir string
}

func main() {
	var dir string
	var port string
	flag.StringVar(&dir, "dir", "Public", "the directory to serve files from. Defaults to the current dir")
	flag.StringVar(&port, "port", ":9993", "The Port the server is listening on")
	flag.Parse()
	i := Init{dir}
	p := &i
	fmt.Print(p)
	r := mux.NewRouter()

	r.HandleFunc("/api/upload/", i.HandleUpload)
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer("./"+http.Dir(dir))))
	r.HandleFunc("/", i.HandleRoot)

	srv := &http.Server{
		Handler:      r,
		Addr:         port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}

func (i *Init) HandleRoot(w http.ResponseWriter, r *http.Request) {

	path := fmt.Sprint("C:/Users/smithay/Documents/godev/src/github.houston.softwaregrp.net/CSB/fileserver/web/fileserver/dist/fileserver/index.html")
	file, err := os.Open(path)
	if err != nil {
		fmt.Print(err, "\n 1 \n")
	}
	fileContents, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Print(err, "\n 2 \n")
	}
	w.Header().Set("Content-Type", "text/html")
	w.Write(fileContents)
	}

func (i *Init)HandleStatic(w http.ResponseWriter,r *http.Request) {
	fmt.Print("dir",i.Dir)
	bob := RipFunc(i.Dir)
	rob:=fmt.Sprintf("%s%s%s","<ul>",bob,"</ul>")
	fmt.Print(rob)

	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(rob))
	fmt.Print(rob)
}
func (i *Init) HandleUpload(w http.ResponseWriter, r *http.Request){ //string,string,*multipart.FileHeader,string) {
if r.Method=="POST" {
	fmt.Print("\n here2 \n")
	patho, _ := os.Getwd()
	l, handler := PullUploader(r, w)
	pather := string(patho + "/" + handler.Filename)
	i.Move(pather, patho, handler, l)

}else{
	fmt.Fprint(w,"405 Method not allowed")
}

}
func PullUploader(r *http.Request,w http.ResponseWriter)(string,*multipart.FileHeader) {

	file, header, err := r.FormFile("uploadfile")
	if err != nil {
		fmt.Fprintln(w, err)

	}
	defer file.Close()

	out, err := os.Create(header.Filename)
	if err != nil {
		fmt.Print(err)
}

	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		fmt.Fprintln(w, err)
	}

	l := r.FormValue("filepath")
	fmt.Print("\n here5 \n")
	fmt.Fprint(w,"200 File successfully uploaded")
	return l,header
}

func RipFunc (root string) (fum string) {
	fi, err := ioutil.ReadDir(root)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("num of file descriptor things: %d\n", len(fi))
	for _,fo := range fi {

		currentFile := fileserver.FileFilter(fo)
		if currentFile != "" {
			fum = fmt.Sprintf("%s%s", fum, currentFile)
		}
	}

	return
}


func (i *Init)Move(pather string ,patho string, handler *multipart.FileHeader,l string) {

	if l == pather {
		npat := string("./"+l+handler.Filename)
		fmt.Print("\n npat2:", npat)
		err := os.Rename(pather, npat)
		if err != nil {
			fmt.Print(err)
		}
	}else{
		os.MkdirAll(l,0666)
		npat := string("./"+l+handler.Filename)
		err := os.Rename(pather, npat)
		if err != nil {
			fmt.Print(err)
		}
	}
}
