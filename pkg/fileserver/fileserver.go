package fileserver

import (
	"fmt"
	"os"
)

func FileFilter(f os.FileInfo) (fe string) {
		if f.IsDir() == false {
			fe= fmt.Sprint("<li>","<a href=","/static/",f.Name()," download target=","_blank",">",f.Name(),"</a>","</li>")
		}
	return
}
