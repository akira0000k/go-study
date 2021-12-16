package main
import (
	"fmt"
	"time"
)
/*
   subject : Time Formatting / Parsing
*/
func main() {
	p := fmt.Println

	//まずは RFC3339 に対応するレイアウトを使って時刻をフォーマットする方法を紹介する。
	

	t := time.Now()
	p(t.Format(time.RFC3339))//2021-10-14T17:59:54+09:00

	//時刻をパースするときも Format と同じレイアウトを使う。
	

	t1, e := time.Parse(
		time.RFC3339,
		"2012-11-01T22:08:41+00:00")
	p(t1)//2012-11-01 22:08:41 +0000 +0000

	//Format、Parse では例示に基づいてレイアウトを決める。
	//ふつうは time モジュールに定義されている定数をレイアウトの例として使うが、独自のレイアウトを使ってもよい。
	//レイアウトでは特定の時刻 Mon Jan 2 15:04:05 MST 2006 を表している必要があり、プログラムはこれに従って時刻をフォーマットしたり、文字列をパースしたりする。
	//時刻の例はちょうど以下のようなものだ。
	//例えば、以下の例では、年は2006に、時間は15に、曜日は月曜になっている。
	

	p(t.Format("3:04PM"))//5:59PM
	p(t.Format("Mon Jan _2 15:04:05 2006"))//Thu Oct 14 17:59:54 2021
	p(t.Format("2006-01-02T15:04:05.999999-07:00"))//2021-10-14T17:59:54.429731+09:00

	form := "3 04 PM"
	t2, e := time.Parse(form, "8 41 PM")
	p(t2)//0000-01-01 20:41:00 +0000 UTC

	//フォーマットしたいのが数値なら、時刻値の部分を抜き出し、文字列のフォーマット機能を使ってもよい。
	

	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())//2021-10-14T17:59:54-00:00

	//Parse は入力が不正だと、パース時に起きた問題を説明するエラーを返す。
	

	ansic := "Mon Jan _2 15:04:05 2006"
	_, e = time.Parse(ansic, "8:41PM")
	p(e)//parsing time "8:41PM" as "Mon Jan _2 15:04:05 2006": cannot parse "8:41PM" as "Mon"
}
// -*- mode: compilation; default-directory: "~/go/src/practice/08misc/" -*-
// Compilation started at Thu Oct 14 18:09:03
//  
// go run misc11.go
// 2021-10-14T18:09:03+09:00
// 2012-11-01 22:08:41 +0000 +0000
// 6:09PM
// Thu Oct 14 18:09:03 2021
// 2021-10-14T18:09:03.977893+09:00
// 0000-01-01 20:41:00 +0000 UTC
// 2021-10-14T18:09:03-00:00
// parsing time "8:41PM" as "Mon Jan _2 15:04:05 2006": cannot parse "8:41PM" as "Mon"
//  
// Compilation finished at Thu Oct 14 18:09:03
