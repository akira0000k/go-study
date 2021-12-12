package main
import (
	"fmt"
	"time"
)
/*
   subject: Worker Pools
*/
func worker(id int, jobs <-chan int, results chan<-int, wkend chan<-int) {
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
}

func printres(results <-chan int, wkend chan<-int) {
	for ires := range results {
		fmt.Println("Res=", ires)
	}
	fmt.Println("stop results")
	wkend <- 0
	//time.Sleep(time.Second)
	fmt.Println("End  results")
}

func main() {
	//_ = time.Second
	const numJobs = 10
	const numWks = 3

	fmt.Println("start")


	jobs := make(chan int, numJobs) //なくても動く
	results := make(chan int)//, numJobs) //バツファを付けないとpanicを起こす
	wkend := make(chan int)

	for w := 1; w <= numWks; w++ {
		go worker(w, jobs, results, wkend)
	}
	go printres(results, wkend)
	
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
	//time.Sleep(time.Second * 2)
	fmt.Println("end")
}
// -*- mode: compilation; default-directory: "~/go/src/practice/02sync/" -*-
// Compilation started at Fri Oct 29 18:41:49
//  
// go run sync08.go
// start
// job 1 sent
//  			 worker 1 started  job 1
//  									 worker 3 started  job 2
// job 2 sent
// job 3 sent
//  						 worker 2 started  job 3
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
//  									 worker 3 finished job 2
//  									 worker 3 started  job 5
// Res= 4
//  						 worker 2 finished job 3
// Res= 6
//  						 worker 2 started  job 6
//  			 worker 1 finished job 4
//  			 worker 1 started  job 7
// Res= 8
//  									 worker 3 finished job 5
//  									 worker 3 started  job 8
// Res= 10
//  						 worker 2 finished job 6
//  						 worker 2 started  job 9
// Res= 12
//  			 worker 1 finished job 7
//  			 worker 1 started  job 10
// Res= 14
//  									 worker 3 finished job 8
//  									 stop worker 3
//  									 End  worker 3
// worker 3 ended
// Res= 16
//  						 worker 2 finished job 9
//  						 stop worker 2
//  						 End  worker 2
// worker 2 ended
// Res= 18
//  			 worker 1 finished job 10
//  			 stop worker 1
//  			 End  worker 1
// Res= 20
// worker 1 ended
// stop results
// End  results
// print  0 ended
// end
//  
// Compilation finished at Fri Oct 29 18:41:54
