package argument

import (
	"context"
	"math"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

var MyValue = float32(20) // want "Magic number: 20"

const MyConstantValue = float32(20)

const (
	FooBlockTypeConversion = int64(512)
	BarBlockTypeConversion = uint8(2)
)

func example() {
	math.Abs(9.5) // want "Magic number: 9.5"
}

func example2() {
	strings.ToUpper("foo")
}

type s struct {
	count int
}

func (s *s) Add(count int) {
	s.count += count
}

func example3() {
	s := &s{}
	s.Add(10) // want "Magic number: 10"
}

func example4() {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second) // want "Magic number: 5"
	defer cancelFunc()

	ctx.Value("foo")
}

func example5() {
	http.StatusText(200) // want "Magic number: 200"
}

func example6() string {
	t := time.Date(2017, time.September, 26, 12, 13, 14, 0, time.UTC)
	return t.String()
}

func example7() {
	c := make(chan int, 1)
	c <- 1
}

func example8() {
	c := make(chan int, 2) // want "Magic number: 2"
	c <- 1
}

func example9() {
	os.Exit(1)
}

func example10() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		duration := 500 * time.Millisecond // want "Magic number: 500"
		time.Sleep(duration)
		wg.Done()
	}()
}

func foobar(a, b int) {}

func example11() {
	foobar(0, 3) // want "Magic number: 3"
}

func foobaz(numbers ...int) {}

func example12() {
	foobaz(0, 1, 3, 5, 6) // want "Magic number: 3" "Magic number: 5" "Magic number: 6"
}

func example13() {
	strconv.FormatInt(10, 32)
	strconv.FormatUint(5, 32)
	strconv.FormatFloat(5.0, 'E', -1, 32)
}

func example14() {
	_, _ = strconv.ParseInt("10", 10, 64)
	_, _ = strconv.ParseUint("5", 10, 64)
	_, _ = strconv.ParseFloat("5.0", 32)
}
