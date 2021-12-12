package main
import (
	"fmt"
	"time"
)
/*
   subject: Worker Pools noname func
*/


func main() {
	//_ = time.Second
	const numJobs = 10
	const numWks = 3

	fmt.Println("start sync08-2")


	jobs := make(chan int, numJobs) //なくても動く
	results := make(chan int)//, numJobs) //バツファを付けないとpanicを起こす
	wkend := make(chan int)

	for w := 1; w <= numWks; w++ {
		// w だけは閉じ込められなかった。引数にしないと全部 worker 4 startedとなる
		go func(id int) {
			sp := ""
			for i:=0; i<id; i++ {
				sp += "\t\t\t"
			}
			for j := range jobs {
				fmt.Println(sp, "worker", id, "started  job", j)
				time.Sleep(time.Second)
				fmt.Println(sp, "worker", id, "finished job", j)
				results <- j*2
			}
			fmt.Println(sp, "stop worker", id)
			wkend <- id
			fmt.Println(sp, "End  worker", id)
		}(w)

	}

	go func() {
		for ires := range results {
			fmt.Println("Res=", ires)
		}
		fmt.Println("stop results")
		wkend <- 0
		fmt.Println("End  results")
	}()
	
	for j:=1; j<=numJobs; j++ {
		jobs<-j
		fmt.Println("job", j, "sent") // jobsにバッファーがないと引っかかる
		time.Sleep(time.Second / 10)
	}
	close(jobs) //workerを終了させるため

	//for a:=1; a<=numJobs; a++ {
	// 	fmt.Println("Res=", <-results)
	//}
	for w := 1; w <= numWks; w++ {
		fmt.Println("worker", <-wkend, "ended")
	}
	close(results)
	fmt.Println("print ", <-wkend, "ended")
	//time.Sleep(time.Second * 20)
	fmt.Println("end")
}
// -*- mode: compilation; default-directory: "~/go/src/practice/02sync/" -*-
// Compilation started at Fri Oct 29 18:45:59
//  
// go run sync08-2.go
// start sync08-2
// job 1 sent
//  			 worker 1 started  job 1
// job 2 sent
//  						 worker 2 started  job 2
//  									 worker 3 started  job 3
// job 3 sent
// job 4 sent
// job 5 sent
// job 6 sent
// job 7 sent
// job 8 sent
// job 9 sent
// job 10 sent
//  			 worker 1 finished job 1
//  			 worker 1 started  job 4
// Res= 2
//  						 worker 2 finished job 2
//  						 worker 2 started  job 5
// Res= 4
//  									 worker 3 finished job 3
//  									 worker 3 started  job 6
// Res= 6
//  			 worker 1 finished job 4
//  			 worker 1 started  job 7
// Res= 8
//  						 worker 2 finished job 5
//  						 worker 2 started  job 8
// Res= 10
//  									 worker 3 finished job 6
//  									 worker 3 started  job 9
// Res= 12
//  			 worker 1 finished job 7
//  			 worker 1 started  job 10
// Res= 14
//  						 worker 2 finished job 8
// Res= 16
//  						 stop worker 2
//  						 End  worker 2
// worker 2 ended
//  									 worker 3 finished job 9
//  									 stop worker 3
//  									 End  worker 3
// worker 3 ended
// Res= 18
//  			 worker 1 finished job 10
//  			 stop worker 1
//  			 End  worker 1
// worker 1 ended
// Res= 20
// stop results
// End  results
// print  0 ended
// end
//  
// Compilation finished at Fri Oct 29 18:46:04
