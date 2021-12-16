package main
import (
	"fmt"
	"os"
	"io"
	"bufio"
)

/*
   subject : fileI/O bufio.NewReader.ReadString('\n'): May disapper last line without LF.
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
		//  ReadString('\n') can not read last line without LF
		line, err := reader.ReadString('\n')
		fmt.Println(err)
		fmt.Print(line)
		if err == io.EOF { break }

		if err != nil {
			return err
		}
		//c, err := fmt.Fprint(writer, line)
		//c, err := writer.Write([]byte(line))
		c, err := writer.WriteString(line)
		if err != nil {
			return err
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
	//err := rwlines("practice20-1.go", "practice20-1-copy.go");
	err := rwlines("dat/combo.txt", "dat/combo-1.txt");
	if err != nil {
		fmt.Println(err)
	}
}
