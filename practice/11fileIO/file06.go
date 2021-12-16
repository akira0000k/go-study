package main
import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)
/*
	subject : サンプルで学ぶ Go 言語：Temporary Files and Directories
*/
func check(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {

	//The easiest way to create a temporary file is by calling ioutil.TempFile.
	//It creates a file and opens it for reading and writing. We provide "" as the first argument,
	//so ioutil.TempFile will create the file in the default location for our OS.
	f, err := ioutil.TempFile("", "sample")
	check(err)

	//Display the name of the temporary file. On Unix-based OSes the directory will likely be /tmp.
	//The file name starts with the prefix given as the second argument to ioutil.TempFile
	//and the rest is chosen automatically to ensure that concurrent calls will always create different file names.
	fmt.Println("Temp file name:", f.Name())

	//Clean up the file after we’re done.
	//The OS is likely to clean up temporary files by itself after some time, but it’s good practice to do this explicitly.
	defer os.Remove(f.Name())

	//We can write some data to the file.
	_, err = f.Write([]byte{1, 2, 3, 4})
	check(err)

	//If we intend to write many temporary files, we may prefer to create a temporary directory.
	//ioutil.TempDir’s arguments are the same as TempFile’s, but it returns a directory name rather than an open file.
	dname, err := ioutil.TempDir("", "sampledir")
	check(err)
	fmt.Println("Temp dir name:", dname)

	defer os.RemoveAll(dname)

	//Now we can synthesize temporary file names by prefixing them with our temporary directory.
	fname := filepath.Join(dname, "file1")
	err = ioutil.WriteFile(fname, []byte{1, 2}, 0666)
	check(err)
}
// -*- mode: compilation; default-directory: "~/go/src/practice/11fileIO/" -*-
// Compilation started at Tue Oct 19 20:40:35
//  
// go run file06.go
// Temp file name: /var/folders/5m/29zwdxmj52q7klt6qnkzxn_40000gp/T/sample100439642
// Temp dir name: /var/folders/5m/29zwdxmj52q7klt6qnkzxn_40000gp/T/sampledir620574705
//  
// Compilation finished at Tue Oct 19 20:40:36
