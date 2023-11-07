# go-url-shortener

A simple URL shortener written in Go for the purpose of learning Go.

## Purpose

This project is just for learning Go and an attempt at creating an ultra-fast URL shortener. The goal is to programmatically generate 10k+ shortened URLs per second (probably more), exploring a bit with Go's concurrency model and standard lib. Feel free to steal any code here or contribute whatever you think would be cool. I will be hosting this project on the domain **[301.li](http://301.li)**.

## Use Cases?

- Efficiently handle massive URL redirection for social media platforms.
- Provide quick and reliable link shortening for marketing and advertising campaigns.
- Offer enterprise-grade link management with robust tracking and analytics.

## Q&A

**Q:** Why?

**A:** Why tf not?

## Production Goals

I want to create a production-ready service that is free for anyone to use, incorporating:
- Enhanced security features.
- Comprehensive analytics for users.
- Scalable architecture to handle growth in demand.
- A friendly UI for easy interaction with the service.
- minimum 99.999% uptime with no cold starts - I'll be tracking that

## Installation

```sh
# Clone the repository
git clone https://github.com/spencerkinney/go-url-shortener.git
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

The short URL key is a 6-char string of uppercase letters and numbers, there are 36 possible characters for each position. With 6 characters in the key, this allows for 36<sup>6</sup> (2,176,782,336) unique combinations, making the probability of collisions very low with a moderate number of URLs. I'll prob later add proper collisions handling but for now let's just say its unlikely. This is what I'll currently use but might change to 8-char strings later.

## Initial Benchmarks

Just got the initial code running and here are the initial performance numbers. It's a start, but honestly...

**Total time for 10000 requests with 100 concurrency:** 1.2382327s  
**Average request time:** 12.311669ms  
**Median request time:** 12.0107ms  
**Requests per second:** 8076.026421  
**Min request time:** < 1ns
**Max request time:** 27.7081ms

...zzzz, this is SLOW. I'll make some changes later to make some easy perf enhancements. this was just the original code to get the thing working. 26ms max is EMBARRASSING. i'll fix later. I also need to check over my loadtest code to make sure it is properly and accurately measuring durations. as some point, I'll set up a proper benchmarking tool. btw this was ran on a AMD Ryzen 7 1700.