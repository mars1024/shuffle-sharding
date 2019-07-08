package main

import (
	"fmt"
	"math/rand"
)

var handQueue []int

func shuffleSharding(v, queue uint64, handSize int) []int {
	if handSize < 1 {
		return nil
	}

	ret := make([]int, handSize)

	for i := 0; i < handSize; i++ {
		ret[i] = int(v % (queue - uint64(i)))
	}

	fmt.Printf("A set: %+v \n", ret)

	for i := 0; i < int(queue); i++ {
		handQueue[i] = 0
	}

	for i := 0; i < handSize; i++ {
		ii, ai := 0, ret[i]+1
		for {
			if handQueue[ii] == 0 {
				ai--
			}
			if ai == 0 {
				handQueue[ii] = 1
				break
			}
			ii++
		}
		ret[i] = ii
	}

	fmt.Printf("I set: %+v \n", ret)

	return ret
}

func main() {
	numQueues := uint64(8)
	handSize := 8

	queue := make([]int, numQueues)
	handQueue = make([]int, numQueues)
	retries := 10

	for i := 0; i < retries; i++ {
		hashValue := rand.Uint64()
		indices := shuffleSharding(hashValue, numQueues, handSize)
		for _, index := range indices {
			// fill queue
			queue[index]++
		}
	}

	fmt.Printf("Queue statistics: %+v \n", queue)
}
