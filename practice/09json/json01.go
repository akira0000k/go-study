package main
import (
	"encoding/json"
	"fmt"
	"os"
)
/*
   subject : JSON
*/
type response1 struct {
	Page   int
	Fruits []string
}

type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

type glocery struct {
	Apple int
	Lettuce int
}
func main() {
	if true {
		bolB, _ := json.Marshal(true)
		fmt.Printf("%T %v\n", bolB, bolB)
		fmt.Println(string(bolB))//true

		intB, _ := json.Marshal(1)
		fmt.Printf("%T %v\n", intB, intB)
		fmt.Println(string(intB))//1

		fltB, _ := json.Marshal(2.34)
		fmt.Printf("%T %v\n", fltB, fltB)
		fmt.Println(string(fltB))//2.34

		strB, _ := json.Marshal("gopher")
		fmt.Printf("%T %v\n", strB, strB)
		fmt.Println(string(strB))//"gopher"

		slcD := []string{"apple", "peach", "pear"}
		slcB, _ := json.Marshal(slcD)
		fmt.Printf("%T %v\n", slcB, slcB)
		fmt.Println(string(slcB))//["apple","peach","pear"]

		mapD := map[string]int{"apple": 5, "lettuce": 7}
		mapB, _ := json.Marshal(mapD)
		fmt.Println(string(mapB))//{"apple":5,"lettuce":7}

		//var glocerystore glocery
		//glocerystore.Apple = 5
		//glocerystore.Lettuce = 7
		//strcC := &glocerystore
		//fmt.Println(glocerystore)
		//fmt.Println(&glocerystore)
		////strcC := &struct { Apple int; Lettuce int }{5, 7}
		//////strcC := &glocery{ 5, 7 }
		strcC := &glocery{
			Apple: 5,
			Lettuce: 7,
		}
		strcB, ok := json.Marshal(strcC)
		fmt.Println(ok, string(strcB))//<nil> {"Apple":5,"Lettuce":7}
		
		//res1D := &response1{
		// 	Page:	1,
		// 	Fruits: []string{"apple", "peach", "pear"}}
		res1D := &response1{ 1, []string{"apple", "peach", "pear"}}
		fmt.Printf("%+v\n", *res1D)//{Page:1 Fruits:[apple peach pear]}
		fmt.Printf("%#v\n", *res1D)//main.response1{Page:1, Fruits:[]string{"apple", "peach", "pear"}}

		res1B, _ := json.Marshal(res1D)
		fmt.Println(string(res1B))//{"Page":1,"Fruits":["apple","peach","pear"]}

		res2D := &response2{
			Page:   1,
			Fruits: []string{"apple", "peach", "pear"}}
		res2B, _ := json.Marshal(res2D)
		fmt.Println(string(res2B))//{"page":1,"fruits":["apple","peach","pear"]}
		fmt.Println("\n===================================\n")
	}

	if true {
		byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

		var dat map[string]interface{}

		if err := json.Unmarshal(byt, &dat); err != nil {
			panic(err)
		}
		fmt.Println(dat)//map[num:6.13 strs:[a b]]

		num := dat["num"].(float64)
		fmt.Println(num)//6.13

		strs := dat["strs"].([]interface{})
		str1 := strs[0].(string)
		fmt.Println(str1)//a

		str := `{"page": 1, "fruits": ["apple", "peach"]}`
		res := response2{}
		json.Unmarshal([]byte(str), &res)
		fmt.Println(res)//{1 [apple peach]}
		fmt.Println(res.Fruits[0])//apple

		enc := json.NewEncoder(os.Stdout)
		d := map[string]int{"apple": 5, "lettuce": 7}
		enc.Encode(d)//{"apple":5,"lettuce":7}
	}
}
// -*- mode: compilation; default-directory: "~/go/src/practice/09json/" -*-
// Compilation started at Sat Oct  9 22:39:52
//  
// go run json01.go
// []uint8 [116 114 117 101]
// true
// []uint8 [49]
// 1
// []uint8 [50 46 51 52]
// 2.34
// []uint8 [34 103 111 112 104 101 114 34]
// "gopher"
// []uint8 [91 34 97 112 112 108 101 34 44 34 112 101 97 99 104 34 44 34 112 101 97 114 34 93]
// ["apple","peach","pear"]
// {"apple":5,"lettuce":7}
// <nil> {"Apple":5,"Lettuce":7}
// {Page:1 Fruits:[apple peach pear]}
// main.response1{Page:1, Fruits:[]string{"apple", "peach", "pear"}}
// {"Page":1,"Fruits":["apple","peach","pear"]}
// {"page":1,"fruits":["apple","peach","pear"]}
//  
// ===================================
//  
// map[num:6.13 strs:[a b]]
// 6.13
// a
// {1 [apple peach]}
// apple
// {"apple":5,"lettuce":7}
//  
// Compilation finished at Sat Oct  9 22:39:52
