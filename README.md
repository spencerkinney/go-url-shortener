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

## Example Usage

### Shorten a URL

```sh
curl -X POST -H "Content-Type: application/json" -d "{\"url\": \"https://www.example.com\"}" http://localhost:8080/shorten
```

**Example Response:**

```json
{
  "shortUrl": "http://localhost:8080/NHMJWL"
}
```

### Access a Shortened URL

Visit `http://localhost:8080/NHMJWL` in your web browser, or use `curl` to be redirected to the original URL.

<br>

## URL Key Space

The short URL key is a 6-char string of uppercase letters and numbers, there are 36 possible characters for each position. With 6 characters in the key, this allows for 36<sup>6</sup> (2,176,782,336) unique combinations, making the probability of collisions very low with a moderate number of URLs. I'll prob later add proper collisions handling but for now let's just say its unlikely.

## rando thoughts on cache (not final)

the URL shortener should handle high throughput.

- **In-Memory Caching**: maybe use a co-located Redis which should have decent read/write speeds.

- **Data Storage**: Shortened URLs serve as keys that map to their original URL counterparts?

- **Cache Eviction**: maybe have a LRU eviction policy to maintain frequently accessed URLs and discard the least used ones.

- **Data Freshness**: Cache entries have a TTL to guarantee up-to-date information, with cache invalidation on updates to preserve data consistency.

ideally, the URL shortener can handle a large number of requests fast while maintaining data accuracy.
