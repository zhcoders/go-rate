# go-rate

Testing with https://github.com/beefsack/go-rate

```
func TestRateLimiter(t *testing.T) {
	limit := 5
	interval := time.Millisecond
	limiter := New(limit, interval)
	count := 0
	for i := 0; i < 100000000; i++ {
		ok := limiter.Try()
		if ok {
			count++
		}
	}
	fmt.Println("TestRateLimiter success count", count)
}

func TestMyRateLimiter(t *testing.T) {
	limit := 5
	interval := time.Millisecond
	limiter := NewMutexRateLimiter(limit, interval)
	count := 0
	for i := 0; i < 100000000; i++ {
		ok := limiter.Try()
		if ok {
			count++
		}
	}
	fmt.Println("TestMyRateLimiter success count", count)
}

func TestChanRateLimiter(t *testing.T) {
	limit := 5
	interval := time.Millisecond
	limiter := NewChanRateLimiter(limit, interval)
	count := 0
	for i := 0; i < 100000000; i++ {
		ok := limiter.Try()
		if ok {
			count++
		}
	}
	fmt.Println("TestChanRateLimiter success count", count)
}

func TestForLoop(t *testing.T) {
	for i := 0; i < 100000000; i++ {
	}
}

```
