package main
import (
	"fmt"
	"time"
)
/*
   subject : Unix 時間
*/
func main() {

	//time.Now と、Unix か UnixNano を使って、Unix エポックからの経過時間を秒、ナノ秒単位で取得する。
	

	now := time.Now()
	secs := now.Unix()
	nanos := now.UnixNano()
	fmt.Println(now)//2021-10-14 17:44:05.28246 +0900 JST m=+0.000163000

	//UnixMillis はない。 そのため、エポックからの経過時間をミリ秒単位で知りたいときはナノ秒の結果を割る必要がある。
	

	millis := nanos / 1000000
	fmt.Println(secs)   //1634201045
	fmt.Println(millis) //1634201045282
	fmt.Println(nanos)  //1634201045282460000

	//逆に、エポックからの秒、ナノ秒での経過時間をtime に変換することもできる。

	fmt.Println(time.Unix(secs, 0))        //2021-10-14 17:44:05 +0900 JST
	fmt.Println(time.Unix(0, nanos))       //2021-10-14 17:44:05.28246 +0900 JST
	nanopart := nanos - secs * 1000000000
	fmt.Println(time.Unix(secs, nanopart)) //2021-10-14 17:49:53.48518 +0900 JST
}
// -*- mode: compilation; default-directory: "~/go/src/practice/08misc/" -*-
// Compilation started at Thu Oct 14 17:49:52
//  
// go run misc10.go
// 2021-10-14 17:49:53.48518 +0900 JST m=+0.000163742
// 1634201393
// 1634201393485
// 1634201393485180000
// 2021-10-14 17:49:53 +0900 JST
// 2021-10-14 17:49:53.48518 +0900 JST
// 2021-10-14 17:49:53.48518 +0900 JST
//  
// Compilation finished at Thu Oct 14 17:49:53
