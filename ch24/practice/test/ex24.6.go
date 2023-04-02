package main

import (
	"fmt"
	"sync"
	"time"
)

type Job interface {
	Do()
}

type SquareJob struct {
	index int
}

// *SquareJob는 Do() 메서드를 가지고 있다. 따라서, DO() 메서드를 호출할 수 있는 Job 인터페이스를 사용할 수 있다.
// SqureJob은 index 필드를 가지므로, Do() 메서드 호출시 index 필드를 기반으로 작업을 수행할 수 있다.
func (j *SquareJob) Do() {
	fmt.Printf("%d 작업 시작\n", j.index)
	time.Sleep(1 * time.Second)
	fmt.Printf("%d 작업 완료 - 결과 : %d\n", j.index, j.index*j.index)
}

func main() {
	// 10개의 작업을 만든다
	var jobList [10]Job

	// 각 작업마다 다른 메모리를 할당한다.
	for i := 0; i < 10; i++ {
		jobList[i] = &SquareJob{i}
	}

	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		job := jobList[i]
		go func() {
			job.Do()
			wg.Done()
		}()
	}
	wg.Wait()
}
