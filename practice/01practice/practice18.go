package main
import (
	"fmt"
)

/*
   subject : new   とほほのGo言語入門
*/ 
type Book struct {
	bookttl string
}
func main() {
	main1()
	main2()
}
func main1() {
	fmt.Println("start")
	booklist := []*Book{}

	for i:=0; i<10; i++ {
		book := new(Book)
		book.bookttl = fmt.Sprintf("Title#%d", i)
		booklist = append(booklist, book)
	}

	for _, book := range booklist {
		fmt.Println(book.bookttl)
	}
}

//func makebook(ttl string) *Book {
// 	//var book Book = Book{ttl}
// 	//var book Book; 	book.bookttl = ttl
// 	book := Book{ttl}
// 	return &book
//}

//func makebook(ttl string) *Book {
// 	book := new(Book)
// 	book.bookttl = ttl
// 	return &book //./practice18.go:38:9: cannot use &book (type **Book) as type *Book in return argument
//}

func makebook(ttl string) *Book {
	pbook := new(Book)
	pbook.bookttl = ttl
	return pbook
}

func main2() {
	fmt.Println("start")
	booklist := []*Book{}

	for i:=0; i<5; i++ {
		booklist = append(booklist, makebook(fmt.Sprintf("ttl-%d", i)))
	}
	for _, book := range booklist {
		fmt.Println(book.bookttl)
	}
}
// -*- mode: compilation; default-directory: "~/go/src/practice/01practice/" -*-
// Compilation started at Sat Oct 30 15:17:58
//  
// go run practice18.go
// start
// Title#0
// Title#1
// Title#2
// Title#3
// Title#4
// Title#5
// Title#6
// Title#7
// Title#8
// Title#9
// start
// ttl-0
// ttl-1
// ttl-2
// ttl-3
// ttl-4
//  
// Compilation finished at Sat Oct 30 15:17:59
