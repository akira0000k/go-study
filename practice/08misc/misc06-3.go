package main

import (
	"fmt"
	s "strings"
)
/*
   subject : String Functions. 日本語入力に対する処理は如何
*/
var pf = fmt.Printf
var p = fmt.Println

func main() {

	func(name, arg1, arg2 string) {
		pf("%s(%s, %s) = %v\n", name, arg1, arg2, s.Contains(arg1, arg2)) 
	}("Contains", "麻婆茄子", "婆")

	func(name, arg1, arg2 string) {
		pf("%s(%s, %s) = %v\n", name, arg1, arg2, s.Count(arg1, arg2)) 
	}("Count", "麻婆豆腐麻婆茄子", "麻婆")

	func(name, arg1, arg2 string) {
		pf("%s(%s, %s) = %v\n", name, arg1, arg2, s.HasPrefix(arg1, arg2)) 
	}("HasPrefix", "朝辞白帝彩雲間", "朝")

	func(name, arg1, arg2 string) {
		pf("%s(%s, %s) = %v\n", name, arg1, arg2, s.HasSuffix(arg1, arg2)) 
	}("HasSuffix", "両岸猿声啼不住", "住")

	func(name, arg1, arg2 string) {
		pf("%s(%s, %s) = %v\n", name, arg1, arg2, s.Index(arg1, arg2)) 
	}("Index", "麻婆豆腐", "麻")
	func(name, arg1, arg2 string) {
		pf("%s(%s, %s) = %v\n", name, arg1, arg2, s.Index(arg1, arg2)) 
	}("Index", "麻婆豆腐", "婆")
	func(name, arg1, arg2 string) {
		pf("%s(%s, %s) = %v\n", name, arg1, arg2, s.Index(arg1, arg2)) 
	}("Index", "麻婆豆腐", "豆腐")
	func(name, arg1, arg2 string) {
		pf("%s(%s, %s) = %v\n", name, arg1, arg2, s.Index(arg1, arg2)) 
	}("Index", "麻婆豆腐", "腐")
	func(name, arg1, arg2 string) {
		pf("%s(%s, %s) = %v\n", name, arg1, arg2, s.Index(arg1, arg2)) 
	}("Index", "麻婆豆腐", "臭")

	func(name string, arg1 []string, arg2 string) {
		pf("%s(%v, %s) = %v\n", name, arg1, arg2, s.Join(arg1, arg2))
	}("Join", []string{"臭", "豆", "腐"}, "-")
	
	func(name, arg1 string, arg2 int) {
		pf("%s(%s, %v) = %v\n", name, arg1, arg2, s.Repeat(arg1, arg2)) 
	}("Repeat", "餓鬼", 6)

	sReplace("親日親米親ソ親朝親韓", "親", "反", -1)
	sReplace("親日親米親ソ親朝親韓", "親", "反", 0)
	sReplace("親日親米親ソ親朝親韓", "親", "反", 1)
	sReplace("親日親米親ソ親朝親韓", "親", "反", 2)
	sReplace("親日親米親ソ親朝親韓", "親", "反", 3)
	sReplace("親日親米親ソ親朝親韓", "親", "反", 4)
	sReplace("親日親米親ソ親朝親韓", "親", "反", 5)
	sReplace("親日親米親ソ親朝親韓", "親", "反", 6)
	
	sSplit("本日は晴天なり", "は")
	sSplit("天気晴朗なれども波高し", "なれども")
	
	sToLower("テスト")
	sToLower("ﾃｽﾄ")
	sToUpper("てすと")
	sToUpper("ﾃｽﾄ")
	p("\n")
	sLen("こんにちは")
	sLen("今日は")
	sChar("今日は"[0])
	sChar("今日は"[1])
	sChar("今日は"[2])
	sChar("今日は"[3])
	sChar("今日は"[4])
	sRune("雲呑麺", 0)
	sRune("雲呑麺", 1)
	sRune("雲呑麺", 2)
	sRune("wang", 0)
	sRune("wang", 1)
	sRune("wang", 2)
	sRune("wang", 3)
}

