package urlprocessor

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/amirrmonfared/myhttp/pkg/hasher"
	httpclient "github.com/amirrmonfared/myhttp/pkg/http-client"
	"github.com/amirrmonfared/myhttp/pkg/tools/semaphore"
)

func TestURLProcessor_ProcessURLs(t *testing.T) {
	tests := []struct {
		name       string
		HTTPClient httpclient.HTTPClient
		Hasher     hasher.Hasher
		Semaphore  semaphore.Semaphore
		urls       []string
		wantErr    bool
	}{
		{
			name: "OK; with correct URL",
			HTTPClient: &httpclient.FakeHTTPClient{
				GetResponse: &http.Response{
					StatusCode: http.StatusOK,
					Body:       io.NopCloser(strings.NewReader("response body")),
				},
				Err: nil,
			},
			Hasher: &hasher.FakeHasher{
				HashValue: "hashed",
				HashError: nil,
			},
			Semaphore: &semaphore.FakeSemaphore{},
			urls:      []string{"http://example.com"},
			wantErr:   false,
		},
		{
			name: "ok; without scheme in URL(with normalize)",
			HTTPClient: &httpclient.FakeHTTPClient{
				GetResponse: &http.Response{
					StatusCode: http.StatusOK,
					Body:       io.NopCloser(strings.NewReader("response body")),
				},
				Err: nil,
			},
			Hasher: &hasher.FakeHasher{
				HashValue: "hashed",
				HashError: nil,
			},
			Semaphore: &semaphore.FakeSemaphore{},
			urls:      []string{"invalid-url.com"},
			wantErr:   false,
		},
		{
			name: "Error; with hash error",
			HTTPClient: &httpclient.FakeHTTPClient{
				GetResponse: &http.Response{
					StatusCode: http.StatusOK,
					Body:       io.NopCloser(strings.NewReader("response body")),
				},
				Err: nil,
			},
			Hasher: &hasher.FakeHasher{
				HashValue: "",
				HashError: fmt.Errorf("hash error"),
			},
			Semaphore: &semaphore.FakeSemaphore{},
			urls:      []string{"http://example.com"},
			wantErr:   true,
		},
		{
			name: "Error; with http client error",
			HTTPClient: &httpclient.FakeHTTPClient{
				GetResponse: nil,
				Err:         fmt.Errorf("http client error"),
			},
			Hasher: &hasher.FakeHasher{
				HashValue: "hashed",
				HashError: nil,
			},
			Semaphore: &semaphore.FakeSemaphore{},
			urls:      []string{"http://example.com"},
			wantErr:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &URLProcessor{
				HTTPClient: tt.HTTPClient,
				Hasher:     tt.Hasher,
				Semaphore:  tt.Semaphore,
			}
			if err := p.ProcessURLs(tt.urls); (err != nil) != tt.wantErr {
				t.Errorf("URLProcessor.ProcessURLs() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
