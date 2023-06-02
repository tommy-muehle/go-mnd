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
	strconv.FormatComplex(10+3i, 'f', 0, 32)
}

func example14() {
	_, _ = strconv.ParseInt("10", 10, 64)
	_, _ = strconv.ParseUint("5", 10, 64)
	_, _ = strconv.ParseFloat("5.0", 32)
	_, _ = strconv.ParseComplex("10+3i", 32)
}

func stringsSplit() {
	_ = strings.SplitN("a,b,c", ",", 2)
	_ = strings.SplitAfterN("a,b,c", ",", 2)
}

func osChmod() {
	_ = os.Chmod("some-filename", 0644)
}

func osMkdir() {
	_ = os.Mkdir("testdir", 0750)
}

func osMkdirAll() {
	_ = os.MkdirAll("test/subdir", 0750)
}

func osOpenFile() {
	_, _ = os.OpenFile("notes.txt", os.O_RDWR|os.O_CREATE, 0755)
}

func osWriteFile() {
	_ = os.WriteFile("testdata/hello", []byte("Hello, Gophers!"), 0666)
}
