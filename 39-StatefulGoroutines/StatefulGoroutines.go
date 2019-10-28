/*
 * @Descripttion:
 * @version:
 * @Author: renjie.zhang
 * @Date: 2019-10-28 22:28:46
 * @LastEditTime: 2019-10-28 22:42:07
 */
package main
	
import (
    "fmt"
    "math/rand"
    "sync/atomic"
    "time"
)

type readOpertion struct {
	key  int
	resp chan int
}

type writeOpertion struct {
	key  int
	val  int
	resp chan bool
}

func main() {
	var readOps uint64
	var writeOps uint64

	reads := make(chan readOpertion)
	writes := make(chan writeOpertion)

	go func() {
		var state = make(map[int]int)
		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}()

	for i := 0; i < 100; i++ {
		go func() {
			for {
				read := readOpertion{
					key:  rand.Intn(5),
					resp: make(chan int)}
				reads <- read
				<-read.resp
				atomic.AddUint64(&readOps,1)
				time.Sleep(time.Millisecond)
			}
		}()
	}
	for w := 0; w < 10; w++ {
        go func() {
            for {
                write := writeOpertion{
                    key:  rand.Intn(5),
                    val:  rand.Intn(100),
                    resp: make(chan bool)}
                writes <- write
                <-write.resp
                atomic.AddUint64(&writeOps, 1)
                time.Sleep(time.Millisecond)
            }
        }()
    }
    time.Sleep(time.Second)
    readOpsFinal := atomic.LoadUint64(&readOps)
    fmt.Println("readOps:", readOpsFinal)
    writeOpsFinal := atomic.LoadUint64(&writeOps)
    fmt.Println("writeOps:", writeOpsFinal)
}
