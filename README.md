# go-url-shortener

A simple URL shortener written in Go for the purpose of learning Go.

## Installation

```sh
# Clone the repository
git clone https://github.com/spencerkinney/go-url-shortener.git
# Navigate into the project directory
cd go-url-shortener
# Build the project
go build
# Run the server
./go-url-shortener
```

<br>

## Usage

### Shorten a URL

```sh
curl -X POST -H "Content-Type: application/json" -d '{"url": "https://www.example.com"}' http://localhost:8080/shorten
```

**Response:**

```json
{
  "shortUrl": "http://localhost:8080/abc123"
}
```

### Access a Shortened URL

Visit `http://localhost:8080/abc123` in your web browser, or use `curl` to be redirected to the original URL.

<br>

## rando thoughts on cache (not final)

the URL shortener should handle high throughput.

- **In-Memory Caching**: maybe use a co-located Redis which should have decent read/write speeds.

- **Data Storage**: Shortened URLs serve as keys that map to their original URL counterparts?

- **Cache Eviction**: maybe have a LRU eviction policy to maintain frequently accessed URLs and discard the least used ones.

- **Data Freshness**: Cache entries have a TTL to guarantee up-to-date information, with cache invalidation on updates to preserve data consistency.

ideally, the URL shortener can handle a large number of requests fast while maintaining data accuracy.
