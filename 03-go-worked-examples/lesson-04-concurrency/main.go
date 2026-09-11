package main

import (
    "context"
    "fmt"
    "sync"
    "time"
)

type Job struct {
    ID int
}

func worker(ctx context.Context, id int, jobs <-chan Job, wg *sync.WaitGroup) {
    defer wg.Done()
    for {
        select {
        case <-ctx.Done():
            return
        case job, ok := <-jobs:
            if !ok {
                return
            }
            fmt.Printf("worker %d processed job %d\n", id, job.ID)
            time.Sleep(50 * time.Millisecond)
        }
    }
}

func main() {
    ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
    defer cancel()

    jobs := make(chan Job)
    var wg sync.WaitGroup

    for i := 1; i <= 3; i++ {
        wg.Add(1)
        go worker(ctx, i, jobs, &wg)
    }

    for i := 1; i <= 10; i++ {
        jobs <- Job{ID: i}
    }
    close(jobs)
    wg.Wait()
}
