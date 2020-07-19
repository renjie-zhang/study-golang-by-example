/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package main

import (
	"fmt"
	"time"
)

func main() {
	channelDemo()

	synchronizationByChannel()

	closeChannel()

	rangOverChannel()
}

func channelDemo() {
	messages := make(chan string)

	go func() {
		messages <- "ping"
	}()
	go func() {
		messages <- "pong"
	}()

	for i := 0; i < 2; i++ {
		fmt.Println("messages", <-messages)
	}

	bufferChannel := make(chan string, 2)
	bufferChannel <- "one"
	bufferChannel <- "two"
	fmt.Println("one message :", <-bufferChannel)
	fmt.Println("two message :", <-bufferChannel)
}

func synchronizationByChannel() {
	done := make(chan bool, 1)
	go worker(done)
	<-done
}

func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("DONE")
	done <- true
}

func closeChannel() {
	jobs := make(chan int, 5)
	done := make(chan bool, 1)
	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Printf("receive a job, job is %d\n", j)
			} else {
				fmt.Printf("don't receive a job\n")
				done <- true
				return
			}
		}
	}()

	for j := 0; j <= 3; j++ {
		jobs <- j
		fmt.Printf("send a job %d\n", j)
	}

	close(jobs)

	fmt.Println("send all jobs")

	<-done
}

func rangOverChannel(){
	demo := make(chan int,2)
	demo <- 1
	demo <-2
	close(demo)
	for v := range demo {
		fmt.Printf("a element %d from channel\n",v)
	}
}
