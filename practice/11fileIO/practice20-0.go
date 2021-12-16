package main
import (
	"fmt"
	"io/ioutil"
	"os"
)

/*
   subject : fileI/O ioutil.ReadFile ioutil.WriteFile:  Ifile == Ofile. complete.
*/ 

func readfile(filepath string) ([]byte, error) {
	dat, err := ioutil.ReadFile(filepath)
	return dat, err
}
func writefile(filepath string, dat []byte) error {
	err := ioutil.WriteFile(filepath, dat, 0666)
	return err
}

func main() {
	fmt.Println("start")
	dat, err := readfile("dat/combo.txt") //practice20-0.go")
	if err != nil {
		fmt.Println("read error")
		fmt.Fprintln(os.Stderr, err)
	} else {
		//err = writefile("practice20-0-copy.go", dat)
		err = writefile("dat/combo-0.txt", dat)
		if err != nil {
			fmt.Fprintln(os.Stderr, "write error:", err)
		}
	}
	fmt.Println("err=", err)
}
