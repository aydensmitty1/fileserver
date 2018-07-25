package fileserver

import (
	"fmt"
	"os"
)

func FileFilter(f os.FileInfo) (fe string) {
		if f.IsDir() == false {
			fe= fmt.Sprint("<li>",f.Name(),"</li>")
		}
	return
}
//func HandleRoot(fn string)string {
//	root := "./"
//	fi, err := ioutil.ReadDir(root)
//	if err != nil {
//		log.Fatal(err)
//	}
//	for _, fo := range fi {
///
//
//		fn = FileFilter(fn, fo)
//	}
//	return fn
//}