package main
import (
	"fmt"
	"os"
	"bufio"
)

/*
   subject : fileI/O bufio.NewScanner.Scan/Text writer.WriteString: DOS->unix, keep last line without LF(add LF)
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

	scanner := bufio.NewScanner(file)
	writer := bufio.NewWriter(file2)
	for scanner.Scan() {
		line := scanner.Text()
		_, err := writer.WriteString(line)
		if err != nil {
			return err
		}
		_, err = writer.WriteString("\n")
		if err != nil {
			return err
		}
		writer.Flush()
	}
	err = scanner.Err()
	if err != nil {
		return err
	}
	fmt.Println("---------done---------")
	return nil
}

func main() {
	fmt.Println("start")
	//err := rwlines("practice20-3.go", "practice20-3-copy.go");
	err := rwlines("dat/combo.txt", "dat/combo-3-1.txt");
	if err != nil {
		fmt.Println(err)
	}
}
