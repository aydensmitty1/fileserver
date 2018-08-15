package fileserver

import (
	"fmt"
		"os"
)

func FileFilter(f os.FileInfo) (fe string) {

	fe= fmt.Sprint("<li>","<a href=","/static/",f.Name(),">",f.Name(),"</a>","</li>")
	return
}