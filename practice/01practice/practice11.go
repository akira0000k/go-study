package main

import "fmt"
/*
   subject : boolean flag check
 */ 
func main1() {
	fmt.Println("start flag test")
	var flag bool=false
	flag = true
	i:=3
	flag = 0 < i && i < 10
	
	fmt.Println(flag)
	fmt.Printf("%v\n", flag)
	fmt.Printf("%x\n", flag)
	
	if flag {
		fmt.Println("flag=TRUE")
	} else {
		fmt.Println("flag=FALSE")
	}
	if !flag {
		fmt.Println("flag=FALSE")
	} else {
		fmt.Println("flag=TRUE")
	}
	if flag!=false {
		fmt.Println("flag=TRUE")
	} else {
		fmt.Println("flag=FALSE")
	}
	if flag==false {
		fmt.Println("flag=FALSE")
	} else {
		fmt.Println("flag=TRUE")
	}
	if flag==true {
		fmt.Println("flag=TRUE")
	} else {
		fmt.Println("flag=FALSE")
	}
	if flag!=true {
		fmt.Println("flag=FALSE")
	} else {
		fmt.Println("flag=TRUE")
	}
		
	fmt.Println("end")
}

func main2() {
	fmt.Println("start int flag test")
	var flag int=0
	flag = 1
	//i:=3
	//flag = 0 < i && i < 10
	
	fmt.Printf("%v\n", flag)
	fmt.Printf("%x\n", flag)
	
	//if flag {
	// 	fmt.Println("flag=TRUE")
	//} else {
	// 	fmt.Println("flag=FALSE")
	//}
	//if !flag {
	// 	fmt.Println("flag=FALSE")
	//} else {
	// 	fmt.Println("flag=TRUE")
	//}
	if flag!=0 {
		fmt.Println("flag=TRUE")
	} else {
		fmt.Println("flag=FALSE")
	}
	if flag==0 {
		fmt.Println("flag=FALSE")
	} else {
		fmt.Println("flag=TRUE")
	}
	if flag==1 {
		fmt.Println("flag=TRUE")
	} else {
		fmt.Println("flag=FALSE")
	}
	if flag!=1 {
		fmt.Println("flag=FALSE")
	} else {
		fmt.Println("flag=TRUE")
	}
		
	fmt.Println("end")
}

func main() {
	main1()
	main2()
}
// -*- mode: compilation; default-directory: "~/go/src/practice/01practice/" -*-
// Compilation started at Sat Oct 30 00:15:08
//  
// go run practice11.go
// start flag test
// true
// true
// %!x(bool=true)
// flag=TRUE
// flag=TRUE
// flag=TRUE
// flag=TRUE
// flag=TRUE
// flag=TRUE
// end
// start int flag test
// 1
// 1
// flag=TRUE
// flag=TRUE
// flag=TRUE
// flag=TRUE
// end
//  
// Compilation finished at Sat Oct 30 00:15:08
