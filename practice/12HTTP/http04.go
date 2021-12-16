package main

//crypto/* パッケージには何種類かのハッシュ関数が含まれている。

import (
	"crypto/sha1"
	"fmt"
)
/*
   subject : サンプルで学ぶ Go 言語：SHA1 Hashes   https://www.spinute.org/go-by-example/sha1-hashes.html
*/
func main() {
	s := "sha1 this string"

	//ハッシュ値を生成するには sha1.New()、sha1.Write(bytes)、sha1.Sum([]byte{}) の順で関数を呼ぶ。 まずは新たなハッシュ関数を生成する。
	h := sha1.New()

	//Write の入力はバイト列である。 文字列 s のハッシュ値を計算したいなら、[]byte(s) と書いてハッシュ値に変換してやらなければならない。
	h.Write([]byte(s))

	//バイトのスライスとして、最終的なハッシュ値を得る。 Sum の引数を、これまで入力したバイト列に追記できるが、普通はこれは使わない。
	bs := h.Sum(nil)

	//SHA1 のハッシュ値は、Git がそうしているように、16進記数法で表示することが多い。 フォーマット文字列に %x と書けば、ハッシュ計算の結果を16進文字列に変換できる。
	fmt.Println(s)
	fmt.Printf("%x\n", bs)
}
// -*- mode: compilation; default-directory: "~/go/src/practice/12HTTP/" -*-
// Compilation started at Tue Oct 19 23:21:05
//  
// go run http04.go
// sha1 this string
// cf23df2207d99a74fbe169e3eba035e633b65d94
//  
// Compilation finished at Tue Oct 19 23:21:05
