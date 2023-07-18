package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/amirrmonfared/myhttp/pkg/hasher"
	httpclient "github.com/amirrmonfared/myhttp/pkg/http-client"
	"github.com/amirrmonfared/myhttp/pkg/tools/semaphore"
	"github.com/amirrmonfared/myhttp/pkg/urlprocessor"
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

	httpClient := httpclient.NewDefaultHTTPClient()

	hasher := hasher.NewURLHasher()

	sem := semaphore.NewDefaultSemaphore(options.maxParallelRequests)

	urlProcessor := urlprocessor.NewURLProcessor(httpClient, hasher, sem)

	if err := urlProcessor.ProcessURLs(options.urls); err != nil {
		log.Fatal(err)
	}

	log.Println("URLs processed successfully")
}
