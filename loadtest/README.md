# Load Test for go-url-shortener

## Prerequisites
- go-url-shortener server is up at `http://localhost:8080/shorten`
- Go is installed (get it from [Go's official site](https://golang.org/dl/))

## Run the Load Test
1. Open a new terminal.
2. Change directory to where `loadtest.go` is:  
   `cd path/to/loadtest`
3. Run the script:  
   `go run loadtest.go`

## Notes
- The script sends many HTTP POST requests to the `/shorten` endpoint.
- Adjust `requestsNumber` and `concurrency` in the script as needed.
- Results show total and average time for requests.
- you can pip install locust and python and run `locust -f locustfile.py` to do a load test, might be slower than using Jmeter though
