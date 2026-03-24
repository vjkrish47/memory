import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

const (
	nodes      = 5
	iterations = 200000
)

type Metric struct {
	bw   int
	loss int
	q    int
}

type Packet struct {
	id int
	m  Metric
}

func collect() Metric {
	return Metric{
		bw:   rand.Intn(1000),
		loss: rand.Intn(100),
		q:    rand.Intn(500),
	}
}

func nodeWorker(id int, out chan<- Packet, wg *sync.WaitGroup) {
	for i := 0; i < iterations; i++ {
		m := collect()
		out <- Packet{id: id, m: m}
	}
	wg.Done()
}

func middleware(in <-chan Packet, done chan<- bool) {
	store := make([]Metric, 0)

	for p := range in {
		store = append(store, p.m)
	}

	total := 0
	for _, m := range store {
		total += m.bw + m.q - m.loss
	}

	fmt.Println("Aggregate:", total)
	done <- true
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	start := time.Now()

	ch := make(chan Packet, 1000)
	done := make(chan bool)

	var wg sync.WaitGroup

	go middleware(ch, done)

	for i := 0; i < nodes; i++ {
		wg.Add(1)
		go nodeWorker(i, ch, &wg)
	}

	wg.Wait()
	close(ch)
	<-done

	fmt.Println("Elapsed:", time.Since(start))
}
