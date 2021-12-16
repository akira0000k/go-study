package main
import (
	"fmt"
	"time"
)
/*
   subject : time functions
*/
func main() {
	p := fmt.Println

	//まずは現在時刻を取得する。
	

	now := time.Now()//2021-10-14 17:31:25.021721 +0900 JST m=+0.000093062
	p(now)

	// time 構造体を、年・月・日などを指定して作ることもできる。 時刻は常に位置（Location）、すなわちタイムゾーンと結びついている。
	

	then := time.Date(
		2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	p(then)//2009-11-17 20:34:58.651387237 +0000 UTC

	// 時刻の構成要素の取り出しかたは直感的だ。
	

	p(then.Year())//2009
	p(then.Month())//November
	p(then.Day())//17
	p(then.Hour())//20
	p(then.Minute())//34
	p(then.Second())//58
	p(then.Nanosecond())//651387237
	p(then.Location())//UTC

	// 曜日（Monday-Sunday）を表す Weekday もある。
	

	p(then.Weekday())//Tuesday

	// 以下のメソッドは2つの時刻を比べて、1つ目が2つ目より前か、後か、それとも同時刻かをそれぞれテストする。
	

	p(then.Before(now))//true
	p(then.After(now))//false
	p(then.Equal(now))//false

	// メソッド Sub は2つの時刻の間隔を表す経過時間 Duration を返す。
	

	diff := now.Sub(then)
	p(diff)//104363h56m26.370333763s

	// 経過時間の長さの単位を変換できる。
	

	p(diff.Hours())//104363.94065842604
	p(diff.Minutes())//6.261836439505563e+06
	p(diff.Seconds())//3.757101863703338e+08
	p(diff.Nanoseconds())//375710186370333763

	// Add を使って指定した期間だけ時刻を進めることもできる。- を使えば、時刻を戻すこともできる。
	
	p(then.Add(diff))//2021-10-14 08:31:25.021721 +0000 UTC
	p(then.Add(-diff))//1997-12-22 08:38:32.281053474 +0000 UTC
}
// -*- mode: compilation; default-directory: "~/go/src/practice/08misc/" -*-
// Compilation started at Thu Oct 14 17:31:22
//  
// go run misc09.go
// 2021-10-14 17:31:25.021721 +0900 JST m=+0.000093062
// 2009-11-17 20:34:58.651387237 +0000 UTC
// 2009
// November
// 17
// 20
// 34
// 58
// 651387237
// UTC
// Tuesday
// true
// false
// false
// 104363h56m26.370333763s
// 104363.94065842604
// 6.261836439505563e+06
// 3.757101863703338e+08
// 375710186370333763
// 2021-10-14 08:31:25.021721 +0000 UTC
// 1997-12-22 08:38:32.281053474 +0000 UTC
//  
// Compilation finished at Thu Oct 14 17:31:25
