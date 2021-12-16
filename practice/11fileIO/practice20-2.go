package main
import (
	"fmt"
	"os"
	"io"
	"bufio"
)

/*
   subject : fileI/O bufio.NewReader.ReadLine: DOS->unix, keep last line without LF(add LF)
        NewScanner.Scan is more convenient.
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

	reader := bufio.NewReader(file)
	writer := bufio.NewWriter(file2)
	for {
		buff, isPrefix, err := reader.ReadLine()
		if err == io.EOF { break }

		if err != nil {
			return err
		}
		c, err := writer.Write(buff)
		if err != nil {
			return err
		}
		if !isPrefix {
			c, err = writer.Write([]byte("\n"))
			if err != nil {
				return err
			}
		}
		_ = c
		//fmt.Println("len=", c)
		writer.Flush()
	}
	fmt.Println("---------done---------")
	return nil
}

func main() {
	fmt.Println("start")
	//err := rwlines("practice20-2.go", "practice20-2-copy.go");
	err := rwlines("dat/combo.txt", "dat/combo-2.txt");
	if err != nil {
		fmt.Println(err)
	}
}
