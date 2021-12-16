package main
import (
	"fmt"
	"os"
	"io"
	"bufio"
)

/*
   subject : fileI/O clss: bufio.NewReader.ReadString('\n')
*/ 
//------------class ---------file line read-----------
type filehandle struct {
	file *os.File
	reader *bufio.Reader
}
func (f *filehandle) Open(filepath string) error {
	file, err := os.Open(filepath)
	if err!= nil {
		return err
	}
	f.file = file
	f.reader = bufio.NewReader(f.file)
	return nil
}
func (f *filehandle) Close() {
	f.file.Close()
}
func (f *filehandle) Read() (string, error) {
	line, err := f.reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return line, nil
}
func openLineRead(filepath string) (*filehandle, error) {
	fil := filehandle{}
	err := fil.Open(filepath)
	return &fil, err
}
//---------------------------------------------------

func readlines2(filepath string) error {

	fil, err := openLineRead(filepath)
	if err!= nil {
		return err
	}
	defer fil.Close()

	for {
		line, err := fil.Read()
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

	err := readlines2("practice19-class.go");
	if err != nil {
		fmt.Println(err)
	}
}
