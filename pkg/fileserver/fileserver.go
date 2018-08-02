package fileserver

import (
	"fmt"
	"mime/multipart"
	"os"
)

func FileFilter(f os.FileInfo) (fe string) {

	fe= fmt.Sprint("<li>","<a href=","/static/",f.Name(),">",f.Name(),"</a>","</li>")
	return
}

func Move(pather string, patho string, handler *multipart.FileHeader, l string, Dir string) {
	if l == "" {
		npat := fmt.Sprint(patho + "/" + Dir + "/" + handler.Filename)
		fmt.Print("\n npat1:", npat)
		err := os.Rename(pather, npat)
		if err != nil {
			fmt.Print(err)
		}
	} else {
		npat := fmt.Sprint(patho + l + handler.Filename)
		fmt.Print("\n npat2:", npat)
		err := os.Rename(pather, npat)
		if err != nil {
			fmt.Print(err)
		}
	}

}
