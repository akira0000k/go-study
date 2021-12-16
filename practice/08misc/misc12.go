package main
import (
	"fmt"
	"math/rand"
	"time"
)
/*
   subject : Random Numbers
*/
func main() {

	//例えば、rand.Intn は int 型の値 n であって、0 <= n < 100 を満たすものをランダムに選んで返す。
	

	fmt.Print(rand.Intn(100), ",")
	fmt.Print(rand.Intn(100))
	fmt.Println()

	//rand.Float64 は float64 型の値 f であって、0.0 <= f < 1.0 を満たすものをランダムに選んで返す。
	

	fmt.Println(rand.Float64())

	//これを使って別の範囲の float 型の値を作ることもできる。 例えば5.0 <= f' < 10.0 だと以下のようになる。
	

	fmt.Print((rand.Float64() * 5) + 5, ",")
	fmt.Print((rand.Float64() * 5) + 5)
	fmt.Println()

	//デフォルトの生成機は決定的なので、プログラムを実行するたびに同じ数の列を生成する。
	//値を変えるためには、プログラムを実行するたびに変わるシードを渡せばよい。
	//なお、math/rand の生成する乱数は、暗号用に使うべきではない。
	//このような用途では、crypto/rand を使うべきだ。
	

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	//r1 は rand パッケージの関数と同様のメソッドを持っている。
	

	fmt.Print(r1.Intn(100), ",")
	fmt.Print(r1.Intn(100))
	fmt.Println()

	//同じ数をシードにすると、同じ乱数列を生成する。
	

	s2 := rand.NewSource(42)
	r2 := rand.New(s2)
	fmt.Print(r2.Intn(100), ",")
	fmt.Print(r2.Intn(100))
	fmt.Println()
	s3 := rand.NewSource(42)
	r3 := rand.New(s3)
	fmt.Print(r3.Intn(100), ",")
	fmt.Print(r3.Intn(100))
}
// -*- mode: compilation; default-directory: "~/go/src/practice/08misc/" -*-
// Compilation started at Thu Oct 14 18:19:58
//  
// go run misc12.go 
// 81,87
// 0.6645600532184904
// 7.1885709359349015,7.123187485356329
// 36,85
// 5,87
// 5,87
// Compilation finished at Thu Oct 14 18:19:59
