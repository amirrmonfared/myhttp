# myhttp

`myhttp` is a command-line tool written in Go for making HTTP requests to multiple URLs and printing the address of each request along with the MD5 hash of the response. The tool is designed to perform requests in parallel to optimize processing time.

## Installation

Before using `myhttp`, ensure that you have Go installed and properly configured on your computer. For Go installation instructions, please refer to the [official Go installation guide](https://golang.org/doc/install).

## Usage

To use `myhttp`, run the following command in your terminal:

```bash
go run ./cmd/myhttp [flags] [URL1] [URL2] ... [URLN]
```

The following flag is available:
````bash
  -parallel int
        number of parallel requests (default 10)
````

## Examples

```bash
go run ./cmd/myhttp -parallel 5 http://www.google.com http://www.yahoo.com http://www.bing.com
```

```bash
go run ./cmd/myhttp -parallel 2 http://www.google.com yahoo.com www.bing.com
```

## Testing

To run the unit tests, run the following command in your terminal:

```bash
go test ./...
```

## Dependencies
MyHTTP does not have any external dependencies beyond the Go standard library. All the required functionality is implemented using the built-in packages provided by Go.