func sReplace(arg1, arg2, arg3 string, arg4 int) {
	pf("%s(%s, %s, %s, %d) = %v\n", "Replace", arg1, arg2, arg3, arg4, s.Replace(arg1, arg2, arg3, arg4))
}
func sSplit(arg1, arg2 string) {
	pf("%s(%s, %s) = %v\n", "Split", arg1, arg2, s.Split(arg1, arg2)) 
}
func sToLower(arg1 string) {
	pf("%s(%s) = %v\n", "ToLower", arg1, s.ToLower(arg1)) 
}
func sToUpper(arg1 string) {
	pf("%s(%s) = %v\n", "ToUpper", arg1, s.ToUpper(arg1)) 
}
func sLen(arg1 string) {
	pf("%s(%s) = %v\n", "len", arg1, len(arg1)) 
}

func sChar(arg1 byte) {
	pf("%c %d\n", arg1, arg1)
}
func sRune(arg1 string, arg2 int) {
	count := 0
	for i, r := range arg1 {
		if count == arg2 {
			pf("%d %d %T %c\n", count, i, r, r)
			return
		}
		count++
	}
}
// -*- mode: compilation; default-directory: "~/go/src/practice/08misc/" -*-
// Compilation started at Sat Oct  9 17:20:29
//  
// go run misc06-3.go
// Contains(麻婆茄子, 婆) = true
// Count(麻婆豆腐麻婆茄子, 麻婆) = 2
// HasPrefix(朝辞白帝彩雲間, 朝) = true
// HasSuffix(両岸猿声啼不住, 住) = true
// Index(麻婆豆腐, 麻) = 0
// Index(麻婆豆腐, 婆) = 3
// Index(麻婆豆腐, 豆腐) = 6
// Index(麻婆豆腐, 腐) = 9
// Index(麻婆豆腐, 臭) = -1
// Join([臭 豆 腐], -) = 臭-豆-腐
// Repeat(餓鬼, 6) = 餓鬼餓鬼餓鬼餓鬼餓鬼餓鬼
// Replace(親日親米親ソ親朝親韓, 親, 反, -1) = 反日反米反ソ反朝反韓
// Replace(親日親米親ソ親朝親韓, 親, 反, 0) = 親日親米親ソ親朝親韓
// Replace(親日親米親ソ親朝親韓, 親, 反, 1) = 反日親米親ソ親朝親韓
// Replace(親日親米親ソ親朝親韓, 親, 反, 2) = 反日反米親ソ親朝親韓
// Replace(親日親米親ソ親朝親韓, 親, 反, 3) = 反日反米反ソ親朝親韓
// Replace(親日親米親ソ親朝親韓, 親, 反, 4) = 反日反米反ソ反朝親韓
// Replace(親日親米親ソ親朝親韓, 親, 反, 5) = 反日反米反ソ反朝反韓
// Replace(親日親米親ソ親朝親韓, 親, 反, 6) = 反日反米反ソ反朝反韓
// Split(本日は晴天なり, は) = [本日 晴天なり]
// Split(天気晴朗なれども波高し, なれども) = [天気晴朗 波高し]
// ToLower(テスト) = テスト
// ToLower(ﾃｽﾄ) = ﾃｽﾄ
// ToUpper(てすと) = てすと
// ToUpper(ﾃｽﾄ) = ﾃｽﾄ
//  
//  
// len(こんにちは) = 15
// len(今日は) = 9
// ä 228
// » 187
//  138
// æ 230
//  151
// 0 0 int32 雲
// 1 3 int32 呑
// 2 6 int32 麺
// 0 0 int32 w
// 1 1 int32 a
// 2 2 int32 n
// 3 3 int32 g
//  
// Compilation finished at Sat Oct  9 17:20:29
