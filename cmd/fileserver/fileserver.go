package main

import (
	"flag"
	"github.com/gorilla/mux"
	"net/http"
	"time"
	"log"

	"io/ioutil"
	"github.houston.softwaregrp.net/CSB/fileserver/pkg/fileserver"
	"fmt"
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
	i:=Init{dir}
	p:=&i
	fmt.Print(p)
	fmt.Print(*p)
	r :=mux.NewRouter()
	r.HandleFunc("/", i.HandleRoot)
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/",http.FileServer(http.Dir(dir))))
	srv := &http.Server{
		Handler: r,
		Addr:    port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}


	//node, err := html.Parse()
	///if err != nil{
	//fmt.Print(err)
//}
	//isLiElement(node)
	//isUlElement(node)
	log.Fatal(srv.ListenAndServe())
}

// purpose of this handler is to do a system call for a listing on the configured directory
// this system call will
// eg. something in package os (maybe i forgot)

// essentially google how to list files in a directory
// convert them into html in the form of <ul><li>file1.txt</li><li>file2.txt</li></ul>
// ignore sub directories for now
// send the response with content type text/html


func (i *Init)HandleRoot(w http.ResponseWriter,r *http.Request) {
	fmt.Print("dir",i.Dir)
	bob := RipFunc(i.Dir)
	rob:=fmt.Sprintf("%s%s%s","<ul>",bob,"</ul>")
	fmt.Print(rob)

	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(rob))
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
	//fmt.Print(fum)
	return
}