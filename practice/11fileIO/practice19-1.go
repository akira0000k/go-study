package main
import (
	"fmt"
	"os"
	"io"
	"bufio"
)

/*
   subject : fileI/O bufio.NewReader.ReadString('\n')
*/ 
func fileclose(f *os.File) {
	fmt.Println("deferd file close", f)
	f.Close()
	fmt.Println("closed", f)
}

func readlines(filepath string) error {
	//var file *os.File
	
	file, err := os.Open(filepath)
	if err!= nil {
		return err
	}
	defer fileclose(file) //file.Close()

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF { break }

		if err != nil {
			return err
		}
		fmt.Print(line)
	}
	fmt.Println("---------done---------")
	return nil
}

func main() {
	fmt.Println("start")
	err := readlines("practice19-1.go");
	if err != nil {
		fmt.Println(err)
	}
}
