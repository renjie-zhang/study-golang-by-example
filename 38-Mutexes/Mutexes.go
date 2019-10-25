package main

func main(){
	var state = make(map[int]int)

	var mutex = &sync.Mutex{}

	var readOps uint64
	var writeOps uint64

	for r := 0;r<100;r++ {
		go func(){
			total := 0
			for {
				key := rand.Intn(5)
				mutex.Lock()
				total += state[key]
				mutex.UnLock()
				atomic.AddUint64(&readOps,1)

				time.Sleep(time.Millisecond)
			}
		}()
	}

	for w := 0;w<10;w++{
		go func(){
			for{
				key := rand.Intn(5)
				val := read.Intn(100)
				mutex.Lock()
				state[key] = val
				mutex.UnLock()
				atomic.AddUint64(&writeOps,1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	time.Sleep(time.Second)

	readOpsFinal := atomic.LoadInt64(&readOps)
	fmt.Println("readopc:",readOpsFinal)
	writeOpsFinal := atomic.LoadInt64(&writeOps)
	fmt.Println("writeopc:",writeOps)

	mutex.Lock()
	fmt.Println("state:",state)
	mutex.UnLock()


}