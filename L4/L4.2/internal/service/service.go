package service

import (
	"context"

	"DistributedGrep/internal/app"
	"DistributedGrep/internal/grep"
	"DistributedGrep/internal/storage"
)

type GrepService struct {
	workers int
	nodeName string
}

func NewService(workers int, nodeName string) *GrepService {
	return &GrepService{
		workers:  workers,
		nodeName: nodeName,
	}
}

func (s *GrepService) Search(
	ctx context.Context,
	req app.GrepRequest,
) (*app.GrepResponse, error) {

	lines, errCh := storage.ReadLines(req.File)

	results := grep.RunWorkers(
		s.workers,
		req,
		lines,
	)

	resp := &app.GrepResponse{
		Node: s.nodeName,
	}

	for {

		select {

		case <-ctx.Done():
			return nil, ctx.Err()

		case err := <-errCh:

			if err != nil {
				return nil, err
			}

			errCh = nil

		case line, ok := <-results:

			if !ok {

				if req.CountOnly {
					resp.Lines = nil
				}

				return resp, nil
			}

			resp.Count++

			if !req.CountOnly {
				resp.Lines = append(
					resp.Lines,
					line,
				)
			}
		}
	}
}