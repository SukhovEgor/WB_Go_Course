package main

import (
	"fmt"
	"log"

	"DistributedGrep/internal/app"
	"DistributedGrep/internal/config"
	"DistributedGrep/internal/network"
	"DistributedGrep/internal/quorum"

	"flag"
)

func main() {

	pattern := flag.String("pattern", "", "search pattern")
	file := flag.String("file", "", "input file")
	ignore := flag.Bool("i", false, "ignore case")
	invert := flag.Bool("v", false, "invert match")
	count := flag.Bool("count", false, "count only")

	flag.Parse()

	cfg := config.MustLoad("config/local.yaml")

	client := network.NewClient()

	q := quorum.New(client)

	req := app.GrepRequest{
		Pattern:    *pattern,
		File:       *file,
		IgnoreCase: *ignore,
		Invert:     *invert,
		CountOnly:  *count,
	}

	resp, err := q.Execute(
		cfg.Cluster.Nodes,
		req,
	)

	if err != nil {
		log.Fatal(err)
	}

	if req.CountOnly {

		fmt.Println(resp.Count)
		return
	}

	for _, line := range resp.Lines {
		fmt.Println(line)
	}
}