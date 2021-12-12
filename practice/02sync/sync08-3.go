package main
import (
	"fmt"
	"time"
	"sync"
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
	results := make(chan int, numJobs) //Res受け取りに戸惑っても投げられる
	var wg, wgp sync.WaitGroup
	
	for w := 1; w <= numWks; w++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			sp := ""
			for i:=0; i<id; i++ {
				sp += "\t\t\t"
			}
			for j := range jobs {
				fmt.Print(sp, "worker ", id, " started  job ", j, "\n")
				time.Sleep(time.Second)
				fmt.Print(sp, "worker ", id, " finished job ", j, "\n")
				results <- j*2
			}
			//fmt.Println(sp, "stop worker", id)
			//wkend <- id
			fmt.Printf("%sEnd  worker %d\n", sp,  id)
		}(w)

	}
	wgp.Add(1)
	go func() {
		defer wgp.Done()
		for ires := range results {
			fmt.Println("Res=", ires)
			time.Sleep(time.Second)
		}
		fmt.Println("End  results")
	}()
	
	for j:=1; j<=numJobs; j++ {
		jobs<-j
		fmt.Println("job", j, "sent") // jobsにバッファーがないと引っかかる
		time.Sleep(time.Second / 10)
	}
	close(jobs) //workerを終了させるため

	fmt.Println("Waiting.....")
	wg.Wait()
	fmt.Println(".....all worker done")

	fmt.Println("close results.....")
	close(results) //printを終了させるため

	wgp.Wait()
	fmt.Println(".....results ended")

	
	//time.Sleep(time.Second * 20)
	fmt.Println("end")
}
// -*- mode: compilation; default-directory: "~/go/src/practice/02sync/" -*-
// Compilation started at Fri Oct 29 18:49:34
//  
// go run sync08-3.go
// start sync08-2
// job 1 sent
//  			worker 1 started  job 1
// job 2 sent
//  									worker 3 started  job 2
// job 3 sent
//  						worker 2 started  job 3
// job 4 sent
// job 5 sent
// job 6 sent
// job 7 sent
// job 8 sent
// job 9 sent
// job 10 sent
//  			worker 1 finished job 1
//  			worker 1 started  job 4
// Res= 2
// Waiting.....
//  									worker 3 finished job 2
//  									worker 3 started  job 5
//  						worker 2 finished job 3
//  						worker 2 started  job 6
// Res= 4
//  			worker 1 finished job 4
//  			worker 1 started  job 7
//  									worker 3 finished job 5
//  									worker 3 started  job 8
//  						worker 2 finished job 6
//  						worker 2 started  job 9
// Res= 6
//  			worker 1 finished job 7
//  			worker 1 started  job 10
//  									worker 3 finished job 8
//  									End  worker 3
//  						worker 2 finished job 9
//  						End  worker 2
//  			worker 1 finished job 10
//  			End  worker 1
// .....all worker done
// close results.....
// Res= 8
// Res= 10
// Res= 12
// Res= 14
// Res= 16
// Res= 18
// Res= 20
// End  results
// .....results ended
// end
//  
// Compilation finished at Fri Oct 29 18:49:46
