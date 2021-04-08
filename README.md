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
The result

![image](https://user-images.githubusercontent.com/20734293/113982275-25f2ca80-987b-11eb-8935-42e5af4f4f89.png)

			count++
