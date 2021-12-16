package main
import (
	"fmt"
	"io/ioutil"
)

/*
   subject : fileI/O ioutil.ReadFile
*/ 

func readfile(filepath string) error {
	dat, err := ioutil.ReadFile(filepath)
	if err!= nil {
		return err
	}
	fmt.Print(string(dat))
	return nil
}

func main() {
	fmt.Println("start")
	err := readfile("practice19-0.go")
	if err != nil {
		fmt.Println(err)
	}
}
