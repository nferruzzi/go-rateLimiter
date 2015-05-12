# go-rateLimiter
Schedule go subroutine using a simple rate limiter (yes another one!) :)

example:

```go
func main() {
	limiter := RateLimiter.NewRateLimiter(time.Second / 10)

	n := time.Now()
	for i := 0; i < 10; i++ {
		log.Print("Scheduled:", i)

		e := i
		f := func() {
			log.Printf("%d: %v", e, time.Now().Sub(n))
		}

		limiter.Go(f)
	}

	time.Sleep(time.Second)
}
```

output:

```
2015/05/12 23:15:59 Scheduled:0
2015/05/12 23:15:59 Scheduled:1
2015/05/12 23:15:59 Scheduled:2
2015/05/12 23:15:59 Scheduled:3
2015/05/12 23:15:59 Scheduled:4
2015/05/12 23:15:59 Scheduled:5
2015/05/12 23:15:59 Scheduled:6
2015/05/12 23:15:59 Scheduled:7
2015/05/12 23:15:59 Scheduled:8
2015/05/12 23:15:59 Scheduled:9
2015/05/12 23:15:59 0: 326.23µs
2015/05/12 23:15:59 1: 100.50801ms
2015/05/12 23:15:59 2: 200.477413ms
2015/05/12 23:16:00 3: 300.825863ms
2015/05/12 23:16:00 4: 401.17095ms
2015/05/12 23:16:00 5: 501.532535ms
2015/05/12 23:16:00 6: 601.363186ms
2015/05/12 23:16:00 7: 701.697503ms
2015/05/12 23:16:00 8: 801.24942ms
2015/05/12 23:16:00 9: 901.623269ms
```
