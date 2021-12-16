package main
import (
	"fmt"
	"os"
	"io"
	"io/ioutil"
	"bufio"
)

/*
   subject : fileI/O bufio.NewReader.ReadString('\n')
*/ 
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readfile(filepath string) error {
	dat, err := ioutil.ReadFile(filepath)
	if err!= nil {
		return err
	}
	fmt.Print(string(dat))
	return nil
}

func fileclose(f *os.File) {
	fmt.Println("deferd file close", f)
	f.Close()
	fmt.Println("closed", f)
}

func readlines(filepath string) error {
	var file *os.File
	
	defer fmt.Println("---deferd-return--")
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

	//var fil filehandle
	//err := fil.Open(filepath)

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
	//	readfile("practice19.go")
	err := readlines2("practice19-2.go");
	if err != nil {
		fmt.Println(err)
	}
}
