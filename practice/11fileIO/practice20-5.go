package main
import (
	"fmt"
	"os"
	"io"
)

/*
   subject : fileI/O file.Read(b) file.WriteString(string(b[:c])):  Ifile == Ofile. complete.
*/ 
func fileclose(f *os.File) {
	fmt.Println("deferd file close", f)
	f.Close()
	fmt.Println("closed", f)
}

func rwlines(fileorig string, filecopy string) error {
	
	file, err := os.Open(fileorig)
	if err!= nil {
		return err
	}
	defer fileclose(file)

	file2, err := os.Create(filecopy)
	if err!= nil {
		return err
	}
	defer fileclose(file2)

	b := make([]byte, 100)
	for {
		c, err := file.Read(b)
		fmt.Println("c=", c, err)
		if err == io.EOF { break }

		if err != nil {
			return err
		}
		//fmt.Println("cr=", c)
		line := string(b[:c])
		c, err = file2.WriteString(line)
		//fmt.Println("cw=", c)
		if err != nil {
			return err
		}
		//_ = c
	}
	fmt.Println("---------done---------")
	return nil
}

func main() {
	fmt.Println("start")
	//err := rwlines("practice20-5.go", "practice20-5-copy.go");
	err := rwlines("dat/combo.txt", "dat/combo-5.txt");
	if err != nil {
		fmt.Println(err)
	}
}
//====================================================================================================================================================================================================//
