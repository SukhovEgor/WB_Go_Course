package or

import (
	"testing"
	"time"
)

func signal(after time.Duration) <-chan interface{} {

	ch := make(chan interface{})

	go func() {
		defer close(ch)
		time.Sleep(after)
	}()

	return ch
}

func TestOr_NoChannels(t *testing.T) {

	done := Or()

	if done != nil {
		t.Fatal("expected nil channel")
	}
}

func TestOr_OneChannel(t *testing.T) {

	done := signal(50 * time.Millisecond)

	start := time.Now()

	<-Or(done)

	if time.Since(start) > 100*time.Millisecond {
		t.Fatal("channel closed too late")
	}
}

func TestOr_TwoChannels(t *testing.T) {

	start := time.Now()

	<-Or(
		signal(50*time.Millisecond),
		signal(time.Second),
	)

	elapsed := time.Since(start)

	if elapsed > 150*time.Millisecond {
		t.Fatalf("expected first channel, got %v", elapsed)
	}
}

func TestOr_ManyChannels(t *testing.T) {

	start := time.Now()

	<-Or(
		signal(2*time.Second),
		signal(1500*time.Millisecond),
		signal(1200*time.Millisecond),
		signal(700*time.Millisecond),
		signal(100*time.Millisecond),
		signal(900*time.Millisecond),
		signal(3*time.Second),
	)

	elapsed := time.Since(start)

	if elapsed > 200*time.Millisecond {
		t.Fatalf("expected earliest channel, got %v", elapsed)
	}
}

func TestOr_ClosedChannel(t *testing.T) {

	ch := make(chan interface{})
	close(ch)

	select {

	case <-Or(ch):

	case <-time.After(time.Second):
		t.Fatal("expected immediate close")
	}
}

func TestOr_WithNilChannel(t *testing.T) {

	start := time.Now()

	<-Or(
		nil,
		signal(100*time.Millisecond),
	)

	elapsed := time.Since(start)

	if elapsed > 200*time.Millisecond {
		t.Fatal("nil channel blocked")
	}
}

func TestOr_Stress(t *testing.T) {

	channels := make([]<-chan interface{}, 0, 100)

	for i := 0; i < 99; i++ {
		channels = append(
			channels,
			signal(5*time.Second),
		)
	}

	channels = append(
		channels,
		signal(100*time.Millisecond),
	)

	start := time.Now()

	<-Or(channels...)

	if time.Since(start) > 300*time.Millisecond {
		t.Fatal("stress test failed")
	}
}