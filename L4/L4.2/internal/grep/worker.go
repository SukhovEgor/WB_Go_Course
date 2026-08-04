package grep

import (
	"sync"

	"DistributedGrep/internal/app"
)

func RunWorkers(

	workers int,

	req app.GrepRequest,

	lines <-chan string,

) <-chan string {

	results := make(chan string)

	var wg sync.WaitGroup

	wg.Add(workers)

	for i := 0; i < workers; i++ {

		go func() {

			defer wg.Done()

			for line := range lines {

				if Match(line, req) {
					results <- line
				}
			}

		}()

	}

	go func() {

		wg.Wait()

		close(results)

	}()

	return results
}