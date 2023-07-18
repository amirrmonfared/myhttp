package main

import (
	"flag"
	"fmt"
	"log"
)

const (
	defaultMaxParallelRequests = 10
)

type options struct {
	maxParallelRequests int
	urls                []string
}

func (o *options) validateOptions() error {
	if o.maxParallelRequests < 1 {
		return fmt.Errorf("parallel must be greater than 0")
	}

	if len(o.urls) == 0 {
		return fmt.Errorf("no URLs provided")
	}

	return nil
}

func gatherOptions() options {
	o := options{}

	flag.IntVar(&o.maxParallelRequests, "parallel", defaultMaxParallelRequests, "Maximum number of parallel requests")
	flag.Parse()

	o.urls = flag.Args()

	return o
}

func main() {
	options := gatherOptions()
	if err := options.validateOptions(); err != nil {
		log.Fatal(fmt.Errorf("invalid options: %s", err))
	}

	log.Println("URLs processed successfully")
}
