package quorum

import (
	"sync"

	"DistributedGrep/internal/app"
	"DistributedGrep/internal/network"
)

type Quorum struct {
	client *network.Client
}

func New(client *network.Client) *Quorum {
	return &Quorum{
		client: client,
	}
}

type result struct {
	resp *app.GrepResponse
	err  error
}

// Execute отправляет запрос всем серверам,
// собирает результаты и проверяет,
// достигнут ли кворум.
func (q *Quorum) Execute(
	nodes []string,
	req app.GrepRequest,
) (*app.GrepResponse, error) {

	majority := len(nodes)/2 + 1

	results := make(chan result, len(nodes))

	var wg sync.WaitGroup

	for _, node := range nodes {

		wg.Add(1)

		go func(addr string) {
			defer wg.Done()

			resp, err := q.client.Send(addr, req)

			results <- result{
				resp: resp,
				err:  err,
			}
		}(node)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	success := 0

	final := &app.GrepResponse{}

	for res := range results {

		if res.err != nil {
			continue
		}

		success++

		final.Count += res.resp.Count
		final.Lines = append(final.Lines, res.resp.Lines...)
	}

	if success < majority {
		return nil, ErrNoQuorum
	}

	return final, nil
}