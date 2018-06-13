package main

import (
	"fmt"
	ss "github.com/nevermosby/how-to-use-golang/sshutils"
	_ "math/rand"
	"sync"
	"time"
)

type Job struct {
	id int
	//randomno int
	command string
}
type Result struct {
	job Job
	//sumofdigits int
	output string
}

var jobs = make(chan Job, 10)
var results = make(chan Result, 10)

func sshOutput(command string) string {
	client, session, err := ss.ConnectToHost("user", "ip:22")
	if err != nil {
		panic(err)
	}
	defer client.Close()
	defer session.Close()
	out, err := session.CombinedOutput(command)
	if err != nil {
		panic(err)
	}
	return string(out)
}

func digits(number int) int {
	sum := 0
	no := number
	for no != 0 {
		digit := no % 10
		sum += digit
		no /= 10
	}
	time.Sleep(2 * time.Second)
	return sum
}
func worker(wg *sync.WaitGroup) {
	for job := range jobs {
		//output := Result{job, digits(job.randomno)}
		output := Result{job, sshOutput(job.command)}

		results <- output
	}
	wg.Done()
}
func createWorkerPool(noOfWorkers int) {
	var wg sync.WaitGroup
	for i := 0; i < noOfWorkers; i++ {
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
	close(results)
}
func allocate(noOfJobs int) {
	command := ""
	for i := 0; i < noOfJobs; i++ {
		if i%2 == 0 {
			command = "uname -a"
		} else {
			command = "hostname -i"
		}
		job := Job{i, command}
		jobs <- job
	}
	close(jobs)
}
func result(done chan bool) {
	for result := range results {
		//fmt.Printf("Job id %d, input random no %d , sum of digits %d\n", result.job.id, result.job.randomno, result.sumofdigits)
		fmt.Printf("Job command %s,  job result %s\n", result.job.command, result.output)

	}
	done <- true
}
func main() {
	startTime := time.Now()
	noOfJobs := 100
	go allocate(noOfJobs)
	done := make(chan bool)
	go result(done)
	noOfWorkers := 10
	createWorkerPool(noOfWorkers)
	<-done
	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Println("total time taken ", diff.Seconds(), "seconds")
}
