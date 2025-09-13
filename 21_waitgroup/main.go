package main

import "sync"

//waitgropus are used to wait for a collection of goroutines to finish executing before proceeding
//they provide a way to synchronize the completion of multiple goroutines 

func task(id int, w *sync.WaitGroup){
	defer w.Done() //decrement the counter when the goroutine completes
	println("Task", id, "is being processed")
}


func main(){
	var wg sync.WaitGroup

	for i:=1; i<=5; i++{
		wg.Add(1) //increment the counter for each goroutine
		go task(i, &wg) //launch a goroutine to perform the task
	}

	wg.Wait()

}