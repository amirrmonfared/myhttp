package urlprocessor

import (
	"fmt"
	"net/url"
	"sync"

	"github.com/amirrmonfared/myhttp/pkg/hasher"
	httpclient "github.com/amirrmonfared/myhttp/pkg/http-client"
	"github.com/amirrmonfared/myhttp/pkg/tools/semaphore"
)

type URLProcessor struct {
	HTTPClient httpclient.HTTPClient
	Hasher     hasher.Hasher
	Semaphore  semaphore.Semaphore
}

func NewURLProcessor(httpClient httpclient.HTTPClient, hasher hasher.Hasher, semaphore semaphore.Semaphore) *URLProcessor {
	return &URLProcessor{
		HTTPClient: httpClient,
		Hasher:     hasher,
		Semaphore:  semaphore,
	}
}

func (p *URLProcessor) ProcessURLs(urls []string) error {
	var wg sync.WaitGroup
	wg.Add(len(urls))

	parsedURLs, err := parseURLs(urls)
	if err != nil {
		return err
	}

	var errors []error
	for _, parsedURL := range parsedURLs {
		p.Semaphore.Acquire()
		go func(parsedURL url.URL) {
			defer wg.Done()
			defer p.Semaphore.Release()
			if err := p.processURL(parsedURL); err != nil {
				errors = append(errors, err)
			}
		}(parsedURL)
	}

	wg.Wait()

	if len(errors) > 0 {
		return fmt.Errorf("%v", errors)
	}

	return nil
}

func (p *URLProcessor) processURL(url url.URL) error {
	resp, err := p.HTTPClient.Get(url.String())
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	hash, err := p.Hasher.Hash(resp.Body)
	if err != nil {
		return err
	}

	fmt.Printf("%s %s\n", url.String(), hash)

	return nil
}

func parseURLs(urlArgs []string) ([]url.URL, error) {
	var urls []url.URL
	for _, urlArg := range urlArgs {
		url, err := url.Parse(urlArg)
		if err != nil {
			return nil, err
		}

		urls = append(urls, normalizeURL(*url))
	}
	return urls, nil
}

func normalizeURL(url url.URL) url.URL {
	if url.Scheme == "" {
		url.Scheme = "http"
	}

	return url
}
