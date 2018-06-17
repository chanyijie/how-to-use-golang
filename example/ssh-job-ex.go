package main

import (
	"fmt"
	ss "github.com/nevermosby/how-to-use-golang/utils/ssh"
	_ "math/rand"
	"sync"
	"time"
)

type Job struct {
	id int
	command string
}
type Result struct {
	job Job
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

func worker(wg *sync.WaitGroup) {
	for job := range jobs {
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
