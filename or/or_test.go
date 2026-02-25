package or

import (
	"testing"
	"time"
)

func sig(after time.Duration) <-chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(after)
	}()
	return c
}
func TestOr(t *testing.T) {
	start := time.Now()
	//тут важный момент:
	//Если будет очень много горутин и самая быстрая из них передается в конце функции
	//то функция может завершиться немного позже
	//чтобы это избежать можно просто расширить case на случай если переданы 2, 3, ... каналов
	//тогда задержка сократиться из-за того что будет меньше рекурсия, но код менее читабелен
	<-Or(
		sig(2*time.Hour),
		sig(2*time.Minute),
		sig(3*time.Minute),
		sig(3*time.Second),
		sig(3*time.Hour),
		sig(1*time.Second),
	)
	elapsed := time.Since(start)
	if elapsed >= 1100*time.Millisecond {
		t.Fatalf("Or did not return early enough, delayed for %v", elapsed)
	}
}
func TestOrZeroChannels(t *testing.T) {
	if Or() != nil {
		t.Fatal("Expected nil for zero channels")
	}
}

func TestOrSingleChannel(t *testing.T) {
	ch := make(chan interface{})
	result := Or(ch)

	if result != ch {
		t.Fatal("Expected same channel for single input")
	}
}
