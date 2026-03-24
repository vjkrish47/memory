import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)
const (
	machines   = 3
	iterations = 200000
	bufferSize = 1000
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
func analyze(m Metric) int {
	s := 0
	for i := 0; i < 50; i++ {
		s += m.bw + m.q - m.loss
	}
	return s
}
func machine(id int, out chan<- Packet, wg *sync.WaitGroup) {
	buffer := make([]Metric, 0, bufferSize)

	for i := 0; i < iterations; i++ {
		m := collect()
		buffer = append(buffer, m)
		analyze(m)
		out <- Packet{id: id, m: m}
	}
	wg.Done()
}
func middleware(in <-chan Packet, out chan<- Packet, done chan<- bool) {
	for p := range in {
		out <- p
	}
	done <- true
}
func collector(in <-chan Packet, done chan<- bool) {
	total := 0
	for p := range in {
		total += analyze(p.m)
	}
	fmt.Println("Aggregate:", total)
	done <- true
}
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	start := time.Now()
	local := make(chan Packet, 1000)
	mid := make(chan Packet, 1000)
	done1 := make(chan bool)
	done2 := make(chan bool)
	var wg sync.WaitGroup
	go middleware(local, mid, done1)
	go collector(mid, done2)
	for i := 0; i < machines; i++ {
		wg.Add(1)
		go machine(i, local, &wg)
	}
	wg.Wait()
	close(local)
	<-done1
	close(mid)
	<-done2
	fmt.Println("Elapsed:", time.Since(start))
}
